# Student

## Create Exams
* HTTP Request : ```POST http://api.com/exams```
* Send exams's data in the request body in the follow format 

	``` 
		[
			{
				"classid"	:	ObjectId,
				"title"     	:	String
			},...
		]
	```
* http StatusCreated (201) will be sent if the exam has been created correctly

## Get all Exams
* HTTP Request : ```GET http://api.com/exams```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       	:	ObjectId,
				"classid"   	:	ObjectId,
				"title" 	:	String
			},...
		]
    ```

## Get all Exams from a specific class
* HTTP Request : ```GET http://api.com/exams/{classid}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       	:	ObjectId,
				"classid"	:	ObjectId,
				"title" 	:	String
			},...
		]
    ```

## Update Students
* HTTP Request : ```PUT http://api.com/exams```
* Send data in the request body in the follow format
	``` 
		[
			{  
				"_id".      	:   	ObjectId,
				"classid"   	:	ObjectId,
				"title" 	:	String
			},...
		]
	```
* http StatusCreated (201) will be sent if the exams has been updated correctly


## Delete Exams
* HTTP Request : ```DELETE http://api.com/exams```
* Send data in the request body in the follow format

	``` 
		[
			{  
				"_id"	:	ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the exams have been deleted correctly

