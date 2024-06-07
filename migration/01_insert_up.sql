-- Connect to the postgres database (assuming username is "postgres" and password is set to your actual password)
psql -U postgres -h localhost -p 5432

-- Check if connected successfully
\conninfo

-- If connected, proceed with creating the tables

-- Create the Customer table
CREATE TABLE Customer (
  CustomerID SERIAL PRIMARY KEY,
  NIK VARCHAR(20) UNIQUE NOT NULL,
  FullName VARCHAR(100) NOT NULL,
  LegalName VARCHAR(100) NOT NULL,
  BirthPlace VARCHAR(50) NOT NULL,
  BirthDate DATE NOT NULL,
  Salary DECIMAL(15, 2) NOT NULL,
  PhotoKTP TEXT NOT NULL,
  PhotoSelfie TEXT NOT NULL
);

-- Create the CreditLimit table
CREATE TABLE CreditLimit (
  LimitID SERIAL PRIMARY KEY,
  CustomerID INT NOT NULL,
  Tenor INT NOT NULL,
  LimitAmount DECIMAL(15, 2) NOT NULL,
  FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID)
);

-- Create the Transaction table
CREATE TABLE Transaction (
  TransactionID SERIAL PRIMARY KEY,
  CustomerID INT NOT NULL,
  TransactionDate DATE NOT NULL,
  OTR DECIMAL(15, 2) NOT NULL,
  AdminFee DECIMAL(15, 2) NOT NULL,
  InstallmentAmount DECIMAL(15, 2) NOT NULL,
  InterestAmount DECIMAL(15, 2) NOT NULL,
  AssetName VARCHAR(100) NOT NULL,
  FOREIGN KEY (CustomerID) REFERENCES Customer(CustomerID)
);

-- Create the Contract table
CREATE TABLE Contract (
  ContractNumber VARCHAR(20) PRIMARY KEY,
  ContractDate DATE NOT NULL,
  Terms TEXT NOT NULL,
  TransactionID INT NOT NULL,
  Status VARCHAR(50) NOT NULL,
  FOREIGN KEY (TransactionID) REFERENCES Transaction(TransactionID)
);
