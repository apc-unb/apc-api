{
        "ID": "5d386a91ae48f2dab9d14e81",
        "professorfirstname": "Carla",
        "professorlastname": "Denise Castanho",
        "classname": "A",
        "address": "PAT 101",
        "year": 2019,
        "season": 2
    },



# APC Class

## Get all Classes
* HTTP Request : ```GET http://api.com/classes```
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
* HTTP Request : ```POST http://api.com/classes```
* Send Classes's data in the request body in the follow format 

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

* http StatusCreated (201) will be sent if the student has been created correctly


## Update Classes
* HTTP Request : ```PUT http://api.com/classes```
* Send data in the request body in the follow format

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
* HTTP Request : ```DELETE http://api.com/classes```
* Send data in the request body in the follow format

	``` 
		[
			{  
				"_id" : ObjectId
			},...
		]
	```
* http StatusOK (200) will be sent if the students have been deleted correctly
