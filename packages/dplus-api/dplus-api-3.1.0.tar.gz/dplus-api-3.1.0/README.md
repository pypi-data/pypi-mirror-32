#The Dplus Python API


The D+ Python API allows using the D+ backend from Python, instead of the ordinary D+ application.

The Python API works on both Windows and Linux.

##Installation

Installing the Python API is done using PIP:

    pip install dplus-api
    
The API was tested with Python 3.5 and newer. It *may* work with older versions of Python, although Python 2 is probably not supported.

 ##Overview
Throughout the manual, code examples are given with filenames.
These files must be located in the same directory as the script itself, or alternately the code can be modified to the contain the full path of the file's location.

The overall structure of the Python API is as follows:

The data to be used for the calculation is built by the user in an instance of the `CalculationInput` class 
(either `GenerateInput` or `FitInput`). 

This instance is then passed to a `CalculationRunner` class (either `LocalRunner` or `WebRunner`),
and the calculation function is called (`generate`, `generate_async`, `fit`, or `fit_async`)

The `CalculationRunner` class returns a `CalculationResult` class.

Here is a very simple example of what this might look line in main.py: 

```
from dplus.CalculationInput import GenerateInput
from dplus.CalculationRunner import LocalRunner

calc_data = GenerateInput.load_from_state_file("mystate.state")
runner = LocalRunner()
result = runner.generate(calc_data)
print(result.graph)
```

A detailed explanation of the class types and their usage follows


##CalculationRunner

There are two kinds of CalculationRunners, Local and Web.

The LocalRunner is intended for users who have the D+ executable files installed on their system. It takes two optional
initialization arguments:

* `exe_directory` is the folder location of the D+ executables. By default, its value is None, and the python interface 
will attempt to find the executables on its own. 
* `session_directory` is the folder where the arguments for the calculation are stored, as well as the output results,
amplitude files, and pdb files, from the c++ executable. By default, its value is None, and an automatically generated 
temporary folder will be used. 

```
from dplus.CalculationRunner import LocalRunner

exe_dir = r"C:\Program Files\D+\bin"
sess_dir = r"sessions"
runner = LocalRunner(exe_dir, sess_dir)
#also possible:
#runner = LocalRunner()
#runner = LocalRunner(exe_dir)
#runner = LocalRunner(session_directory=sess_dir)
```

The WebRunner is intended for users accessing the D+ server. It takes two required initialization arguments, with no
default values:

* `url` is the address of the server.
* `token` is the authentication token granting access to the server. 

```
from dplus.CalculationRunner import WebRunner

url = r'http://localhost:8000/'
token = '4bb25edc45acd905775443f44eae'
runner = WebRunner(url, token)
```

Both runner classes have the same four methods: 

generate(calc_data), generate_async(calc_data), fit(calc_data), fit_async(calc_data)

All four methods take the same single argument, `calc_data` - an instance of a CalculationData class

generate and fit return a `CalculationResult`

generate_async and fit_async return a `RunningJob`

##RunningJob

The user should not be initializing this class. When returned from an async function in CalculationRunner, the user can 
use the following methods to interact with the RunningJob:

* `get_status()`: get a json dictionary reporting the job's current status
* `get_result(calc_data)`: get a `CalculationResult`. Requires a copy of the CalculationInput used to create the job. 
should only be called when job is completed. It is the user's responsibility to verify job completion with get_status 
before calling. 
* `abort()`: end a currently running job

```
from dplus.CalculationInput import GenerateInput
from dplus.CalculationRunner import LocalRunner

 calc_data = GenerateInput.load_from_state_file("mystate.state")
 runner = LocalRunner()
 job = runner.generate_async(calc_data)
 start_time = datetime.datetime.now()
 status = job.get_status()
 while status['isRunning']:
     status = job.get_status()
     run_time = datetime.datetime.now() - start_time
     if run_time > datetime.timedelta(seconds=50):
         job.abort()
         raise TimeoutError("Job took too long")
 result = job.get_result(calc_data)
```

##CalculationInput

There are two kinds of CalculationInput, FitInput and GenerateInput.

GenerateInput contains an instance of a `State` class and an x vector. It is used to generate the signal of a given
parameter tree.

FitInput contains a `State` class, an x vector, and a y vector. It is used to fit a parameter tree with an 
existing signal (the y vector) 

The `State` class is described in the next section.

The x and y vectors are simply lists of floating point coordinates. They can be generated from parameters in the state 
class or loaded from a file.
CalculationInput has the following methods:

* `get_model`: get a model by either its `name` or its `model_ptr`
* `get_models_by_type`: returns a list of `Models` with a given `type_name`, e.g. UniformHollowCylinder
* `get_mutable_params`: returns a list of `Parameters` in the state class, whose property `mutable` is True
* `get_mutable_parameter_values`: returns a list of floats, matching the values of the mutable parameters
* `set_mutable_parameter_values`: given a list of floats, sets the mutable parameters of the state (in the order given by 
get_mutable_parameter_values)

GenerateInput has the following static methods to create an instance of GenerateInput:

* `web_load(args_dict)` receives a json dictionary of "args", which contains "state" and "x"
* `load_from_state_file(filename)` receives the location of a file that contains a serialized parameter tree (state)
* `load_from_FitInput` creates a GenerateInput from a FitInput
* `load_from_PDB` receives the location of a PDB file, and automatically creates the state and x parameters based on the pdb 

FitInput has the following static methods to create an instance of FitInput:

* `web_load(args_dict)` receives a json dictionary of "args", which contains "state" and "x"
* `load_from_state_file(filename)` receives the location of a file that contains a serialized parameter tree (state)

GenerateInput and FitInput can both also be initialized through the initialization function.

```
from dplus.CalculationInput import GenerateInput, FitInput

gen_input=GenerateInput.load_from_state_file('sphere.state')
s=State()
gen_input2=GenerateInput(s)
fit_input=FitInput.load_from_state_file('sphere.state')
```

###State

The state class contains an instance of each of three classes: DomainPreferences, FittingPreferences, and Domain. 
They are described in the upcoming sections.

It has the methods  `get_model`, `get_models_by_type`, `get_mutable_params`,  `get_mutable_parameter_values`, and
`set_mutable_parameter_values`, just as CalculationInput does.

In fact, CalculationInput simply invokes these functions from State when they are called from CalculationInput

State, _and every class and sub class contained within state_ (ie preferences, models, parameters), all have the functions 
`load_from_dictionary` and `serialize`.

`load_from_dictionary` sets the values of the various fields within a class to match those contained within a suitable dictionary. 
It can behave recursively as necessary, for example with a model that has children.

`serialize` saves the contents of a class to a dictionary. Note that there may be additional fields in the dictionary
beyond those described in this document, because some defunct (outdated, irrelevant, or not-yet-implemented) fields are 
still saved in the serialized dictionary.

State also has a function `get_simple_json` which returns a simplified json of just parameters, for Fitting.
(it matches a similarly named function in the c++ code)


####DomainPreferences
The DomainPreferences class contains properties that are copied from the D+ interface. Their usage is explained in 
the D+ documentation.

We create a new instance of DomainPreferences by calling the python __init__ function:

`dom_pref= DomainPreferences()`

There are no arguments given to the initialization function, and all the properties are set to default values:

|Property Name | Default Value | Allowed values|
|---|---|---|
|signal_file|	""|"", or a valid file location|
|convergence|	0.001||
|grid_size|	100|Even integer greater than 20|
|orientation_iterations|	100||
|orientation_method|	"Monte Carlo (Mersenne Twister)"|"Monte Carlo (Mersenne Twister)", "Adaptive (VEGAS) Monte Carlo", "Adaptive Gauss Kronrod"|
|use_grid|	False|True, False|
|q_max|	7.5|Positive number. If signal file is provided, must match highest x value|

Any property can then be changed easily.

`dom_pref.q_max= 10`

If the user tries to set a property to an invalid value (for example, setting q_max to something other than a positive number) they will get an error.

If a signal file is provided, the value of q_max will automatically be set to the highest x value in the signal file.


####Fitting Preferences
The FittingPreferences class contains properties that are copied from the D+ interface. Their usage is explained in the D+ documentation.

We create a new instance of FittingPreferences by calling the python __init__ function:

`fit_pref= FittingPreferences()`

There are no arguments given to the initialization function, and all the properties are set to default values:

|Property Name | Default Value |Allowed Values|Required when|
|---|---|---|---|
|convergence|	0.1| Positive numbers||
|der_eps|	0.1| Positive numbers||
|fitting_iterations|	20|Positive integers||
|step_size|0.01| Positive numbers||
|loss_function|"Trivial Loss"| "Trivial Loss","Huber Loss","Soft L One Loss","Cauchy Loss","Arctan Loss","Tolerant Loss"||
|loss_func_param_one|0.5|Number|Required for all loss_function except "Trivial Loss"|
|loss_func_param_two|0.5|Number|Required when loss_function is "Tolerant Loss"|
|x_ray_residuals_type|"Normal Residuals"|"Normal Residuals","Ratio Residuals","Log Residuals"||
|minimizer_type|"Trust Region"|"Line Search","Trust Region"||
|trust_region_strategy_type|"Dogleg"|"Levenberg-Marquardt","Dogleg"|minimizer_type is "Trust Region"|
|dogleg_type|"Traditional Dogleg"|"Traditional Dogleg","Subspace Dogleg"|trust_region_strategy_type is "Dogleg"|
|line_search_type|"Armijo"|"Armijo","Wolfe"|minimizer_type is "Line Search"|
|line_search_direction_type|"Steepest Descent"|"Steepest Descent","Nonlinear Conjugate Gradient","L-BFGS","BFGS"|minimizer_type is "Line Search". if line_search_type is "Armijo", cannot be "BFGS" or "L-BFGS". |
|nonlinear_conjugate_gradient_type|""|"Fletcher Reeves","Polak Ribirere","Hestenes Stiefel"|linear_search_direction_type is "Nonlinear Conjugate Gradient"|

Any property can then be changed easily.

`fit_pref.convergence= 0.5`

If the user tries to set a property to an invalid value they will get an error.


####Domain

The Domain class describes the parameter tree. 

The root of the tree is the `Domain` class. This contains an array of `Population` classes. 
Each `Population` can contain a number of `Model` classes. Some models have children, also models.

#####Models

Domain and Population are two special kinds of models.

The Domain model is the root of the parameter tree, which can contain multiple populations. 
Populations can contain standard types of models.

The available standard model classes are:

* UniformHollowCylinder
* Sphere
* SymmetricLayeredSlabs
* AsymmetricLayeredSlabs
* Helix
* DiscreteHelix
* SpacefillingSymmetry
* ManualSymmetry
* PDB
* AMP

You can create any model by calling its initialization. 
Please note that models are dynamically loaded from those available in DPlus. 
Therefore, your code editor may underline the model in red.

Models have Layer Parameters, Extra Parameters, and Location Parameters. These are all collection of instances of the `Parameter` class

All of these can be modified. They are accessed using dictionaries.
Example:

```
from dplus.DataModels.models import UniformHollowCylinder

uhc=UniformHollowCylinder()
uhc.layer_params[1]["Radius"].value=2.0
uhc.extra_params["Height"].value=3.0
uhc.location_params["x"].value=2
```

######Parameters

The Parameter class contains the following properties:

value: a float whose default value is 0

sigma: a float whose default value is 0

mutable: a boolean whose default value is False

constraints: an instance of the Constraints class, by default it is the default Constraints

`p=Parameter(4)`

######Constraints

The Constraints class contains the following properties:

MaxValue: a float whose default value is infinity

MinValue: a float whose default value is -infinity

`c=Constraints(min_val=5)`

##CalculationResult

The CalculationResult class is returned by the CalculationRunner. 
The user should generally not be instantiating the class themselves. 

The class has the following properties accessible:

* 'graph': an OrderedDict whose keys are x values and whose values are y values.
* 'y': The raw list of y values from the results json
* 'headers': an OrderDict of headers, whose keys are ModelPtrs and whose values are the header associated. 
This property is not necessarily present in fitting results
* 'parameter_tree': A json of parameters (can be used to create a new state with state's load_from_dictionary). Only present in fitting,
not generate, results
* 'error' : returns the json error report from the dplus run

In addition, Calculation results has the following public functions:

* 'get_amp(model_ptr, destination_folder)': returns the file location of the amplitude file for given model_ptr. 
destination_folder has a default value of None, but if provided, the amplitude file will be copied to that location,
and then have its address returned 
* 'get_pdb(mod_ptr, destination_folder)': returns the file location of the pdb file for given model_ptr. 
destination_folder has a default value of None, but if provided, the pdb file will be copied to that location,
and then have its address returned 
* 'save_to_out_file(path)': receives path (path+filename.out), and saves the results to the file.

###Amplitude

Within the CalculationResult module there is an Amplitude class. It has a static method, 'load', 
which receives a filename and qmax vvalue and creates an instance of the Amplitude class. 

Alternately one can create an empty instance  of Amplitude and then call the function 'read_amp', 
which accomplishes the same thing.  

In addition the class has the following functions:
* q_indices - return array of [q, theta, phi] for each amplitude item in the amplitude array 
* num_indices - return the numbers of  trios [q, theta, phi] in Amplitude file
* complex_amplitude_array - return complex array of amplitudes

All this functions assume that the user call 'load' or 'read_amp'.
In case the user didn't call them, the functions return None arrays/ 0 num of indices

The Amplitude class contains three properties, the poorly named 'arr' (until a better name is suggested) that stores the 
Amplitude values,  'headers', which stores a list of headers and 'step_size - the 'step' between q values.

The function 'save' saved the contents of arr and headers to a new .amp file

```
from dplus.FileReaders import Amplitude

my_amp=Amplitude.load('myamp.amp', 25)
#insert various modifications of amplitude here
my_amp.save('myamp-modified.amp')
```

##Additional Usage examples

From the Dplus GUI it is possible to create a state file by selecting File>Export All Parameters.
Alternately one can create a State by hand, by adding populations, models, fittingpreferences, etc.
In addition, there is an option for generating a pdb, load_from_pdb. It requires the address of the pdb file, and the value of q_max (the largest q value). It automatically populates the rest of the state with reasonable default values.


***Example One***

```
from dplus.CalculationInput import FitInput
from dplus.CalculationRunner import LocalRunner

exe_directory = r"C:\Program Files\D+\bin"
sess_directory = r"session"
runner= LocalRunner(exe_directory, sess_directory)


input=FitInput.load_from_state_file('spherefit.state')
result=runner.fit(input)
print(result.graph)
```

***Example Two***

```
from dplus.CalculationInput import GenerateInput
from dplus.CalculationRunner import LocalRunner
from dplus.DataModels import ModelFactory, Population
from dplus.State import State
from dplus.DataModels.models import UniformHollowCylinder

sess_directory = r"session"
runner= LocalRunner(session_directory=sess_directory)

uhc=UniformHollowCylinder()
s=State()
s.Domain.populations[0].add_model(uhc)

caldata = GenerateInput(s)
result=runner.generate(caldata)
print(result.graph)
```

***Example Three***

```
from dplus.CalculationRunner import LocalRunner
from dplus.CalculationInput import GenerateInput

runner=LocalRunner()
caldata=GenerateInput.load_from_PDB('1JFF.pdb', 5)
result=runner.generate(caldata)
print(result.graph)
```

***Example Four***

```
from dplus.CalculationRunner import LocalRunner
from dplus.CalculationInput import GenerateInput, FitInput
API=LocalRunner()
input = GenerateInput.load_from_state_file("uhc.state")
cylinder = input.get_model("test_cylinder")

print("Original radius is ", cylinder.layer_params[1]['Radius'].value)
result = API.generate(input)

fit_input = FitInput(input.state, result.graph)
cylinder = fit_input.get_model("test_cylinder")
cylinder.layer_params[1]['Radius'].value = 2
cylinder.layer_params[1]['Radius'].mutable = True

fit_result = API.fit(fit_input)
print(fit_result.parameter_tree)
fit_input.combine_results(fit_result)
print("Result radius is ", cylinder.layer_params[1]['Radius'].value)
```

###Python Fitting
It is possible to fit a curve using the results from Generate and numpy's built in minimzation/curve fitting functions. This is a new functionality that is sill very much under development. An example follows:```

```
import numpy as np
from scipy import optimize
from dplus.CalculationInput import GenerateInput, FitInput
from dplus.CalculationRunner import LocalRunner

input=FitInput.load_from_state_file(r"2_pops.state")
generate_runner=LocalRunner()

def run_generate(xdata, *params):
    '''
    scipy's optimization algorithms require a function that receives an x array and an array of parameters, and
    returns a y array.
    this function will be called repeatedly, until scipy's optimization has completed.
    '''
    input.set_mutable_parameter_values(params) #we take the parameters given by scipy and place them inside our parameter tree
    generate_results=generate_runner.generate(input) #call generate
    return np.array(generate_results.y) #return the results of the generate call

x_data=input.x
y_data=input.y
p0 = input.get_mutable_parameter_values()
method='lm' #lenenberg-marquadt (see scipy documentation)
popt, pcov =optimize.curve_fit(run_generate, x_data, y_data, p0=p0, method=method)

#popt is the optimized set of parameters from those we have indicated as mutable
#we can insert them back into our CalculationInput and create the optmized parameter tree
input.set_mutable_parameter_values(popt)
#we can run generate to get the results of generate with them
best_results=generate_runner.generate(input)
```