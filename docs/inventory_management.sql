-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Aug 05, 2025 at 02:12 PM
-- Server version: 8.0.40-0ubuntu0.20.04.1
-- PHP Version: 8.3.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `inventory_management`
--

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` int NOT NULL,
  `permission` varchar(255) NOT NULL,
  `description` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `permission`, `description`, `created_at`) VALUES
(1, 'user.create', 'Permission to create a new user', '2025-07-27 23:58:05'),
(2, 'user.update', 'Permission to update an existing user', '2025-07-27 23:58:05'),
(3, 'user.fetch', 'Permission to fetch a specific user', '2025-07-27 23:58:05'),
(4, 'user.list', 'Permission to list users', '2025-07-27 23:58:05'),
(5, 'user.delete', 'Permission to delete a user', '2025-07-27 23:58:05'),
(6, 'product.create', 'Permission to create a new product', '2025-07-27 23:58:05'),
(7, 'product.update', 'Permission to update an existing product', '2025-07-27 23:58:05'),
(8, 'product.fetch', 'Permission to fetch a specific product', '2025-07-27 23:58:05'),
(9, 'product.list', 'Permission to list products', '2025-07-27 23:58:05'),
(10, 'product.delete', 'Permission to delete an product', '2025-07-27 23:58:05'),
(11, 'stockHistory.record', 'Permission to record a new stock history', '2025-07-27 23:58:05'),
(12, 'stockHistory.list', 'Permission to list stock histories', '2025-07-27 23:58:05');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `sku` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `live_qty` int NOT NULL DEFAULT '0',
  `warehouse_qty` int NOT NULL DEFAULT '0',
  `threshold_qty` int NOT NULL DEFAULT '0',
  `created_by` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`) VALUES
(1, 'ADMIN'),
(2, 'MANAGER'),
(3, 'CUSTOMER');

-- --------------------------------------------------------

--
-- Table structure for table `role_permissions`
--

CREATE TABLE `role_permissions` (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `role_permissions`
--

INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`) VALUES
(1, 1, '2025-07-27 18:42:10'),
(1, 2, '2025-07-27 18:42:10'),
(1, 3, '2025-07-27 18:42:10'),
(1, 4, '2025-07-27 18:42:10'),
(1, 5, '2025-07-27 18:42:10'),
(1, 6, '2025-07-27 18:42:10'),
(1, 7, '2025-07-27 18:42:10'),
(1, 8, '2025-07-27 18:42:10'),
(1, 9, '2025-07-27 18:42:10'),
(1, 10, '2025-07-27 18:42:10'),
(1, 11, '2025-07-27 23:20:51'),
(1, 12, '2025-07-27 23:20:51'),
(2, 6, '2025-08-05 20:09:31'),
(2, 7, '2025-08-05 20:09:31'),
(2, 8, '2025-08-05 20:09:31'),
(2, 9, '2025-08-05 20:09:31'),
(2, 10, '2025-08-05 20:09:31'),
(2, 11, '2025-08-05 20:09:31'),
(2, 12, '2025-08-05 20:09:31'),
(3, 8, '2025-08-05 20:10:32'),
(3, 9, '2025-08-05 20:10:32');

-- --------------------------------------------------------

--
-- Table structure for table `stock_histories`
--

CREATE TABLE `stock_histories` (
  `id` int NOT NULL,
  `product_id` int NOT NULL,
  `qty` int NOT NULL,
  `operation_type` varchar(255) NOT NULL,
  `operation` varchar(255) NOT NULL,
  `created_by` int DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(100) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `role_id` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `first_name`, `last_name`, `role_id`, `created_at`, `updated_at`) VALUES
(1, 'test@yahoo.com', '$2a$10$LF3in1iy/jKWXATDrezrpefzt89rdevMGanqfkJOGSx59m48koiVG', 'Test', 'User', 1, '2025-08-05 12:11:01', '2025-08-05 18:11:16');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `permission` (`permission`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `sku` (`sku`),
  ADD KEY `fk_products_created_by` (`created_by`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD PRIMARY KEY (`role_id`,`permission_id`),
  ADD KEY `fk_role_permissions_permission_id` (`permission_id`);

--
-- Indexes for table `stock_histories`
--
ALTER TABLE `stock_histories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_stock_histories_product_id` (`product_id`),
  ADD KEY `fk_stock_histories_created_by` (`created_by`) USING BTREE;

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_role_id` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `stock_histories`
--
ALTER TABLE `stock_histories`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_products_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT;

--
-- Constraints for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD CONSTRAINT `fk_role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  ADD CONSTRAINT `fk_role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `stock_histories`
--
ALTER TABLE `stock_histories`
  ADD CONSTRAINT `fk_stock_histories_created_by` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE RESTRICT,
  ADD CONSTRAINT `fk_stock_histories_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `fk_users_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
