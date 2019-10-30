# Student

## Create Students
* HTTP Request : ```POST http://api.com/student```
* Send student's data in the request body in the following format 

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
* HTTP Request : ```POST http://api.com/student/file```
* Send data in the request body in the following format

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
* HTTP Request : ```GET http://api.com/student```
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
                "uri" 	 :	String
            },
            "photourl"  :	String,
            "email"     :	String,
            "grades"    :	{
                "exams"     :	[]float64,
                "projects"  :	[]float64,
                "lists"     :	[]float64
            }
        },...
    ]
    ```

## Get all Students from a class
* HTTP Request : ```GET http://api.com/student/{classid}```
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
                "uri" 	 :	String
            },
            "photourl"  :	String,
            "email"     :	String,
            "grades"    :	{
                "exams"    :	[]float64,
                "projects" :	[]float64,
                "lists"    :	[]float64
            }
        }...
    ]
    ```

## Update Students
* HTTP Request : ```PUT http://api.com/student```
* Send data in the request body in the following format
* PS : Student can only uptade empty handles

	``` 
    [
        {  
            "_id".      	:   ObjectId,
            "email" 	:   String,
            "password"	:   String,
            "newpassword":   String,
            "handles"   	: 	{
                    "codeforces" :	String,
                    "uri" 	:	String
            }
        },...
    ]
	```
* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete Students
* HTTP Request : ```DELETE http://api.com/student```
* Send data in the request body in the following format
* *PS* : Only admin with `Professor` : `True` can make this request

	``` 
    [
        {  
            "_id" : ObjectId
        },...
    ]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly


## Log in Student
* HTTP Request : ```POST http://api.com/student/login```

    ``` 
        {
            "matricula" :	String,
            "password"  :	String,
        }
    ```
* Return a json format as follow

	```
    {
        "userexist"     :	Boolean,
        "student"       :	StudentInfo,
        "class"	       :	SchoolClass,
        "news"	       :	[]News,
        "Progress": {
            "done"  : String,
            "total" : String
        },
    }
	```
 
 ## Get Student Codeforces Progress
 * HTTP Request : ```GET /student/contest/{studentid}```
 
 * Return a json format as follow
 
 	```
    [
        {
            "name": String,
            "url": String
            "done": String,
            "total": String,
        }...,
    ]
 	```
