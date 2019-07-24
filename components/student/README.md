
# Student

## Get all Students
* HTTP Request : ```GET http://api.com/students```
* Return a list of object in json format as follow
    ``` 
        [
			{
				"_id"       : ObjectId,
				"classid" 	: ObjectId,
				"firstname" : String,
				"matricula" : String,
				"password"  : String,
				"handles"   : []String,
				"photourl"  : String,
				"email"     : String,
				"grades" :	{
					"exams"    : []float64,
					"projects" : []float64,
					"lists"    : []float64
				}
			}...
		]
    ```

## Create Students
* HTTP Request : ```POST http://api.com/students```
* Send student's data in the request body in the follow format 
``` 
    [
			{
				"classid" 	: ObjectId,
				"firstname" : String,
				"matricula" : String,
				"password"  : String,
				"handles"   : []String,
				"email"     : String
			},...
	]
```
* http StatusCreated (201) will be sent if the student has been created correctly
    
## Delete Students
* HTTP Request : ```DELETE http://api.com/students```
* Send data in the request body in the follow format
``` 
    [
        {  
            "_id":       ObjectId
        },...
    ]
```
* http StatusOK (200) will be sent if the students have been deleted correctly

## Update Students
* HTTP Request : ```PUT http://api.com/students```
* Send data in the request body in the follow format
``` 
    [
        {  
            "_id"		:   ObjectId,
            "email" 	:	String,
            "password"	:   String,
        },...
    ]
```
* http StatusCreated (201) will be sent if the student has been updated correctly
