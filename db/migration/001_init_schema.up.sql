CREATE TABLE `user` (
                        `id` int PRIMARY KEY AUTO_INCREMENT,
                        `full_name` varchar(255),
                        `email` varchar(255),
                        `tel` varchar(255),
                        `password_hash` varchar(255),
                        `otp` varchar(255),
                        `otp_timeout` datetime,
                        `is_otp_active` bit,
                        `avatar_path` varchar(255),
                        `update_reason` varchar(255),
                        `status` tinyint,
                        `created_at` datetime
);

CREATE TABLE `organization` (
                                `id` int PRIMARY KEY AUTO_INCREMENT,
                                `name` varchar(255),
                                `info` tinytext,
                                `root_organization` int,
                                `type` varchar(255)
);

CREATE TABLE `user_organization` (
                                     `id` int PRIMARY KEY AUTO_INCREMENT,
                                     `user_id` int,
                                     `organization_id` int,
                                     `created_at` datetime,
                                     `update_at` datetime,
                                     `is_active` bit
);

CREATE TABLE `role` (
                        `id` int PRIMARY KEY AUTO_INCREMENT,
                        `name` varchar(255),
                        `description` varchar(255)
);

CREATE TABLE `user_role` (
                             `id` int PRIMARY KEY AUTO_INCREMENT,
                             `user_id` int,
                             `role_id` int
);

CREATE TABLE `permission` (
                              `id` int PRIMARY KEY AUTO_INCREMENT,
                              `name` varchar(255),
                              `description` varchar(255)
);

CREATE TABLE `role_permission` (
                                   `id` int PRIMARY KEY AUTO_INCREMENT,
                                   `permission_id` int,
                                   `role_id` int
);

ALTER TABLE `organization` ADD FOREIGN KEY (`root_organization`) REFERENCES `organization` (`id`);

ALTER TABLE `user_organization` ADD FOREIGN KEY (`organization_id`) REFERENCES `user_organization` (`id`);

ALTER TABLE `user_organization` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `user_organization` ADD FOREIGN KEY (`organization_id`) REFERENCES `organization` (`id`);

ALTER TABLE `user_role` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `user_role` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`);

ALTER TABLE `role_permission` ADD FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);