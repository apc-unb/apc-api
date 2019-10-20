# APC Class

## Get all Classes
* HTTP Request : ```GET http://api.com/class```
* Return a list of object in json format as follow

    ``` 
        [
			{
				"_id"                   :	ObjectId,
				"professorfirstname"    :	String,
				"professorlastname"     :	String,
				"classname"             :	String,
				"address"               :	String,
				"year"                  :	Integer,
				"season"                :	Integer
			}...
		]
    ```

## Create Classes
* HTTP Request : ```POST http://api.com/class```
* Send Classes's data in the request body in the following format 

	``` 
        [
			{
				"professorfirstname"    :	String,
				"professorlastname"     :	String,
				"classname"             :	String,
				"address"               :	String,
				"year"                  :	Integer,
				"season"                :	Integer
			}...
		]

* http StatusCreated (201) will be sent if the class has been created correctly


## Update Classes
* HTTP Request : ```PUT http://api.com/class```
* Send data in the request body in the following format

   ``` 
        [
			{
				"_id"                   :	ObjectId,
				"professorfirstname"    :	String,
				"professorlastname"     :	String,
				"classname"             :	String,
				"address"               :	String,
				"year"                  :	Integer,
				"season"                :	Integer
			}...
		]
    ```

* http StatusCreated (201) will be sent if the student has been updated correctly


## Delete Classes
* HTTP Request : ```DELETE http://api.com/class```
* Send data in the request body in the following format

	``` 
		[
			{  
				"_id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
