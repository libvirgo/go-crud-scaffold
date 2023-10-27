-- Create "admin_users" table
CREATE TABLE `admin_users` (`id` bigint NOT NULL AUTO_INCREMENT, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `wallet_address` varchar(255) NOT NULL, PRIMARY KEY (`id`), INDEX `adminuser_wallet_address` (`wallet_address`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `wallet_address` varchar(255) NOT NULL, PRIMARY KEY (`id`), INDEX `user_wallet_address` (`wallet_address`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_activities" table
CREATE TABLE `user_activities` (`id` bigint NOT NULL AUTO_INCREMENT, `create_time` timestamp NOT NULL, `update_time` timestamp NOT NULL, `type` bigint NOT NULL, `user_user_activity` bigint NULL, PRIMARY KEY (`id`), INDEX `user_activities_users_user_activity` (`user_user_activity`), CONSTRAINT `user_activities_users_user_activity` FOREIGN KEY (`user_user_activity`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
