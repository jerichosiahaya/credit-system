-- Active: 1717488635357@@127.0.0.1@5432@postgres@public
INSERT INTO Customer (NIK, FullName, LegalName, BirthPlace, BirthDate, Salary, PhotoKTP, PhotoSelfie)
VALUES 
('1234567890123456', 'Budi', 'Budi', 'Jakarta', '1985-05-15', 5000000.00, 'base64photoktp1', 'base64photoselfie1'),
('2345678901234567', 'Annisa', 'Annisa', 'Bandung', '1990-08-22', 10000000.00, 'base64photoktp2', 'base64photoselfie2');

-- Insert CreditLimit data
INSERT INTO CreditLimit (CustomerID, Tenor, LimitAmount)
VALUES 
-- Budi and Annisa's limits
(1, 1, 100000.00),
(1, 2, 200000.00),
(1, 3, 500000.00),
(1, 4, 700000.00),
(2, 1, 1000000.00),
(2, 2, 1200000.00),
(2, 3, 1500000.00),
(2, 4, 2000000.00);