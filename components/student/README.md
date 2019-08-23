# Student

## Create Students
* HTTP Request : ```POST http://api.com/students```
* Send student's data in the request body in the follow format 

	``` 
		[
			{
				"classid"   :	ObjectId,
				"firstname" :	String,
				"lastname"  :	String,
				"matricula" :	String
			},...
		]
	```
* http StatusCreated (201) will be sent if the student has been created correctly

## Create Students by CSV file
* HTTP Request : ```PUT http://api.com/studentsFile```
* Send data in the request body in the follow format

	|    ANO/SEMESTE/TURMA   |             2019/2/A 
	|------------------------|-------------------------------
	|       160140000        | 	Thiago Veras Machado    
	|       160140000        | 	Giovanni Guidini       
	|       160140000        | 	Vitor Dullens     

	``` 
		{  
			"file"	: file.csv,
		},...
	```

* http StatusCreated (201) will be sent if the student has been created correctly

## Get all Students
* HTTP Request : ```GET http://api.com/students```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"firstname" :	String,
				"lastname"  :	String,
				"matricula" :	String,
				"password"  :	String,
				"handles"   :	{
					"codeforces" :	String,
					"uri" 	     :	String
				},
				"photourl"  :	String,
				"email"     :	String,
				"grades"    :	{
					"exams"    :	[]float64,
					"projects" :	[]float64,
					"lists"    :	[]float64
				}
			},...
		]
    ```

## Get all Students from a specific class
* HTTP Request : ```GET http://api.com/students/{classid}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"firstname" :	String,
				"lastname"  :	String,
				"matricula" :	String,
				"password"  :	String,
				"handles"   :	{
					"codeforces" :	String,
					"uri" 	     :	String
				},
				"photourl"  :	String,
				"email"     :	String,
				"grades"    :	{
					"exams"    :	[]float64,
					"projects" :	[]float64,
					"lists"    :	[]float64
				}
			},...
		]
    ```

## Update Students
* HTTP Request : ```PUT http://api.com/students```
* Send data in the request body in the follow format
* PS : Student can only uptade empty handles

	``` 
		[
			{  
				"_id".      	:   ObjectId,
				"email" 	:   String,
				"password"	:   String,
				"newpassword" 	: 	String,
				"handles"   	: 	{
						"codeforces" :	String,
						"uri" 	     :	String
				}
			},...
		]
	```
* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete Students
* HTTP Request : ```DELETE http://api.com/students```
* Send data in the request body in the follow format

	``` 
		[
			{  
				"_id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly


## Log in Students
* HTTP Request : ```POST http://api.com/student```

    ``` 
		{
			"matricula" :	String,
			"password"  :	String,
		}
    ```
* Return a json format as follow

	```
	{
	    "userexist"	:	Boolean,
	    "student"	:	StudentInfo,
	    "class"	:	SchoolClass,
	    "news"	:	[]News 
	}
	```
