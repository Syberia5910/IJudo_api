package models

type Config struct {
	LogLevel      	   		string `mapstructure:"LOGLEVEL"`
    Database_File_Path      string `mapstructure:"DB_FILE_PATH"`
    DataBase_File_Name      string `mapstructure:"DB_FILE_NAME"`
	HttpServer_Host      	string `mapstructure:"HTTP_SERVER_ADDRESS"`
	HttpServer_Port	  		int    `mapstructure:"HTTP_SERVER_PORT"`
}