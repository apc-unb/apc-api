
# Exam

## Get all Exams
* HTTP Request : ```GET http://api.com/exam```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"title"     :	String
			}...
		]
    ```


## Get all Exams from a class
* HTTP Request : ```GET http://api.com/exam/{classid}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"title"     :	String
			}...
		]
    ```


## Create Exams
* HTTP Request : ```POST http://api.com/exam```
* Send exams's data in the request body in the following format 

	``` 
        [
			{
				"classid"   :	ObjectId,
				"title"     :	String
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been created correctly


## Update Exams
* HTTP Request : ```PUT http://api.com/exam```
* Send data in the request body in the following format

    ``` 
        [
			{
				"_id"       :	ObjectId,
				"classid"   :	ObjectId,
				"title"     :	String
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete Exams
* HTTP Request : ```DELETE http://api.com/exam```
* Send data in the request body in the following format

	``` 
		[
			{  
				"_id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
