
# Exam

## Get all News
* HTTP Request : ```GET http://api.com/news```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"           :	ObjectId,
				"classid"       :	ObjectId,
				"title"         :	String,
				"description"   :	String
                "tags"          :   []String
			}...
		]
    ```


## Get all News from a class
* HTTP Request : ```GET http://api.com/news/{classID}```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"           :	ObjectId,
				"classid"       :	ObjectId,
				"title"         :	String,
				"description"   :	String
                "tags"          :   []String
			}...
		]
    ```


## Create News
* HTTP Request : ```POST http://api.com/news```
* Send News's data in the request body in the following format 

	``` 
        [
			{
				"classid"       :	ObjectId,
				"title"         :	String,
				"description"   :	String
                "tags"          :   []String
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been created correctly


## Update News
* HTTP Request : ```PUT http://api.com/news```
* Send data in the request body in the following format

    ``` 
        [
			{
				"_id"           :	ObjectId,
				"classid"       :	ObjectId,
				"title"         :	String,
				"description"   :	String
                "tags"          :   []String
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete News
* HTTP Request : ```DELETE http://api.com/news```
* Send data in the request body in the following format

	``` 
		[
			{  
				"_id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
