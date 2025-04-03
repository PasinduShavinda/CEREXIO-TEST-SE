CREATE DATABASE inventory;

USE inventory;


CREATE TABLE `inventory_master` (
  `product_id` int unsigned NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `stock_in_hand` int DEFAULT NULL,
  `unit_price` decimal(10,0) DEFAULT NULL,
  `supplier_id` int DEFAULT NULL,
  `date_added` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp NULL DEFAULT NULL,
  `remarks` text,
  PRIMARY KEY (`product_id`)
);


CREATE TABLE `inventory_details` (  
  `transaction_id` int unsigned NOT NULL AUTO_INCREMENT,  
  `product_id` int unsigned DEFAULT NULL,  
  `transaction_type` varchar(255) DEFAULT NULL,  
  `quantity` int DEFAULT NULL,  
  `transaction_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,  
  `transaction_description` varchar(255) DEFAULT NULL,  
  `transaction_amount` decimal(10,0) DEFAULT NULL,  
  `transaction_status` varchar(255) DEFAULT NULL,  
  `transaction_notes` text COLLATE utf8mb4_general_ci,  
  PRIMARY KEY (`transaction_id`),  
  CONSTRAINT `fk_inventory_product`  
    FOREIGN KEY (`product_id`)  
    REFERENCES `inventory_master` (`product_id`)  
);
