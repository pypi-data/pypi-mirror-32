# EasyMunk Python add-on for Munkres

EasyMunk provides a versatile interface alternative and pre-processing handler to the Munkres module for score-based (Hungarian algorithm) assignments.

## Purpose

The EasyMunk module interfaces with the Munkres module to:
1. Provide an easy and adaptable method of pair assignment.
2. Reduce the amount of pre-processing required for munkres.
3. Automatically account for some special-case munkres decisions.
4. Do all of the above without reducing munkres' functionality.

What might one use the Hungarian algorithm to accomplish?

* Assigning pairs.
  * For example: objects in images to information about those objects.
* Assigning tasks.
* Evaluating changing cost attributes over function iterations. (duplicate your primaries per iteration and forbid!)
* Handling and evaluating all potential values for the best single value for every actor.
* And more!

See this module's easymunk.py for use instructions, or the [munkres documentation][] for more algorithm details.

[munkres documentation]: https://github.com/bmc/munkres

## Installation

1. Ensure Python and Pip are in the environment path variable.
2. Inside a terminal; enter the command "pip install easymunk"

Built in Python 3.6 - should now be compatible with Python 2.6+ and Python 3.6+.

## Use

* At the top of your .py file, add "from easymunk import EasyMunk"
* Although EasyMunk is a class, self is never used. You can call easymunk.EasyMunk class methods without creating an instance.
* The two public methods are EasyMunk.print_info() and EasyMunk.sort()

### EasyMunk.print_info()
Prints the most recent .sort() method's profit (if it exists) and cost matrices using munkres.print_matrix() for readability, the chosen indices in the (primary objects, secondary objects) array created by EasyMunk, and the sum total cost & profit for the chosen solution.

### EasyMunk.sort()

Please see this module's Documentation at the top of ./easymunk/easymunk.py for further explanation & examples, and at the top of the EasyMunk.score() method definition for further details on parameters and functions.

### External Functions

All external function parameters passed as positional, excluding \*\*kwargs. Optional parameters (such as list of lists of dicts containing pair score / primary object / secondary object, or \*args) which conflict with an external function's call are removed starting with \*args, then \*\*kwargs; at which point the least likely optional parameter is removed and \*args and \*\*kwargs are introduced again before continuing the cull cycle to just the minimum parameters. The working function parameter syntax is cached, but this process will always run a full cycle starting with the last cached syntax if a TypeError is met - meaning you can have dynamic method parameters or methods.

### Parameters

* list primary_objects: List of anything you wish to assign an object/task/etc.
* list secondary_objects: List of anything you wish to assign primary_objects to.
* function score_function: Score function for each possible pair; must return a numeric metric.
  * score_function input: primary_object, secondary_object. 
  * *dynamically-detected **optional*** input: \*args, \*\*kwargs.
  * score_function return float: Your cost function's return should be an int or float or long int or munkres.DISALLOWED 
                                 representing that match's cost or profit.
* function assignment_function: Function for operating on pair assignments externally.
  * assignment_function input: primary_object, secondary_object. 
  * *dynamically-detected **optional*** input: \*args, \*\*kwargs.
  * assignment_function return None: Your assignment function should not return anything.
* **optional** bool convert_profit_to_cost: Determine if profit to cost function is used. Default is False.
* **optional** function profit_to_cost_function: Run after score_function if convert_profit_to_cost is True. 
  * A default conversion function of (cost_index_value = max(profit_array) - profit_index_value) is used if left as None.
  * profit_to_cost_function input: single profit value. 
  * *dynamically-detected **optional*** input: value's primary object, value's secondary object, entire list of lists of 
  profit dicts as {"score": profit, "primary": primary_object, "secondary": secondary_object}, \*args, \*\*kwargs.
  * profit_to_cost_function return float: Your conversion function's return should be the same
                                          type as the cost_function's return; in other words, a numeric value for pair score.
* **optional** tuple \*args: Optional arguments passed to score_function, assignment_function, and profit_to_cost_function.
* **optional** dict \*\*kwargs: Optional keyword arguments passed to score_function, assignment_function, and profit_to_cost_function.

### Class Variables

EasyMunk.change_function_syntax_defaults_and_order() & EasyMunk.change_default_dict_keys()

* There is no need to change either of these. However, you might want to do so anyways without changing a \_private attribute. Recommendations on using the provided EasyMunk class attributes to change different operations such as dict key names and first-attempt function syntax / syntax attempt order to be completed. However, documentation on what each class variable is responsible for is available within the EasyMunk module.

## Copyright

EasyMunk module by Andrew M. Hogan. (EasyMunk &copy; 2018 Hogan Consulting Group)

## License

Licensed under the Apache License.
