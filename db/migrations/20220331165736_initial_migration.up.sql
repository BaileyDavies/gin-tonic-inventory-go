
CREATE TABLE IF NOT EXISTS gin.Country (
    country_id VARCHAR (50) PRIMARY KEY,
    country_name VARCHAR (50) NOT NULL
);

CREATE TABLE IF NOT EXISTS gin.Brewers (
    brewer_id VARCHAR (50) PRIMARY KEY,
    brewer_name VARCHAR (50) NOT NULL,
    brewer_description VARCHAR (50) NOT NULL,
    brewer_website VARCHAR (50) NULL,
    country_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (country_id)
                     REFERENCES Country (country_id)
);

CREATE TABLE IF NOT EXISTS gin.Gins (
    gin_id VARCHAR (50) PRIMARY KEY,
    gin_name VARCHAR (50) NOT NULL,
    gin_description VARCHAR (50) NOT NULL,
    gin_abv VARCHAR (50) NULL,
    brewer_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (brewer_id)
                  REFERENCES Brewers (brewer_id)
);

CREATE TABLE IF NOT EXISTS gin.Tonics (
    tonic_id VARCHAR (50) PRIMARY KEY,
    tonic_name VARCHAR (50) NOT NULL,
    tonic_description VARCHAR (50) NOT NULL
);

CREATE TABLE IF NOT EXISTS gin.Regions (
    region_id VARCHAR (50) PRIMARY KEY,
    region_name VARCHAR (50) NOT NULL
);


CREATE TABLE IF NOT EXISTS gin.Garnishes (
    garnish_id VARCHAR (50) PRIMARY KEY,
    garnish_name VARCHAR (50) NOT NULL,
    garnish_description VARCHAR (50) NOT NULL,
    garnish_region VARCHAR (50) NULL,
    FOREIGN KEY (garnish_region)
                       REFERENCES Regions(region_id)
);

CREATE TABLE IF NOT EXISTS gin.Gin_Tonics (
    gin_id VARCHAR (50) NOT NULL,
    tonic_id VARCHAR (50) NOT NULL,
   FOREIGN KEY (gin_id)
                        REFERENCES Gins (gin_id),
   FOREIGN KEY (tonic_id)
                        REFERENCES Tonics(tonic_id)
);

CREATE TABLE IF NOT EXISTS gin.Gin_Garnish(
    gin_id VARCHAR (50) NOT NULL,
    garnish_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (gin_id)
                        REFERENCES Gins(gin_id),
    FOREIGN KEY (garnish_id)
                        REFERENCES Garnishes(garnish_id)
);


CREATE TABLE IF NOT EXISTS gin.Users (
    user_id VARCHAR (50) PRIMARY KEY,
    user_name VARCHAR (50) NOT NULL,
    user_email VARCHAR (50) NOT NULL
);

CREATE TABLE IF NOT EXISTS gin.GinRatings(
    rating_id VARCHAR (50) NOT NULL,
    rating_score FLOAT NOT NULL,
    user_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (user_id)
                        REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS gin.TonicRatings(
    rating_id VARCHAR(50) NOT NULL,
    rating_score FLOAT NOT NULL,
    user_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (user_id)
                         REFERENCES Users(user_id)

);

CREATE TABLE IF NOT EXISTS gin.Sellers (
    seller_id VARCHAR(50) PRIMARY KEY,
    seller_name VARCHAR(50) NOT NULL,
    seller_website VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS gin.Gins_Sellers (
    gin_id VARCHAR(50) NOT NULL,
    seller_id VARCHAR(50) NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (gin_id)
                          REFERENCES Gins(gin_id),
    FOREIGN KEY (seller_id)
                          REFERENCES Sellers(seller_id)
);

CREATE TABLE IF NOT EXISTS gin.Tonics_Sellers (
    seller_id VARCHAR(50) NOT NULL,
    tonic_id VARCHAR(50) NOT NULL,
    price FLOAT NOT NULL,
    FOREIGN KEY (seller_id)
                            REFERENCES Sellers(seller_id),
    FOREIGN KEY (tonic_id)
                            REFERENCES Tonics(tonic_id)
);


CREATE TABLE IF NOT EXISTS gin.Users_Gins (
    user_id VARCHAR (50) NOT NULL,
    gin_id VARCHAR (50) NOT NULL,
    list_type list_type NOT NULL,
    FOREIGN KEY (user_id)
                        REFERENCES Users(user_id),
    FOREIGN KEY (gin_id)
                        REFERENCES Gins(gin_id)

);

CREATE TABLE IF NOT EXISTS gin.Users_Tonics (
    user_id VARCHAR (50) NOT NULL,
    tonic_id VARCHAR (50) NOT NULL,
    list_type list_type NOT NULL,
    FOREIGN KEY (user_id)
                        REFERENCES Users(user_id),
    FOREIGN KEY (tonic_id)
                        REFERENCES Tonics(tonic_id)

);

CREATE TABLE IF NOT EXISTS gin.Serves (
    serve_id VARCHAR (50) PRIMARY KEY,
    serve_name VARCHAR(50) NOT NULL,
    gin_id VARCHAR(50) NOT NULL,
    tonic_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (gin_id)
                    REFERENCES Gins(gin_id),
    FOREIGN KEY (tonic_id)
                    REFERENCES Tonics(tonic_id),
    FOREIGN KEY (user_id)
                    REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS gin.ServeRatings (
    rating_id VARCHAR(50) PRIMARY KEY,
    rating_score VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id)
                          REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS gin.Serves_Garnish (
    serve_id VARCHAR (50) NOT NULL,
    tonic_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (serve_id)
                    REFERENCES Serves(serve_id),
    FOREIGN KEY (tonic_id)
                    REFERENCES Tonics(tonic_id)
);

CREATE TABLE IF NOT EXISTS gin.Serves_Users (
    serve_id VARCHAR (50) NOT NULL,
    user_id VARCHAR (50) NOT NULL,
    FOREIGN KEY (serve_id)
                          REFERENCES Serves(serve_id),
    FOREIGN KEY (user_id)
                          REFERENCES Users(user_id)
);



