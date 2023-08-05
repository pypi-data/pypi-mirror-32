from functools import partial, wraps
import copy

MSG_CHECK_FALLBACK = "Your submission is incorrect. Try again!"

def requires_ast(f):
    @wraps(f)
    def wrapper(*args, **kwargs):
        state = kwargs.get('state', args[0] if len(args) else None)
        state_ast = [state.student_ast, state.solution_ast]

        # fail if no ast parser in use
        if any(ast is None for ast in state_ast): 
            raise TypeError("Trying to use ast, but it is None. Are you using a parser?")

        # check whether the parser passed or failed for some code
        ParseError = state.ast_dispatcher.ParseError

        parse_fail = any(isinstance(ast, ParseError) for ast in state_ast)

        if parse_fail: return state              # skip test
        else: return f(*args, **kwargs)          # proceed with test

    return wrapper

@requires_ast
def check_node(state, name, index=0, missing_msg="Could not find the {index}{node_name}.", priority=None):
    """Select a node from abstract syntax tree (AST), using its name and index position.
    
    Args:
        state: State instance describing student and solution code. Can be omitted if used with Ex().
        name : the name of the abstract syntax tree node to find.
        index: the position of that node (see below for details).
        missing_msg: feedback message if node is not in student AST.
        priority: the priority level of the node being searched for. This determines whether to
                  descend into other AST nodes during the search. Higher priority nodes descend
                  into lower priority. Currently, the only important part of priority is that 
                  setting a very high priority (e.g. 99) will search every node.
                

    :Example:
        If both the student and solution code are.. ::

            SELECT a FROM b; SELECT x FROM y;

        then we can focus on the first select with::
        
            # approach 1: with manually created State instance
            state = State(*args, **kwargs)
            new_state = check_node(state, 'SelectStmt', 0)
            
            # approach 2:  with Ex and chaining
            new_state = Ex().check_node('SelectStmt', 0)

    """
    df = partial(state.ast_dispatcher, name, slice(None), priority=priority)

    sol_stmt_list = df(state.solution_ast) 
    try: sol_stmt = sol_stmt_list[index]
    except IndexError: raise IndexError("Can't get %s statement at index %s"%(name, index))

    stu_stmt_list = df(state.student_ast)
    try: stu_stmt = stu_stmt_list[index]
    except IndexError: 
        # use speaker on ast dialect module to get message, or fall back to generic
        _msg = state.ast_dispatcher.describe(sol_stmt, missing_msg, index = index)
        state.do_test(_msg or MSG_CHECK_FALLBACK)

    action = {'type': 'check_node', 'kwargs': {'name': name, 'index': index}, 'node': stu_stmt}
    
    return state.to_child(student_ast = stu_stmt, solution_ast = sol_stmt,
                          history = state.history + (action,)) 


@requires_ast
def check_field(state, name, index=None, missing_msg="Could not find the {index}{field_name} of the {node_name}."):
    """Select an attribute from an abstract syntax tree (AST) node, using the attribute name.
    
    Args:
        state: State instance describing student and solution code. Can be omitted if used with Ex().
        name: the name of the attribute to select from current AST node.
        index: entry to get from field. If too few entires, will fail with missing_msg.
        missing_msg: feedback message if attribute is not in student AST.

    :Example:
        If both the student and solution code are.. ::
            
            SELECT a FROM b; SELECT x FROM y;

        then we can get the from_clause using ::

            # approach 1: with manually created State instance -----
            state = State(*args, **kwargs)
            select = check_node(state, 'SelectStmt', 0)
            clause = check_field(select, 'from_clause')

            # approach 2: with Ex and chaining ---------------------
            select = Ex().check_node('SelectStmt', 0)     # get first select statement
            clause =  select.check_field('from_clause')    # get from_clause (a list)
            clause2 = select.check_field('from_clause', 0) # get first entry in from_clause
    """
    try: 
        sol_attr = getattr(state.solution_ast, name)
        if index is not None: sol_attr = sol_attr[index]
    except IndexError: 
        raise IndexError("Can't get %s attribute"%name)

    try: 
        stu_attr = getattr(state.student_ast, name)
        if index is not None: stu_attr = stu_attr[index]
    except: 
        # use speaker on ast dialect module to get message, or fall back to generic
        _msg = state.ast_dispatcher.describe(state.student_ast, missing_msg, field = name, index = index)
        state.do_test(_msg or MSG_CHECK_FALLBACK)

    # fail if attribute exists, but is none only for student
    if stu_attr is None and sol_attr is not None:
        _msg = state.ast_dispatcher.describe(state.student_ast, missing_msg, field = name, index = index)
        state.do_test(_msg)

    action = {'type': 'check_field', 'kwargs': {'name': name, 'index': index}}

    return state.to_child(student_ast = stu_attr, solution_ast = sol_attr,
                          history = state.history + (action,)) 

import re

def test_student_typed(state, text, msg="Submission does not contain the code `{}`.", fixed=False):
    """Test whether the student code contains text.

    Args:
        state: State instance describing student and solution code. Can be omitted if used with Ex().
        text : text that student code must contain.
        msg  : feedback message if text is not in student code.
        fixed: whether to match text exactly, rather than using regular expressions.

    Note:
        Functions like ``check_node`` focus on certain parts of code.
        Using these functions followed by ``test_student_typed`` will only look
        in the code being focused on.

    :Example:
        If the student code is.. ::

            SELECT a FROM b WHERE id < 100

        Then the first test below would (unfortunately) pass, but the second would fail..::

            # contained in student code
            Ex().test_student_typed(text="id < 10")

            # the $ means that you are matching the end of a line
            Ex().test_student_typed(text="id < 10$")

        By setting ``fixed = True``, you can search for fixed strings::

            # without fixed = True, '*' matches any character
            Ex().test_student_typed(text="SELECT * FROM b")               # passes
            Ex().test_student_typed(text="SELECT \\\\* FROM b")             # fails
            Ex().test_student_typed(text="SELECT * FROM b", fixed=True)   # fails

        You can check only the code corresponding to the WHERE clause, using ::

            where = Ex().check_node('SelectStmt', 0).check_field('where_clause')
            where.test_student_typed(text = "id < 10)

    """
    stu_ast = state.student_ast
    stu_code = state.student_code

    # fallback on using complete student code if no ast
    ParseError = state.ast_dispatcher.ParseError
    stu_text = stu_ast._get_text(stu_code) if not isinstance(stu_ast, ParseError) else stu_code

    _msg = msg.format(text)

    # either simple text matching or regex test
    res = text in stu_text if fixed else re.search(text, stu_text)

    if not res:
        state.do_test(_msg)

    return state


@requires_ast
def has_equal_ast(state, 
                  msg="Check the {ast_path}. Your code does not seem to match the solution.",
                  sql=None,
                  start="sql_script",
                  exact=True):
    """Test whether the student and solution code have identical AST representations

    Args:
        state: State instance describing student and solution code. Can be omitted if used with Ex().
        msg  : feedback message if student and solution ASTs don't match
        sql  : optional code to use in place of the solution ast
        start: if sql arg is used, the parser rule to parse the sql code
        exact: whether to require an exact match (True), or only that the 
               student AST contains the solution AST.

    :Example:
        Suppose the student and solution code is `SeLeCt 1` and `SELECT 1`, respectively.
        In this case, the SCT `Ex().has_equal_ast()` will pass, since both
        select statements return identical ASTs.

        If the solution code is..::

            SELECT * FROM b WHERE id > 1 AND name = 'filip'

        Then the following SCT makes sure ``id > 1`` was used somewhere in the WHERE clause.::

            Ex().check_node('SelectStmt') \\
                .check_field('where_clause') \\/
                .has_equal_ast(sql = 'id > 1', start='expression', exact=False)
        
    """
    ast = state.ast_dispatcher.ast
    sol_ast = state.solution_ast if sql is None else ast.parse(sql, start)

    stu_rep = repr(state.student_ast)
    sol_rep = repr(sol_ast)

    _msg = msg.format(ast_path = state.get_ast_path())
    if       exact and (sol_rep != stu_rep):     state.do_test(_msg or MSG_CHECK_FALLBACK)
    elif not exact and (sol_rep not in stu_rep): state.do_test(_msg or MSG_CHECK_FALLBACK)

    return state

def verify_ast_parses(state):
    asts = [state.student_ast, state.solution_ast]
    if any(isinstance(c, state.ast_dispatcher.ParseError) for c in asts):
        state.do_test("AST did not parse")

    return state
