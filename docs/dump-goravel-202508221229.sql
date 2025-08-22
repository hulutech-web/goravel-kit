-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: goravel
-- ------------------------------------------------------
-- Server version	8.0.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `contract_templates`
--

DROP TABLE IF EXISTS `contract_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contract_templates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contract_templates`
--

LOCK TABLES `contract_templates` WRITE;
/*!40000 ALTER TABLE `contract_templates` DISABLE KEYS */;
/*!40000 ALTER TABLE `contract_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contracts`
--

DROP TABLE IF EXISTS `contracts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contracts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `landlord_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `tenant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `type` enum('代理合同','租赁合同') NOT NULL DEFAULT '代理合同',
  `content` text NOT NULL,
  `tenant_sign` varchar(255) NOT NULL,
  `landlord_sign` varchar(255) NOT NULL,
  `signed_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `signed_location` varchar(255) NOT NULL COMMENT '签约地点',
  `paper_contract` varchar(255) DEFAULT NULL COMMENT '纸质合同',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contracts`
--

LOCK TABLES `contracts` WRITE;
/*!40000 ALTER TABLE `contracts` DISABLE KEYS */;
/*!40000 ALTER TABLE `contracts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `failed_jobs`
--

DROP TABLE IF EXISTS `failed_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `failed_jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL,
  `connection` text NOT NULL,
  `queue` text NOT NULL,
  `payload` longtext NOT NULL,
  `exception` longtext NOT NULL,
  `failed_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `failed_jobs`
--

LOCK TABLES `failed_jobs` WRITE;
/*!40000 ALTER TABLE `failed_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `failed_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `file_cates`
--

DROP TABLE IF EXISTS `file_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `file_cates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '类目名称',
  `sort` int(11) NOT NULL COMMENT '排序',
  `type` varchar(255) NOT NULL COMMENT '类目类型: [image=图片, video=视频, audio=音频, file=文件]',
  `pid` bigint(20) NOT NULL COMMENT '父类目ID',
  `tenant_id` bigint(20) NOT NULL COMMENT '租户ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `file_cates`
--

LOCK TABLES `file_cates` WRITE;
/*!40000 ALTER TABLE `file_cates` DISABLE KEYS */;
/*!40000 ALTER TABLE `file_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `files`
--

DROP TABLE IF EXISTS `files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cid` int(11) NOT NULL COMMENT '类目ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `type` varchar(255) NOT NULL COMMENT '文件类型: [image=图片, video=视频, audio=音频, file=文件]',
  `name` varchar(255) NOT NULL COMMENT '文件名称',
  `uri` varchar(255) NOT NULL COMMENT '文件路径',
  `ext` varchar(255) NOT NULL COMMENT '文件扩展',
  `size` int(11) NOT NULL COMMENT '文件大小',
  `engine` varchar(255) NOT NULL COMMENT '存储引擎',
  `path` varchar(255) NOT NULL COMMENT '访问路径',
  `tenant_id` int(11) NOT NULL COMMENT '租户ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `files`
--

LOCK TABLES `files` WRITE;
/*!40000 ALTER TABLE `files` DISABLE KEYS */;
/*!40000 ALTER TABLE `files` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `houses`
--

DROP TABLE IF EXISTS `houses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `houses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `landlord_id` bigint(20) unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `monthly_rent` double NOT NULL,
  `deposit` double NOT NULL,
  `header_img` varchar(255) NOT NULL,
  `poster` varchar(255) NOT NULL,
  `albums` text NOT NULL,
  `location` geometry DEFAULT NULL COMMENT '经纬度',
  `area` double NOT NULL,
  `facilities` text NOT NULL,
  `property_fee` double NOT NULL,
  `traffic` text NOT NULL,
  `shopping` text NOT NULL,
  `video` text NOT NULL,
  `status` varchar(255) NOT NULL,
  `swipers` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `houses_landlord_id_index` (`landlord_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `houses`
--

LOCK TABLES `houses` WRITE;
/*!40000 ALTER TABLE `houses` DISABLE KEYS */;
/*!40000 ALTER TABLE `houses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `queue` varchar(255) NOT NULL,
  `payload` longtext NOT NULL,
  `attempts` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `reserved_at` datetime DEFAULT NULL,
  `available_at` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `jobs_queue_index` (`queue`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menus`
--

DROP TABLE IF EXISTS `menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级ID',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `path` varchar(255) DEFAULT NULL COMMENT '路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件',
  `icon` varchar(255) NOT NULL DEFAULT 'AlertOutlined' COMMENT '图标',
  `menu_type` varchar(255) NOT NULL COMMENT '菜单类型: [page=页面, action=操作]',
  `cacheable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否缓存',
  `render_menu` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否渲染菜单',
  `permission` varchar(255) DEFAULT NULL COMMENT '权限标识',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `target` varchar(255) DEFAULT NULL COMMENT '目标',
  `badge` varchar(255) DEFAULT NULL COMMENT '角标',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menus`
--

LOCK TABLES `menus` WRITE;
/*!40000 ALTER TABLE `menus` DISABLE KEYS */;
INSERT INTO `menus` VALUES (1,0,'首页','workplace','/workplace','@/pages/workplace/index.vue','HomeOutlined','',0,1,'workplace',0,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(2,0,'系统','system','/system','@/components/layout/BlankView.vue','ControlOutlined','',0,1,'system',100,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(3,2,'菜单管理','system.auth.menu','/system/auth/menu','@/pages/system/auth/menu/index.vue','SettingOutlined','',0,1,'system:auth:menu',100,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(4,2,'角色管理','system.auth.role','/system/auth/role','@/pages/system/auth/role/index.vue','UserSwitchOutlined','',0,1,'system:auth:role',100,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(5,2,'权限管理','permission','/system/auth/permission','@/pages/system/auth/permission/index.vue','VerifiedOutlined','',0,1,'system:auth:permission',100,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(6,2,'用户管理','user','/system/user','@/pages/system/user/index.vue','UserOutlined','',0,1,'system:user',1,'','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(7,2,'附件中心','netdisk','/system/netdisk','@/pages/system/netdisk/index.vue','FolderOutlined','page',1,1,'system:netdisk',0,'_self','','2025-08-22 00:53:31','2025-08-22 00:53:31'),(8,2,'代码生成','system.crud.index','/system/crud/index','@/pages/system/crud/index.vue','ApiOutlined','page',1,1,'system:crud:index',0,'_self','',NULL,'2025-08-22 01:08:23'),(9,2,'数据表字段','system.crud.column','/system/crud/:id/column','@/pages/system/crud/column.vue','DatabaseOutlined','page',1,0,'system:crud:column',0,'_self','',NULL,NULL);
/*!40000 ALTER TABLE `menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) NOT NULL,
  `batch` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (16,'20210101000001_create_users_table',1),(17,'20210101000002_create_jobs_table',1),(18,'20250812104711_create_houses_table',1),(19,'20250812122145_create_file_cates_table',1),(20,'20250812122233_create_files_table',1),(21,'20250812122424_create_menus_table',1),(22,'20250812180742_create_roles_table',1),(23,'20250812180749_create_permissions_table',1),(24,'20250812181002_create_role_permissions_table',1),(25,'20250812181028_create_user_roles_table',1),(26,'20250815095805_create_rooms_table',1),(27,'20250815100241_create_orders_table',1),(28,'20250815101623_create_contracts_table',1),(29,'20250815101745_create_contract_templates_table',1),(30,'20250816105218_create_pdf_gens_table',1);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `target_type` enum('house','room') NOT NULL COMMENT '目标类型',
  `house_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `room_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `tenant_id` bigint(20) unsigned NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `monthly_rent` double NOT NULL DEFAULT '0',
  `status` enum('待支付','已支付','执行中','已结束') NOT NULL DEFAULT '待支付',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pdf_gens`
--

DROP TABLE IF EXISTS `pdf_gens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pdf_gens` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `params` json NOT NULL,
  `html` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pdf_gens`
--

LOCK TABLES `pdf_gens` WRITE;
/*!40000 ALTER TABLE `pdf_gens` DISABLE KEYS */;
INSERT INTO `pdf_gens` VALUES (1,'pdf_gen2','{\"type\": \"object\", \"title\": \"销售合同\", \"$schema\": \"http://json-schema.org/draft-07/schema#\", \"required\": [\"contractTitle\", \"contractNo\", \"sellerName\", \"buyerName\", \"productName\", \"quantity\", \"unitPrice\"], \"properties\": {\"quantity\": {\"type\": \"number\", \"title\": \"数量\", \"minimum\": 1}, \"signDate\": {\"type\": \"string\", \"title\": \"签署日期\", \"format\": \"date\"}, \"buyerName\": {\"type\": \"string\", \"title\": \"买方名称\"}, \"unitPrice\": {\"type\": \"number\", \"title\": \"单价\", \"minimum\": 0}, \"contractNo\": {\"type\": \"string\", \"title\": \"合同编号\", \"pattern\": \"^CT-[0-9]{8}$\"}, \"sellerName\": {\"type\": \"string\", \"title\": \"卖方名称\"}, \"productName\": {\"type\": \"string\", \"title\": \"产品名称\"}, \"totalAmount\": {\"type\": \"number\", \"title\": \"总金额\", \"readOnly\": true}, \"deliveryDate\": {\"type\": \"string\", \"title\": \"交货日期\", \"format\": \"date\"}, \"paymentTerms\": {\"enum\": [\"预付全款\", \"货到付款\", \"30天账期\", \"60天账期\"], \"type\": \"string\", \"title\": \"付款条款\"}, \"specialTerms\": {\"rows\": 4, \"type\": \"string\", \"title\": \"特别条款\", \"format\": \"textarea\"}, \"contractTitle\": {\"type\": \"string\", \"title\": \"合同标题\", \"default\": \"产品销售合同\"}, \"warrantyPeriod\": {\"enum\": [\"6个月\", \"1年\", \"2年\", \"3年\", \"5年\"], \"type\": \"string\", \"title\": \"保修期\"}}, \"description\": \"销售合同表单数据\"}','<!DOCTYPE html>\n<html lang=\"zh-CN\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <title>合同PDF动态生成系统</title>\n  <script src=\"https://unpkg.com/vue@3/dist/vue.global.js\"></script>\n  <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css\">\n  <style>\n    * {\n      margin: 0;\n      padding: 0;\n      box-sizing: border-box;\n      font-family: \'Segoe UI\', Tahoma, Geneva, Verdana, sans-serif;\n    }\n    \n    body {\n      background: linear-gradient(135deg, #1a2a6c, #b21f1f, #1a2a6c);\n      min-height: 100vh;\n      padding: 20px;\n      color: #333;\n    }\n    \n    .container {\n      max-width: 1800px;\n      margin: 0 auto;\n    }\n    \n    header {\n      text-align: center;\n      padding: 20px 0;\n      color: white;\n      margin-bottom: 30px;\n    }\n    \n    header h1 {\n      font-size: 2.8rem;\n      margin-bottom: 10px;\n      text-shadow: 0 2px 8px rgba(0,0,0,0.3);\n    }\n    \n    header p {\n      font-size: 1.3rem;\n      max-width: 900px;\n      margin: 0 auto;\n      opacity: 0.9;\n      line-height: 1.6;\n    }\n    \n    .tabs {\n      display: flex;\n      background: rgba(255, 255, 255, 0.1);\n      border-radius: 10px;\n      padding: 10px;\n      margin-bottom: 25px;\n      max-width: 1000px;\n      margin: 0 auto 25px;\n      flex-wrap: wrap;\n    }\n    \n    .tab {\n      flex: 1;\n      padding: 15px;\n      text-align: center;\n      color: white;\n      font-size: 1.1rem;\n      font-weight: 600;\n      cursor: pointer;\n      border-radius: 8px;\n      transition: all 0.3s ease;\n      min-width: 200px;\n    }\n    \n    .tab.active {\n      background: rgba(255, 255, 255, 0.2);\n      box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);\n    }\n    \n    .tab i {\n      margin-right: 10px;\n      font-size: 1.2rem;\n    }\n    \n    .app-container {\n      display: flex;\n      gap: 25px;\n      background: rgba(255, 255, 255, 0.95);\n      border-radius: 15px;\n      box-shadow: 0 15px 50px rgba(0, 0, 0, 0.3);\n      overflow: hidden;\n      min-height: 700px;\n    }\n    \n    .panel {\n      padding: 25px;\n      flex: 1;\n      display: flex;\n      flex-direction: column;\n      min-width: 0; /* 修复宽度问题 */\n    }\n    \n    .schema-editor {\n      background: #f8f9fa;\n      border-right: 1px solid #e0e0e0;\n    }\n    \n    .form-preview {\n      background: white;\n    }\n    \n    .panel-title {\n      display: flex;\n      align-items: center;\n      gap: 10px;\n      margin-bottom: 20px;\n      padding-bottom: 15px;\n      border-bottom: 2px solid #2c3e50;\n      color: #2c3e50;\n    }\n    \n    .panel-title i {\n      font-size: 1.6rem;\n    }\n    \n    .panel-title h2 {\n      font-size: 1.8rem;\n      font-weight: 700;\n    }\n    \n    textarea {\n      width: 100%;\n      height: 300px;\n      padding: 15px;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      font-family: monospace;\n      font-size: 15px;\n      resize: none;\n      transition: all 0.3s;\n    }\n    \n    textarea:focus {\n      outline: none;\n      border-color: #3498db;\n      box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);\n    }\n    \n    .actions {\n      display: flex;\n      gap: 15px;\n      margin-top: 20px;\n      flex-wrap: wrap;\n    }\n    \n    button {\n      padding: 14px 20px;\n      color: white;\n      border: none;\n      border-radius: 8px;\n      font-size: 1.1rem;\n      font-weight: 600;\n      cursor: pointer;\n      transition: all 0.3s;\n      display: flex;\n      align-items: center;\n      justify-content: center;\n      gap: 10px;\n      flex: 1;\n      min-width: 200px;\n    }\n    \n    button.primary {\n      background: linear-gradient(to right, #3498db, #2980b9);\n    }\n    \n    button.secondary {\n      background: linear-gradient(to right, #7f8c8d, #95a5a6);\n    }\n    \n    button.success {\n      background: linear-gradient(to right, #27ae60, #2ecc71);\n    }\n    \n    button.danger {\n      background: linear-gradient(to right, #e74c3c, #c0392b);\n    }\n    \n    button:hover {\n      transform: translateY(-3px);\n      box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);\n    }\n    \n    .form-container {\n      flex: 1;\n      overflow-y: auto;\n      padding: 10px;\n    }\n    \n    .form-group {\n      margin-bottom: 25px;\n      padding: 20px;\n      background: #f8f9fa;\n      border-radius: 10px;\n      box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);\n      position: relative;\n    }\n    \n    .form-group h3 {\n      margin-bottom: 15px;\n      color: #2c3e50;\n      font-size: 1.4rem;\n      padding-bottom: 8px;\n      border-bottom: 1px solid #eee;\n    }\n    \n    label {\n      display: block;\n      margin-bottom: 8px;\n      font-weight: 600;\n      color: #2c3e50;\n    }\n    \n    input, select, textarea {\n      width: 100%;\n      padding: 14px;\n      border: 1px solid #ced4da;\n      border-radius: 8px;\n      font-size: 1rem;\n      transition: all 0.3s;\n      background: white;\n    }\n    \n    input:focus, select:focus, textarea:focus {\n      outline: none;\n      border-color: #3498db;\n      box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);\n    }\n    \n    .error {\n      color: #e74c3c;\n      font-size: 0.9rem;\n      margin-top: 5px;\n    }\n    \n    .form-container:empty::before {\n      content: \"请提供有效的JSON Schema来生成表单\";\n      display: block;\n      text-align: center;\n      padding: 40px;\n      color: #7f8c8d;\n      font-size: 1.2rem;\n    }\n    \n    .data-preview {\n      background: #2c3e50;\n      color: white;\n      padding: 25px;\n      border-radius: 10px;\n      margin-top: 20px;\n      font-family: monospace;\n      white-space: pre-wrap;\n      max-height: 250px;\n      overflow-y: auto;\n      line-height: 1.5;\n    }\n    \n    .data-preview-title {\n      display: flex;\n      align-items: center;\n      gap: 10px;\n      margin-bottom: 15px;\n      color: #3498db;\n    }\n    \n    .status {\n      margin-top: 15px;\n      padding: 15px;\n      border-radius: 8px;\n      text-align: center;\n      font-weight: 500;\n      font-size: 1.1rem;\n    }\n    \n    .success {\n      background: #d4edda;\n      color: #155724;\n      border-left: 4px solid #28a745;\n    }\n    \n    .error-status {\n      background: #f8d7da;\n      color: #721c24;\n      border-left: 4px solid #dc3545;\n    }\n    \n    .info {\n      background: #cce5ff;\n      color: #004085;\n      border-left: 4px solid #007bff;\n    }\n    \n    .template-container {\n      display: flex;\n      flex-direction: column;\n      height: 100%;\n    }\n    \n    .editor-container {\n      flex: 1;\n      display: flex;\n      flex-direction: column;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      overflow: hidden;\n    }\n    \n    .editor-header {\n      background: #2c3e50;\n      color: white;\n      padding: 10px 15px;\n      font-weight: 600;\n      display: flex;\n      justify-content: space-between;\n      flex-wrap: wrap;\n    }\n    \n    #templateEditor {\n      height: 400px;\n      width: 100%;\n      border: none;\n      background: #1e1e1e;\n      color: #d4d4d4;\n      padding: 15px;\n      font-family: monospace;\n      font-size: 14px;\n      line-height: 1.5;\n      resize: vertical;\n    }\n    \n    .contract-preview {\n      flex: 1;\n      background: white;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      padding: 20px;\n      overflow-y: auto;\n      margin-top: 20px;\n      font-family: \'Times New Roman\', Times, serif;\n      line-height: 1.8;\n    }\n    \n    .contract-preview h1 {\n      text-align: center;\n      margin-bottom: 30px;\n      color: #2c3e50;\n      border-bottom: 2px solid #2c3e50;\n      padding-bottom: 15px;\n    }\n    \n    .contract-section {\n      margin-bottom: 30px;\n    }\n    \n    .contract-section h2 {\n      color: #2c3e50;\n      margin-bottom: 15px;\n      border-left: 4px solid #3498db;\n      padding-left: 10px;\n    }\n    \n    .contract-variable {\n      background: #fff8e1;\n      padding: 2px 5px;\n      border-radius: 4px;\n      border: 1px dashed #ffc107;\n    }\n    \n    .signature-area {\n      display: flex;\n      justify-content: space-between;\n      margin-top: 60px;\n      padding-top: 20px;\n      border-top: 1px solid #ccc;\n      flex-wrap: wrap;\n    }\n    \n    .signature-box {\n      width: 45%;\n      text-align: center;\n      min-width: 200px;\n      margin-bottom: 20px;\n    }\n    \n    .signature-line {\n      height: 1px;\n      background: #333;\n      margin: 40px 0 10px;\n    }\n    \n    .template-actions {\n      display: flex;\n      gap: 15px;\n      margin-top: 20px;\n      flex-wrap: wrap;\n    }\n    \n    .template-actions button {\n      flex: none;\n      width: auto;\n      padding: 12px 25px;\n    }\n    \n    .tab-content {\n      display: none;\n    }\n    \n    .tab-content.active {\n      display: block;\n    }\n    \n    .pdf-preview {\n      width: 100%;\n      height: 600px;\n      background: #f0f0f0;\n      border-radius: 8px;\n      display: flex;\n      align-items: center;\n      justify-content: center;\n      margin-top: 20px;\n      flex-direction: column;\n    }\n    \n    .pdf-preview img {\n      max-width: 100%;\n      max-height: 500px;\n      box-shadow: 0 5px 15px rgba(0,0,0,0.2);\n    }\n    \n    .pdf-preview p {\n      margin-top: 20px;\n      color: #7f8c8d;\n      font-size: 1.2rem;\n    }\n    \n    .download-btn {\n      margin-top: 25px;\n      padding: 15px 40px;\n      font-size: 1.2rem;\n    }\n    \n    .template-syntax {\n      margin-top: 10px;\n      padding: 15px;\n      background: #f8f9fa;\n      border-radius: 8px;\n      font-size: 0.9rem;\n    }\n    \n    .template-syntax h4 {\n      margin-bottom: 10px;\n      color: #2c3e50;\n    }\n    \n    .syntax-example {\n      background: #2c3e50;\n      color: white;\n      padding: 10px;\n      border-radius: 4px;\n      font-family: monospace;\n      margin-top: 5px;\n    }\n    \n    @media (max-width: 1200px) {\n      .app-container {\n        flex-direction: column;\n      }\n      \n      .schema-editor {\n        border-right: none;\n        border-bottom: 1px solid #eee;\n      }\n      \n      header h1 {\n        font-size: 2.2rem;\n      }\n      \n      .tab {\n        min-width: 150px;\n        padding: 10px;\n        font-size: 1rem;\n      }\n    }\n    \n    @media (max-width: 768px) {\n      .tabs {\n        flex-direction: column;\n      }\n      \n      .signature-box {\n        width: 100%;\n      }\n      \n      .panel-title h2 {\n        font-size: 1.5rem;\n      }\n      \n      button {\n        min-width: 100%;\n      }\n    }\n  </style>\n</head>\n<body>\n  <div id=\"app\">\n    <div class=\"container\">\n      <header>\n        <h1><i class=\"fas fa-file-contract\"></i> 合同PDF动态生成系统</h1>\n        <p>使用JSON Schema设计表单，编辑合同模板，动态生成PDF文件</p>\n      </header>\n      \n      <div class=\"tabs\">\n        <div class=\"tab\" @click=\"currentTab = \'design\'\" :class=\"{active: currentTab === \'design\'}\">\n          <i class=\"fas fa-drafting-compass\"></i> 表单设计\n        </div>\n        <div class=\"tab\" @click=\"currentTab = \'template\'\" :class=\"{active: currentTab === \'template\'}\">\n          <i class=\"fas fa-file-alt\"></i> 模板编辑\n        </div>\n        <div class=\"tab\" @click=\"currentTab = \'generate\'\" :class=\"{active: currentTab === \'generate\'}\">\n          <i class=\"fas fa-file-pdf\"></i> 生成PDF\n        </div>\n      </div>\n      \n      <div class=\"tab-content active\" v-if=\"currentTab === \'design\'\">\n        <div class=\"app-container\">\n          <div class=\"panel schema-editor\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-code\"></i>\n              <h2>Schema 编辑器</h2>\n            </div>\n            \n            <textarea v-model=\"schemaJson\" placeholder=\"在此输入您的JSON Schema...\"></textarea>\n            \n            <div class=\"actions\">\n              <button @click=\"generateForm\" class=\"primary\">\n                <i class=\"fas fa-play\"></i> 生成表单\n              </button>\n              <button @click=\"loadExample\" class=\"secondary\">\n                <i class=\"fas fa-lightbulb\"></i> 加载示例\n              </button>\n              <button @click=\"resetForm\" class=\"danger\">\n                <i class=\"fas fa-trash\"></i> 重置\n              </button>\n            </div>\n            \n            <div :class=\"[\'status\', status.type]\">\n              {{ status.message }}\n            </div>\n            \n            <div class=\"data-preview\">\n              <div class=\"data-preview-title\">\n                <i class=\"fas fa-database\"></i>\n                <h3>表单数据预览</h3>\n              </div>\n              <pre>{{ formData }}</pre>\n            </div>\n          </div>\n          \n          <div class=\"panel form-preview\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-window-maximize\"></i>\n              <h2>表单预览</h2>\n            </div>\n            \n            <div class=\"form-container\">\n              <template v-if=\"formFields.length\">\n                <div v-for=\"(field, index) in formFields\" :key=\"index\" class=\"form-group\">\n                  <h3>{{ field.title }}</h3>\n                  <p v-if=\"field.description\" class=\"description\">{{ field.description }}</p>\n                  \n                  <div v-if=\"field.type === \'string\' && !field.enum\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"text\" \n                      v-model=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.enum\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <select v-model=\"formData[field.key]\" :required=\"field.required\">\n                      <option v-for=\"(option, i) in field.enum\" :key=\"i\" :value=\"option\">\n                        {{ option }}\n                      </option>\n                    </select>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'number\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"number\" \n                      v-model.number=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :min=\"field.minimum\"\n                      :max=\"field.maximum\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                  \n                  <div v-if=\"field.type === \'boolean\'\">\n                    <div class=\"checkbox-group\">\n                      <input type=\"checkbox\" :id=\"field.key\" v-model=\"formData[field.key]\">\n                      <label :for=\"field.key\">{{ field.title }}</label>\n                    </div>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.radioOptions\">\n                    <label>{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <div class=\"radio-group\">\n                      <label v-for=\"(option, i) in field.radioOptions\" :key=\"i\">\n                        <input \n                          type=\"radio\" \n                          :name=\"field.key\" \n                          :value=\"option.value\"\n                          v-model=\"formData[field.key]\"\n                          :required=\"field.required\"\n                        >\n                        {{ option.label }}\n                      </label>\n                    </div>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.format === \'textarea\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <textarea \n                      v-model=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :rows=\"field.rows || 4\"\n                      :required=\"field.required\"\n                    ></textarea>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'date\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"date\" \n                      v-model=\"formData[field.key]\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                </div>\n                \n                <div class=\"form-group\">\n                  <button class=\"success\" @click=\"saveFormData\">\n                    <i class=\"fas fa-save\"></i> 保存表单数据\n                  </button>\n                </div>\n              </template>\n            </div>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"tab-content\" :class=\"{active: currentTab === \'template\'}\" v-if=\"currentTab === \'template\'\">\n        <div class=\"app-container\">\n          <div class=\"panel\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-edit\"></i>\n              <h2>合同模板编辑器</h2>\n            </div>\n            \n            <div class=\"template-container\">\n              <div class=\"editor-container\">\n                <div class=\"editor-header\">\n                  <div>合同模板编辑 (使用[[ ]]作为Golang模板语法)</div>\n                  <div>可用变量: {{ availableVariables.join(\', \') }}</div>\n                </div>\n                <textarea id=\"templateEditor\" v-model=\"contractTemplate\"></textarea>\n              </div>\n              \n              <div class=\"template-actions\">\n                <button class=\"primary\" @click=\"updateTemplatePreview\">\n                  <i class=\"fas fa-sync\"></i> 更新预览\n                </button>\n                <button class=\"success\" @click=\"saveTemplate\">\n                  <i class=\"fas fa-save\"></i> 保存模板\n                </button>\n                <button class=\"secondary\" @click=\"loadTemplateExample\">\n                  <i class=\"fas fa-file-alt\"></i> 加载示例\n                </button>\n              </div>\n              \n              <div class=\"template-syntax\">\n                <h4><i class=\"fas fa-info-circle\"></i> 模板语法说明</h4>\n                <p>使用双中括号 <code>[[ ]]</code> 作为模板变量标记：</p>\n                <div class=\"syntax-example\">\n                  &lt;p&gt;合同编号：[[.contractNo]]&lt;/p&gt;<br>\n                  &lt;p&gt;签署日期：[[.signDate]]&lt;/p&gt;\n                </div>\n                <p>在后端处理时，这些标记会被转换为Golang的标准模板语法 <code>{{ }}</code></p>\n              </div>\n              \n              <div class=\"contract-preview\" v-html=\"compiledTemplate\"></div>\n            </div>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"tab-content\" :class=\"{active: currentTab === \'generate\'}\" v-if=\"currentTab === \'generate\'\">\n        <div class=\"app-container\">\n          <div class=\"panel\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-file-pdf\"></i>\n              <h2>PDF生成</h2>\n            </div>\n            \n            <div class=\"status info\">\n              <i class=\"fas fa-info-circle\"></i> 表单数据已准备就绪，点击下方按钮生成PDF\n            </div>\n            \n            <div class=\"actions\">\n              <button class=\"success\" @click=\"generatePDF\">\n                <i class=\"fas fa-cogs\"></i> 生成PDF合同\n              </button>\n            </div>\n            \n            <div class=\"pdf-preview\" v-if=\"pdfGenerated\">\n              <img src=\"https://cdn-icons-png.flaticon.com/512/337/337946.png\" alt=\"PDF Preview\">\n              <p>合同PDF生成成功！</p>\n              <button class=\"primary download-btn\" @click=\"downloadPDF\">\n                <i class=\"fas fa-download\"></i> 下载PDF文件\n              </button>\n            </div>\n            <div class=\"pdf-preview\" v-else>\n              <i class=\"fas fa-file-pdf\" style=\"font-size: 120px; color: #e74c3c;\"></i>\n              <p>点击上方按钮生成PDF文件</p>\n            </div>\n            \n            <div class=\"data-preview\">\n              <div class=\"data-preview-title\">\n                <i class=\"fas fa-code\"></i>\n                <h3>最终合同HTML</h3>\n              </div>\n              <pre>{{ compiledTemplate }}</pre>\n            </div>\n          </div>\n        </div>\n      </div>\n    </div>\n  </div>\n\n  <script>\n    const { createApp, ref, reactive, computed, watch } = Vue;\n    \n    createApp({\n      setup() {\n        // 初始JSON Schema示例\n        const exampleSchema = {\n          \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n          \"title\": \"销售合同\",\n          \"description\": \"销售合同表单数据\",\n          \"type\": \"object\",\n          \"properties\": {\n            \"contractTitle\": {\n              \"type\": \"string\",\n              \"title\": \"合同标题\",\n              \"default\": \"产品销售合同\"\n            },\n            \"contractNo\": {\n              \"type\": \"string\",\n              \"title\": \"合同编号\",\n              \"pattern\": \"^CT-[0-9]{8}$\"\n            },\n            \"signDate\": {\n              \"type\": \"string\",\n              \"title\": \"签署日期\",\n              \"format\": \"date\"\n            },\n            \"sellerName\": {\n              \"type\": \"string\",\n              \"title\": \"卖方名称\"\n            },\n            \"buyerName\": {\n              \"type\": \"string\",\n              \"title\": \"买方名称\"\n            },\n            \"productName\": {\n              \"type\": \"string\",\n              \"title\": \"产品名称\"\n            },\n            \"quantity\": {\n              \"type\": \"number\",\n              \"title\": \"数量\",\n              \"minimum\": 1\n            },\n            \"unitPrice\": {\n              \"type\": \"number\",\n              \"title\": \"单价\",\n              \"minimum\": 0\n            },\n            \"totalAmount\": {\n              \"type\": \"number\",\n              \"title\": \"总金额\",\n              \"readOnly\": true\n            },\n            \"deliveryDate\": {\n              \"type\": \"string\",\n              \"title\": \"交货日期\",\n              \"format\": \"date\"\n            },\n            \"paymentTerms\": {\n              \"type\": \"string\",\n              \"title\": \"付款条款\",\n              \"enum\": [\"预付全款\", \"货到付款\", \"30天账期\", \"60天账期\"]\n            },\n            \"warrantyPeriod\": {\n              \"type\": \"string\",\n              \"title\": \"保修期\",\n              \"enum\": [\"6个月\", \"1年\", \"2年\", \"3年\", \"5年\"]\n            },\n            \"specialTerms\": {\n              \"type\": \"string\",\n              \"title\": \"特别条款\",\n              \"format\": \"textarea\",\n              \"rows\": 4\n            }\n          },\n          \"required\": [\"contractTitle\", \"contractNo\", \"sellerName\", \"buyerName\", \"productName\", \"quantity\", \"unitPrice\"]\n        };\n        \n        // 响应式数据\n        const currentTab = ref(\'design\');\n        const schemaJson = ref(JSON.stringify(exampleSchema, null, 2));\n        const formFields = ref([]);\n        const formData = ref({});\n        const status = reactive({\n          type: \'success\',\n          message: \'已加载示例Schema，点击\"生成表单\"按钮创建表单\'\n        });\n        const contractTemplate = ref(\'\');\n        const compiledTemplate = ref(\'\');\n        const availableVariables = ref([]);\n        const pdfGenerated = ref(false);\n        \n        // 合同模板示例（使用[[ ]]作为模板语法）\n        const templateExample = `\n<html>\n<head>\n  <title>[[.contractTitle]]</title>\n  <style>\n    .contract-container { \n      font-family: \'Times New Roman\', Times, serif; \n      line-height: 1.8;\n      color: #333;\n      max-width: 800px;\n      margin: 0 auto;\n      padding: 40px;\n      background: #fff;\n    }\n    .header {\n      text-align: center;\n      margin-bottom: 40px;\n    }\n    h1 {\n      font-size: 28px;\n      color: #1a3a6c;\n      margin-bottom: 10px;\n      border-bottom: 2px solid #1a3a6c;\n      padding-bottom: 15px;\n    }\n    .contract-info {\n      display: flex;\n      justify-content: space-between;\n      margin-bottom: 30px;\n      font-size: 16px;\n    }\n    .parties {\n      display: flex;\n      justify-content: space-between;\n      margin: 30px 0;\n    }\n    .party {\n      width: 45%;\n    }\n    .party h2 {\n      font-size: 20px;\n      border-bottom: 1px solid #ccc;\n      padding-bottom: 8px;\n      margin-bottom: 15px;\n    }\n    .terms-section {\n      margin: 30px 0;\n    }\n    .terms-section h2 {\n      font-size: 20px;\n      color: #1a3a6c;\n      border-left: 4px solid #1a3a6c;\n      padding-left: 10px;\n      margin: 25px 0 15px;\n    }\n    .signature-area {\n      display: flex;\n      justify-content: space-between;\n      margin-top: 80px;\n    }\n    .signature-box {\n      width: 45%;\n      text-align: center;\n    }\n    .signature-line {\n      height: 1px;\n      background: #333;\n      margin: 40px 0 10px;\n    }\n    .variable {\n      background: #fff8e1;\n      padding: 2px 5px;\n      border-radius: 4px;\n      border: 1px dashed #ffc107;\n    }\n  </style>\n</head>\n<body class=\"contract-container\">\n  <div class=\"header\">\n    <h1>[[.contractTitle]]</h1>\n    <p>合同编号：<span class=\"variable\">[[.contractNo]]</span></p>\n  </div>\n  \n  <div class=\"contract-info\">\n    <div>签订日期：<span class=\"variable\">[[.signDate]]</span></div>\n    <div>生效日期：<span class=\"variable\">[[.signDate]]</span></div>\n  </div>\n  \n  <div class=\"parties\">\n    <div class=\"party\">\n      <h2>甲方（卖方）</h2>\n      <p>名称：<span class=\"variable\">[[.sellerName]]</span></p>\n    </div>\n    \n    <div class=\"party\">\n      <h2>乙方（买方）</h2>\n      <p>名称：<span class=\"variable\">[[.buyerName]]</span></p>\n    </div>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第一条 产品信息</h2>\n    <p>1.1 产品名称：<span class=\"variable\">[[.productName]]</span></p>\n    <p>1.2 产品数量：<span class=\"variable\">[[.quantity]]</span></p>\n    <p>1.3 产品单价：人民币 <span class=\"variable\">[[.unitPrice]]</span> 元</p>\n    <p>1.4 总金额：人民币 <span class=\"variable\">[[.totalAmount]]</span> 元</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第二条 交货条款</h2>\n    <p>2.1 交货日期：<span class=\"variable\">[[.deliveryDate]]</span></p>\n    <p>2.2 交货地点：买方指定地点</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第三条 付款条款</h2>\n    <p>3.1 付款方式：<span class=\"variable\">[[.paymentTerms]]</span></p>\n    <p>3.2 付款期限：自合同签订之日起30日内</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第四条 保修条款</h2>\n    <p>4.1 保修期限：<span class=\"variable\">[[.warrantyPeriod]]</span></p>\n    <p>4.2 保修范围：产品制造缺陷</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第五条 特别条款</h2>\n    <p><span class=\"variable\">[[.specialTerms]]</span></p>\n  </div>\n  \n  <div class=\"signature-area\">\n    <div class=\"signature-box\">\n      <p>甲方（卖方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n    \n    <div class=\"signature-box\">\n      <p>乙方（买方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n  </div>\n</body>\n</html>\n`.trim();\n        \n        // 设置初始合同模板\n        contractTemplate.value = templateExample;\n        \n        // 生成表单方法\n        const generateForm = () => {\n          try {\n            const schema = JSON.parse(schemaJson.value);\n            \n            if (!schema || schema.type !== \'object\' || !schema.properties) {\n              throw new Error(\'无效的JSON Schema。根类型必须是object且包含properties属性。\');\n            }\n            \n            // 提取表单字段\n            formFields.value = Object.entries(schema.properties).map(([key, field]) => {\n              return {\n                key,\n                ...field,\n                required: (schema.required || []).includes(key)\n              };\n            });\n            \n            // 初始化表单数据\n            formData.value = {};\n            formFields.value.forEach(field => {\n              if (field.type === \'array\') {\n                formData.value[field.key] = [];\n              } else if (field.type === \'object\') {\n                formData.value[field.key] = {};\n                Object.keys(field.properties).forEach(prop => {\n                  formData.value[field.key][prop] = field.properties[prop].default || \'\';\n                });\n              } else {\n                formData.value[field.key] = field.default || \'\';\n                \n                // 设置默认日期\n                if (field.format === \'date\' && !field.default) {\n                  const today = new Date();\n                  formData.value[field.key] = today.toISOString().split(\'T\')[0];\n                }\n              }\n            });\n            \n            // 计算总金额\n            if (formData.value.quantity && formData.value.unitPrice) {\n              formData.value.totalAmount = formData.value.quantity * formData.value.unitPrice;\n            }\n            \n            status.type = \'success\';\n            status.message = `表单生成成功！共 ${formFields.value.length} 个字段`;\n            \n            // 更新可用变量\n            availableVariables.value = Object.keys(schema.properties);\n            \n            // 更新模板预览\n            updateTemplatePreview();\n          } catch (error) {\n            status.type = \'error-status\';\n            status.message = `错误: ${error.message}`;\n            formFields.value = [];\n          }\n        };\n        \n        // 加载示例Schema\n        const loadExample = () => {\n          schemaJson.value = JSON.stringify(exampleSchema, null, 2);\n          status.type = \'success\';\n          status.message = \'示例Schema已加载，点击\"生成表单\"按钮创建表单\';\n          generateForm();\n        };\n        \n        // 重置表单\n        const resetForm = () => {\n          schemaJson.value = \'\';\n          formFields.value = [];\n          formData.value = {};\n          status.type = \'success\';\n          status.message = \'表单已重置，请输入新的JSON Schema\';\n        };\n        \n        // 保存表单数据\n        const saveFormData = () => {\n          console.log(\'表单数据已保存:\', formData.value);\n          status.type = \'success\';\n          status.message = \'表单数据已保存！\';\n          \n          // 更新模板预览\n          updateTemplatePreview();\n        };\n        \n        // 更新模板预览\n        const updateTemplatePreview = () => {\n          try {\n            // 使用[[ ]]作为模板语法，避免与Vue冲突\n            let result = contractTemplate.value;\n            \n            // 替换所有变量\n            Object.keys(formData.value).forEach(key => {\n              const regex = new RegExp(`\\\\[\\\\[\\\\.${key}\\\\]\\\\]`, \'g\');\n              const value = formData.value[key] || \'\';\n              result = result.replace(regex, value);\n            });\n            \n            compiledTemplate.value = result;\n          } catch (error) {\n            compiledTemplate.value = `<div class=\"error\">模板渲染错误: ${error.message}</div>`;\n          }\n        };\n        \n        // 保存模板\n        const saveTemplate = () => {\n          console.log(\'合同模板已保存:\', contractTemplate.value);\n          status.type = \'success\';\n          status.message = \'合同模板已保存到数据库！\';\n          updateTemplatePreview();\n        };\n        \n        // 加载模板示例\n        const loadTemplateExample = () => {\n          contractTemplate.value = templateExample;\n          updateTemplatePreview();\n          status.type = \'success\';\n          status.message = \'已加载合同模板示例\';\n        };\n        \n        // 生成PDF\n        const generatePDF = () => {\n          // 模拟PDF生成过程\n          pdfGenerated.value = true;\n          \n          // 在实际应用中，这里会调用后端API：\n          // 1. 将contractTemplate和formData发送到后端\n          // 2. 后端将[[ ]]替换为{{ }}以符合Golang模板语法\n          // 3. 使用Golang模板引擎填充数据\n          // 4. 使用go-wkhtmltopdf生成PDF\n          // 5. 返回PDF文件URL\n          \n          console.log(\'生成PDF请求已发送\', {\n            template: contractTemplate.value,\n            data: formData.value\n          });\n          \n          status.type = \'success\';\n          status.message = \'PDF生成成功！在实际应用中，PDF将在后端生成并存储。\';\n        };\n        \n        // 下载PDF\n        const downloadPDF = () => {\n          alert(\'在实际应用中，这里会提供生成的PDF文件下载\');\n        };\n        \n        // 初始化\n        generateForm();\n        \n        // 监听表单数据变化\n        watch(formData, (newVal) => {\n          // 自动计算总金额\n          if (newVal.quantity && newVal.unitPrice) {\n            newVal.totalAmount = newVal.quantity * newVal.unitPrice;\n          }\n          updateTemplatePreview();\n        }, { deep: true });\n        \n        return {\n          currentTab,\n          schemaJson,\n          formFields,\n          formData,\n          status,\n          contractTemplate,\n          compiledTemplate,\n          availableVariables,\n          pdfGenerated,\n          generateForm,\n          loadExample,\n          resetForm,\n          saveFormData,\n          updateTemplatePreview,\n          saveTemplate,\n          loadTemplateExample,\n          generatePDF,\n          downloadPDF\n        };\n      }\n    }).mount(\'#app\');\n  </script>\n</body>\n</html>','2025-08-22 00:53:27','2025-08-22 00:53:27'),(2,'pdf_gen2','{\"type\": \"object\", \"title\": \"销售合同\", \"$schema\": \"http://json-schema.org/draft-07/schema#\", \"required\": [\"contractTitle\", \"contractNo\", \"sellerName\", \"buyerName\", \"productName\", \"quantity\", \"unitPrice\"], \"properties\": {\"quantity\": {\"type\": \"number\", \"title\": \"数量\", \"minimum\": 1}, \"signDate\": {\"type\": \"string\", \"title\": \"签署日期\", \"format\": \"date\"}, \"buyerName\": {\"type\": \"string\", \"title\": \"买方名称\"}, \"unitPrice\": {\"type\": \"number\", \"title\": \"单价\", \"minimum\": 0}, \"contractNo\": {\"type\": \"string\", \"title\": \"合同编号\", \"pattern\": \"^CT-[0-9]{8}$\"}, \"sellerName\": {\"type\": \"string\", \"title\": \"卖方名称\"}, \"productName\": {\"type\": \"string\", \"title\": \"产品名称\"}, \"totalAmount\": {\"type\": \"number\", \"title\": \"总金额\", \"readOnly\": true}, \"deliveryDate\": {\"type\": \"string\", \"title\": \"交货日期\", \"format\": \"date\"}, \"paymentTerms\": {\"enum\": [\"预付全款\", \"货到付款\", \"30天账期\", \"60天账期\"], \"type\": \"string\", \"title\": \"付款条款\"}, \"specialTerms\": {\"rows\": 4, \"type\": \"string\", \"title\": \"特别条款\", \"format\": \"textarea\"}, \"contractTitle\": {\"type\": \"string\", \"title\": \"合同标题\", \"default\": \"产品销售合同\"}, \"warrantyPeriod\": {\"enum\": [\"6个月\", \"1年\", \"2年\", \"3年\", \"5年\"], \"type\": \"string\", \"title\": \"保修期\"}}, \"description\": \"销售合同表单数据\"}','<!DOCTYPE html>\n<html lang=\"zh-CN\">\n<head>\n  <meta charset=\"UTF-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n  <title>合同PDF动态生成系统</title>\n  <script src=\"https://unpkg.com/vue@3/dist/vue.global.js\"></script>\n  <link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css\">\n  <style>\n    * {\n      margin: 0;\n      padding: 0;\n      box-sizing: border-box;\n      font-family: \'Segoe UI\', Tahoma, Geneva, Verdana, sans-serif;\n    }\n    \n    body {\n      background: linear-gradient(135deg, #1a2a6c, #b21f1f, #1a2a6c);\n      min-height: 100vh;\n      padding: 20px;\n      color: #333;\n    }\n    \n    .container {\n      max-width: 1800px;\n      margin: 0 auto;\n    }\n    \n    header {\n      text-align: center;\n      padding: 20px 0;\n      color: white;\n      margin-bottom: 30px;\n    }\n    \n    header h1 {\n      font-size: 2.8rem;\n      margin-bottom: 10px;\n      text-shadow: 0 2px 8px rgba(0,0,0,0.3);\n    }\n    \n    header p {\n      font-size: 1.3rem;\n      max-width: 900px;\n      margin: 0 auto;\n      opacity: 0.9;\n      line-height: 1.6;\n    }\n    \n    .tabs {\n      display: flex;\n      background: rgba(255, 255, 255, 0.1);\n      border-radius: 10px;\n      padding: 10px;\n      margin-bottom: 25px;\n      max-width: 1000px;\n      margin: 0 auto 25px;\n      flex-wrap: wrap;\n    }\n    \n    .tab {\n      flex: 1;\n      padding: 15px;\n      text-align: center;\n      color: white;\n      font-size: 1.1rem;\n      font-weight: 600;\n      cursor: pointer;\n      border-radius: 8px;\n      transition: all 0.3s ease;\n      min-width: 200px;\n    }\n    \n    .tab.active {\n      background: rgba(255, 255, 255, 0.2);\n      box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);\n    }\n    \n    .tab i {\n      margin-right: 10px;\n      font-size: 1.2rem;\n    }\n    \n    .app-container {\n      display: flex;\n      gap: 25px;\n      background: rgba(255, 255, 255, 0.95);\n      border-radius: 15px;\n      box-shadow: 0 15px 50px rgba(0, 0, 0, 0.3);\n      overflow: hidden;\n      min-height: 700px;\n    }\n    \n    .panel {\n      padding: 25px;\n      flex: 1;\n      display: flex;\n      flex-direction: column;\n      min-width: 0; /* 修复宽度问题 */\n    }\n    \n    .schema-editor {\n      background: #f8f9fa;\n      border-right: 1px solid #e0e0e0;\n    }\n    \n    .form-preview {\n      background: white;\n    }\n    \n    .panel-title {\n      display: flex;\n      align-items: center;\n      gap: 10px;\n      margin-bottom: 20px;\n      padding-bottom: 15px;\n      border-bottom: 2px solid #2c3e50;\n      color: #2c3e50;\n    }\n    \n    .panel-title i {\n      font-size: 1.6rem;\n    }\n    \n    .panel-title h2 {\n      font-size: 1.8rem;\n      font-weight: 700;\n    }\n    \n    textarea {\n      width: 100%;\n      height: 300px;\n      padding: 15px;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      font-family: monospace;\n      font-size: 15px;\n      resize: none;\n      transition: all 0.3s;\n    }\n    \n    textarea:focus {\n      outline: none;\n      border-color: #3498db;\n      box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);\n    }\n    \n    .actions {\n      display: flex;\n      gap: 15px;\n      margin-top: 20px;\n      flex-wrap: wrap;\n    }\n    \n    button {\n      padding: 14px 20px;\n      color: white;\n      border: none;\n      border-radius: 8px;\n      font-size: 1.1rem;\n      font-weight: 600;\n      cursor: pointer;\n      transition: all 0.3s;\n      display: flex;\n      align-items: center;\n      justify-content: center;\n      gap: 10px;\n      flex: 1;\n      min-width: 200px;\n    }\n    \n    button.primary {\n      background: linear-gradient(to right, #3498db, #2980b9);\n    }\n    \n    button.secondary {\n      background: linear-gradient(to right, #7f8c8d, #95a5a6);\n    }\n    \n    button.success {\n      background: linear-gradient(to right, #27ae60, #2ecc71);\n    }\n    \n    button.danger {\n      background: linear-gradient(to right, #e74c3c, #c0392b);\n    }\n    \n    button:hover {\n      transform: translateY(-3px);\n      box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);\n    }\n    \n    .form-container {\n      flex: 1;\n      overflow-y: auto;\n      padding: 10px;\n    }\n    \n    .form-group {\n      margin-bottom: 25px;\n      padding: 20px;\n      background: #f8f9fa;\n      border-radius: 10px;\n      box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);\n      position: relative;\n    }\n    \n    .form-group h3 {\n      margin-bottom: 15px;\n      color: #2c3e50;\n      font-size: 1.4rem;\n      padding-bottom: 8px;\n      border-bottom: 1px solid #eee;\n    }\n    \n    label {\n      display: block;\n      margin-bottom: 8px;\n      font-weight: 600;\n      color: #2c3e50;\n    }\n    \n    input, select, textarea {\n      width: 100%;\n      padding: 14px;\n      border: 1px solid #ced4da;\n      border-radius: 8px;\n      font-size: 1rem;\n      transition: all 0.3s;\n      background: white;\n    }\n    \n    input:focus, select:focus, textarea:focus {\n      outline: none;\n      border-color: #3498db;\n      box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.2);\n    }\n    \n    .error {\n      color: #e74c3c;\n      font-size: 0.9rem;\n      margin-top: 5px;\n    }\n    \n    .form-container:empty::before {\n      content: \"请提供有效的JSON Schema来生成表单\";\n      display: block;\n      text-align: center;\n      padding: 40px;\n      color: #7f8c8d;\n      font-size: 1.2rem;\n    }\n    \n    .data-preview {\n      background: #2c3e50;\n      color: white;\n      padding: 25px;\n      border-radius: 10px;\n      margin-top: 20px;\n      font-family: monospace;\n      white-space: pre-wrap;\n      max-height: 250px;\n      overflow-y: auto;\n      line-height: 1.5;\n    }\n    \n    .data-preview-title {\n      display: flex;\n      align-items: center;\n      gap: 10px;\n      margin-bottom: 15px;\n      color: #3498db;\n    }\n    \n    .status {\n      margin-top: 15px;\n      padding: 15px;\n      border-radius: 8px;\n      text-align: center;\n      font-weight: 500;\n      font-size: 1.1rem;\n    }\n    \n    .success {\n      background: #d4edda;\n      color: #155724;\n      border-left: 4px solid #28a745;\n    }\n    \n    .error-status {\n      background: #f8d7da;\n      color: #721c24;\n      border-left: 4px solid #dc3545;\n    }\n    \n    .info {\n      background: #cce5ff;\n      color: #004085;\n      border-left: 4px solid #007bff;\n    }\n    \n    .template-container {\n      display: flex;\n      flex-direction: column;\n      height: 100%;\n    }\n    \n    .editor-container {\n      flex: 1;\n      display: flex;\n      flex-direction: column;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      overflow: hidden;\n    }\n    \n    .editor-header {\n      background: #2c3e50;\n      color: white;\n      padding: 10px 15px;\n      font-weight: 600;\n      display: flex;\n      justify-content: space-between;\n      flex-wrap: wrap;\n    }\n    \n    #templateEditor {\n      height: 400px;\n      width: 100%;\n      border: none;\n      background: #1e1e1e;\n      color: #d4d4d4;\n      padding: 15px;\n      font-family: monospace;\n      font-size: 14px;\n      line-height: 1.5;\n      resize: vertical;\n    }\n    \n    .contract-preview {\n      flex: 1;\n      background: white;\n      border: 1px solid #ddd;\n      border-radius: 8px;\n      padding: 20px;\n      overflow-y: auto;\n      margin-top: 20px;\n      font-family: \'Times New Roman\', Times, serif;\n      line-height: 1.8;\n    }\n    \n    .contract-preview h1 {\n      text-align: center;\n      margin-bottom: 30px;\n      color: #2c3e50;\n      border-bottom: 2px solid #2c3e50;\n      padding-bottom: 15px;\n    }\n    \n    .contract-section {\n      margin-bottom: 30px;\n    }\n    \n    .contract-section h2 {\n      color: #2c3e50;\n      margin-bottom: 15px;\n      border-left: 4px solid #3498db;\n      padding-left: 10px;\n    }\n    \n    .contract-variable {\n      background: #fff8e1;\n      padding: 2px 5px;\n      border-radius: 4px;\n      border: 1px dashed #ffc107;\n    }\n    \n    .signature-area {\n      display: flex;\n      justify-content: space-between;\n      margin-top: 60px;\n      padding-top: 20px;\n      border-top: 1px solid #ccc;\n      flex-wrap: wrap;\n    }\n    \n    .signature-box {\n      width: 45%;\n      text-align: center;\n      min-width: 200px;\n      margin-bottom: 20px;\n    }\n    \n    .signature-line {\n      height: 1px;\n      background: #333;\n      margin: 40px 0 10px;\n    }\n    \n    .template-actions {\n      display: flex;\n      gap: 15px;\n      margin-top: 20px;\n      flex-wrap: wrap;\n    }\n    \n    .template-actions button {\n      flex: none;\n      width: auto;\n      padding: 12px 25px;\n    }\n    \n    .tab-content {\n      display: none;\n    }\n    \n    .tab-content.active {\n      display: block;\n    }\n    \n    .pdf-preview {\n      width: 100%;\n      height: 600px;\n      background: #f0f0f0;\n      border-radius: 8px;\n      display: flex;\n      align-items: center;\n      justify-content: center;\n      margin-top: 20px;\n      flex-direction: column;\n    }\n    \n    .pdf-preview img {\n      max-width: 100%;\n      max-height: 500px;\n      box-shadow: 0 5px 15px rgba(0,0,0,0.2);\n    }\n    \n    .pdf-preview p {\n      margin-top: 20px;\n      color: #7f8c8d;\n      font-size: 1.2rem;\n    }\n    \n    .download-btn {\n      margin-top: 25px;\n      padding: 15px 40px;\n      font-size: 1.2rem;\n    }\n    \n    .template-syntax {\n      margin-top: 10px;\n      padding: 15px;\n      background: #f8f9fa;\n      border-radius: 8px;\n      font-size: 0.9rem;\n    }\n    \n    .template-syntax h4 {\n      margin-bottom: 10px;\n      color: #2c3e50;\n    }\n    \n    .syntax-example {\n      background: #2c3e50;\n      color: white;\n      padding: 10px;\n      border-radius: 4px;\n      font-family: monospace;\n      margin-top: 5px;\n    }\n    \n    @media (max-width: 1200px) {\n      .app-container {\n        flex-direction: column;\n      }\n      \n      .schema-editor {\n        border-right: none;\n        border-bottom: 1px solid #eee;\n      }\n      \n      header h1 {\n        font-size: 2.2rem;\n      }\n      \n      .tab {\n        min-width: 150px;\n        padding: 10px;\n        font-size: 1rem;\n      }\n    }\n    \n    @media (max-width: 768px) {\n      .tabs {\n        flex-direction: column;\n      }\n      \n      .signature-box {\n        width: 100%;\n      }\n      \n      .panel-title h2 {\n        font-size: 1.5rem;\n      }\n      \n      button {\n        min-width: 100%;\n      }\n    }\n  </style>\n</head>\n<body>\n  <div id=\"app\">\n    <div class=\"container\">\n      <header>\n        <h1><i class=\"fas fa-file-contract\"></i> 合同PDF动态生成系统</h1>\n        <p>使用JSON Schema设计表单，编辑合同模板，动态生成PDF文件</p>\n      </header>\n      \n      <div class=\"tabs\">\n        <div class=\"tab\" @click=\"currentTab = \'design\'\" :class=\"{active: currentTab === \'design\'}\">\n          <i class=\"fas fa-drafting-compass\"></i> 表单设计\n        </div>\n        <div class=\"tab\" @click=\"currentTab = \'template\'\" :class=\"{active: currentTab === \'template\'}\">\n          <i class=\"fas fa-file-alt\"></i> 模板编辑\n        </div>\n        <div class=\"tab\" @click=\"currentTab = \'generate\'\" :class=\"{active: currentTab === \'generate\'}\">\n          <i class=\"fas fa-file-pdf\"></i> 生成PDF\n        </div>\n      </div>\n      \n      <div class=\"tab-content active\" v-if=\"currentTab === \'design\'\">\n        <div class=\"app-container\">\n          <div class=\"panel schema-editor\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-code\"></i>\n              <h2>Schema 编辑器</h2>\n            </div>\n            \n            <textarea v-model=\"schemaJson\" placeholder=\"在此输入您的JSON Schema...\"></textarea>\n            \n            <div class=\"actions\">\n              <button @click=\"generateForm\" class=\"primary\">\n                <i class=\"fas fa-play\"></i> 生成表单\n              </button>\n              <button @click=\"loadExample\" class=\"secondary\">\n                <i class=\"fas fa-lightbulb\"></i> 加载示例\n              </button>\n              <button @click=\"resetForm\" class=\"danger\">\n                <i class=\"fas fa-trash\"></i> 重置\n              </button>\n            </div>\n            \n            <div :class=\"[\'status\', status.type]\">\n              {{ status.message }}\n            </div>\n            \n            <div class=\"data-preview\">\n              <div class=\"data-preview-title\">\n                <i class=\"fas fa-database\"></i>\n                <h3>表单数据预览</h3>\n              </div>\n              <pre>{{ formData }}</pre>\n            </div>\n          </div>\n          \n          <div class=\"panel form-preview\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-window-maximize\"></i>\n              <h2>表单预览</h2>\n            </div>\n            \n            <div class=\"form-container\">\n              <template v-if=\"formFields.length\">\n                <div v-for=\"(field, index) in formFields\" :key=\"index\" class=\"form-group\">\n                  <h3>{{ field.title }}</h3>\n                  <p v-if=\"field.description\" class=\"description\">{{ field.description }}</p>\n                  \n                  <div v-if=\"field.type === \'string\' && !field.enum\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"text\" \n                      v-model=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.enum\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <select v-model=\"formData[field.key]\" :required=\"field.required\">\n                      <option v-for=\"(option, i) in field.enum\" :key=\"i\" :value=\"option\">\n                        {{ option }}\n                      </option>\n                    </select>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'number\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"number\" \n                      v-model.number=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :min=\"field.minimum\"\n                      :max=\"field.maximum\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                  \n                  <div v-if=\"field.type === \'boolean\'\">\n                    <div class=\"checkbox-group\">\n                      <input type=\"checkbox\" :id=\"field.key\" v-model=\"formData[field.key]\">\n                      <label :for=\"field.key\">{{ field.title }}</label>\n                    </div>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.radioOptions\">\n                    <label>{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <div class=\"radio-group\">\n                      <label v-for=\"(option, i) in field.radioOptions\" :key=\"i\">\n                        <input \n                          type=\"radio\" \n                          :name=\"field.key\" \n                          :value=\"option.value\"\n                          v-model=\"formData[field.key]\"\n                          :required=\"field.required\"\n                        >\n                        {{ option.label }}\n                      </label>\n                    </div>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'string\' && field.format === \'textarea\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <textarea \n                      v-model=\"formData[field.key]\"\n                      :placeholder=\"field.placeholder\"\n                      :rows=\"field.rows || 4\"\n                      :required=\"field.required\"\n                    ></textarea>\n                  </div>\n                  \n                  <div v-if=\"field.type === \'date\'\">\n                    <label :for=\"field.key\">{{ field.title }}<span v-if=\"field.required\"> *</span></label>\n                    <input \n                      type=\"date\" \n                      v-model=\"formData[field.key]\"\n                      :required=\"field.required\"\n                    >\n                  </div>\n                </div>\n                \n                <div class=\"form-group\">\n                  <button class=\"success\" @click=\"saveFormData\">\n                    <i class=\"fas fa-save\"></i> 保存表单数据\n                  </button>\n                </div>\n              </template>\n            </div>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"tab-content\" :class=\"{active: currentTab === \'template\'}\" v-if=\"currentTab === \'template\'\">\n        <div class=\"app-container\">\n          <div class=\"panel\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-edit\"></i>\n              <h2>合同模板编辑器</h2>\n            </div>\n            \n            <div class=\"template-container\">\n              <div class=\"editor-container\">\n                <div class=\"editor-header\">\n                  <div>合同模板编辑 (使用[[ ]]作为Golang模板语法)</div>\n                  <div>可用变量: {{ availableVariables.join(\', \') }}</div>\n                </div>\n                <textarea id=\"templateEditor\" v-model=\"contractTemplate\"></textarea>\n              </div>\n              \n              <div class=\"template-actions\">\n                <button class=\"primary\" @click=\"updateTemplatePreview\">\n                  <i class=\"fas fa-sync\"></i> 更新预览\n                </button>\n                <button class=\"success\" @click=\"saveTemplate\">\n                  <i class=\"fas fa-save\"></i> 保存模板\n                </button>\n                <button class=\"secondary\" @click=\"loadTemplateExample\">\n                  <i class=\"fas fa-file-alt\"></i> 加载示例\n                </button>\n              </div>\n              \n              <div class=\"template-syntax\">\n                <h4><i class=\"fas fa-info-circle\"></i> 模板语法说明</h4>\n                <p>使用双中括号 <code>[[ ]]</code> 作为模板变量标记：</p>\n                <div class=\"syntax-example\">\n                  &lt;p&gt;合同编号：[[.contractNo]]&lt;/p&gt;<br>\n                  &lt;p&gt;签署日期：[[.signDate]]&lt;/p&gt;\n                </div>\n                <p>在后端处理时，这些标记会被转换为Golang的标准模板语法 <code>{{ }}</code></p>\n              </div>\n              \n              <div class=\"contract-preview\" v-html=\"compiledTemplate\"></div>\n            </div>\n          </div>\n        </div>\n      </div>\n      \n      <div class=\"tab-content\" :class=\"{active: currentTab === \'generate\'}\" v-if=\"currentTab === \'generate\'\">\n        <div class=\"app-container\">\n          <div class=\"panel\">\n            <div class=\"panel-title\">\n              <i class=\"fas fa-file-pdf\"></i>\n              <h2>PDF生成</h2>\n            </div>\n            \n            <div class=\"status info\">\n              <i class=\"fas fa-info-circle\"></i> 表单数据已准备就绪，点击下方按钮生成PDF\n            </div>\n            \n            <div class=\"actions\">\n              <button class=\"success\" @click=\"generatePDF\">\n                <i class=\"fas fa-cogs\"></i> 生成PDF合同\n              </button>\n            </div>\n            \n            <div class=\"pdf-preview\" v-if=\"pdfGenerated\">\n              <img src=\"https://cdn-icons-png.flaticon.com/512/337/337946.png\" alt=\"PDF Preview\">\n              <p>合同PDF生成成功！</p>\n              <button class=\"primary download-btn\" @click=\"downloadPDF\">\n                <i class=\"fas fa-download\"></i> 下载PDF文件\n              </button>\n            </div>\n            <div class=\"pdf-preview\" v-else>\n              <i class=\"fas fa-file-pdf\" style=\"font-size: 120px; color: #e74c3c;\"></i>\n              <p>点击上方按钮生成PDF文件</p>\n            </div>\n            \n            <div class=\"data-preview\">\n              <div class=\"data-preview-title\">\n                <i class=\"fas fa-code\"></i>\n                <h3>最终合同HTML</h3>\n              </div>\n              <pre>{{ compiledTemplate }}</pre>\n            </div>\n          </div>\n        </div>\n      </div>\n    </div>\n  </div>\n\n  <script>\n    const { createApp, ref, reactive, computed, watch } = Vue;\n    \n    createApp({\n      setup() {\n        // 初始JSON Schema示例\n        const exampleSchema = {\n          \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n          \"title\": \"销售合同\",\n          \"description\": \"销售合同表单数据\",\n          \"type\": \"object\",\n          \"properties\": {\n            \"contractTitle\": {\n              \"type\": \"string\",\n              \"title\": \"合同标题\",\n              \"default\": \"产品销售合同\"\n            },\n            \"contractNo\": {\n              \"type\": \"string\",\n              \"title\": \"合同编号\",\n              \"pattern\": \"^CT-[0-9]{8}$\"\n            },\n            \"signDate\": {\n              \"type\": \"string\",\n              \"title\": \"签署日期\",\n              \"format\": \"date\"\n            },\n            \"sellerName\": {\n              \"type\": \"string\",\n              \"title\": \"卖方名称\"\n            },\n            \"buyerName\": {\n              \"type\": \"string\",\n              \"title\": \"买方名称\"\n            },\n            \"productName\": {\n              \"type\": \"string\",\n              \"title\": \"产品名称\"\n            },\n            \"quantity\": {\n              \"type\": \"number\",\n              \"title\": \"数量\",\n              \"minimum\": 1\n            },\n            \"unitPrice\": {\n              \"type\": \"number\",\n              \"title\": \"单价\",\n              \"minimum\": 0\n            },\n            \"totalAmount\": {\n              \"type\": \"number\",\n              \"title\": \"总金额\",\n              \"readOnly\": true\n            },\n            \"deliveryDate\": {\n              \"type\": \"string\",\n              \"title\": \"交货日期\",\n              \"format\": \"date\"\n            },\n            \"paymentTerms\": {\n              \"type\": \"string\",\n              \"title\": \"付款条款\",\n              \"enum\": [\"预付全款\", \"货到付款\", \"30天账期\", \"60天账期\"]\n            },\n            \"warrantyPeriod\": {\n              \"type\": \"string\",\n              \"title\": \"保修期\",\n              \"enum\": [\"6个月\", \"1年\", \"2年\", \"3年\", \"5年\"]\n            },\n            \"specialTerms\": {\n              \"type\": \"string\",\n              \"title\": \"特别条款\",\n              \"format\": \"textarea\",\n              \"rows\": 4\n            }\n          },\n          \"required\": [\"contractTitle\", \"contractNo\", \"sellerName\", \"buyerName\", \"productName\", \"quantity\", \"unitPrice\"]\n        };\n        \n        // 响应式数据\n        const currentTab = ref(\'design\');\n        const schemaJson = ref(JSON.stringify(exampleSchema, null, 2));\n        const formFields = ref([]);\n        const formData = ref({});\n        const status = reactive({\n          type: \'success\',\n          message: \'已加载示例Schema，点击\"生成表单\"按钮创建表单\'\n        });\n        const contractTemplate = ref(\'\');\n        const compiledTemplate = ref(\'\');\n        const availableVariables = ref([]);\n        const pdfGenerated = ref(false);\n        \n        // 合同模板示例（使用[[ ]]作为模板语法）\n        const templateExample = `\n<html>\n<head>\n  <title>[[.contractTitle]]</title>\n  <style>\n    .contract-container { \n      font-family: \'Times New Roman\', Times, serif; \n      line-height: 1.8;\n      color: #333;\n      max-width: 800px;\n      margin: 0 auto;\n      padding: 40px;\n      background: #fff;\n    }\n    .header {\n      text-align: center;\n      margin-bottom: 40px;\n    }\n    h1 {\n      font-size: 28px;\n      color: #1a3a6c;\n      margin-bottom: 10px;\n      border-bottom: 2px solid #1a3a6c;\n      padding-bottom: 15px;\n    }\n    .contract-info {\n      display: flex;\n      justify-content: space-between;\n      margin-bottom: 30px;\n      font-size: 16px;\n    }\n    .parties {\n      display: flex;\n      justify-content: space-between;\n      margin: 30px 0;\n    }\n    .party {\n      width: 45%;\n    }\n    .party h2 {\n      font-size: 20px;\n      border-bottom: 1px solid #ccc;\n      padding-bottom: 8px;\n      margin-bottom: 15px;\n    }\n    .terms-section {\n      margin: 30px 0;\n    }\n    .terms-section h2 {\n      font-size: 20px;\n      color: #1a3a6c;\n      border-left: 4px solid #1a3a6c;\n      padding-left: 10px;\n      margin: 25px 0 15px;\n    }\n    .signature-area {\n      display: flex;\n      justify-content: space-between;\n      margin-top: 80px;\n    }\n    .signature-box {\n      width: 45%;\n      text-align: center;\n    }\n    .signature-line {\n      height: 1px;\n      background: #333;\n      margin: 40px 0 10px;\n    }\n    .variable {\n      background: #fff8e1;\n      padding: 2px 5px;\n      border-radius: 4px;\n      border: 1px dashed #ffc107;\n    }\n  </style>\n</head>\n<body class=\"contract-container\">\n  <div class=\"header\">\n    <h1>[[.contractTitle]]</h1>\n    <p>合同编号：<span class=\"variable\">[[.contractNo]]</span></p>\n  </div>\n  \n  <div class=\"contract-info\">\n    <div>签订日期：<span class=\"variable\">[[.signDate]]</span></div>\n    <div>生效日期：<span class=\"variable\">[[.signDate]]</span></div>\n  </div>\n  \n  <div class=\"parties\">\n    <div class=\"party\">\n      <h2>甲方（卖方）</h2>\n      <p>名称：<span class=\"variable\">[[.sellerName]]</span></p>\n    </div>\n    \n    <div class=\"party\">\n      <h2>乙方（买方）</h2>\n      <p>名称：<span class=\"variable\">[[.buyerName]]</span></p>\n    </div>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第一条 产品信息</h2>\n    <p>1.1 产品名称：<span class=\"variable\">[[.productName]]</span></p>\n    <p>1.2 产品数量：<span class=\"variable\">[[.quantity]]</span></p>\n    <p>1.3 产品单价：人民币 <span class=\"variable\">[[.unitPrice]]</span> 元</p>\n    <p>1.4 总金额：人民币 <span class=\"variable\">[[.totalAmount]]</span> 元</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第二条 交货条款</h2>\n    <p>2.1 交货日期：<span class=\"variable\">[[.deliveryDate]]</span></p>\n    <p>2.2 交货地点：买方指定地点</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第三条 付款条款</h2>\n    <p>3.1 付款方式：<span class=\"variable\">[[.paymentTerms]]</span></p>\n    <p>3.2 付款期限：自合同签订之日起30日内</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第四条 保修条款</h2>\n    <p>4.1 保修期限：<span class=\"variable\">[[.warrantyPeriod]]</span></p>\n    <p>4.2 保修范围：产品制造缺陷</p>\n  </div>\n  \n  <div class=\"terms-section\">\n    <h2>第五条 特别条款</h2>\n    <p><span class=\"variable\">[[.specialTerms]]</span></p>\n  </div>\n  \n  <div class=\"signature-area\">\n    <div class=\"signature-box\">\n      <p>甲方（卖方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n    \n    <div class=\"signature-box\">\n      <p>乙方（买方）签字：</p>\n      <div class=\"signature-line\"></div>\n      <p>日期：<span class=\"variable\">[[.signDate]]</span></p>\n    </div>\n  </div>\n</body>\n</html>\n`.trim();\n        \n        // 设置初始合同模板\n        contractTemplate.value = templateExample;\n        \n        // 生成表单方法\n        const generateForm = () => {\n          try {\n            const schema = JSON.parse(schemaJson.value);\n            \n            if (!schema || schema.type !== \'object\' || !schema.properties) {\n              throw new Error(\'无效的JSON Schema。根类型必须是object且包含properties属性。\');\n            }\n            \n            // 提取表单字段\n            formFields.value = Object.entries(schema.properties).map(([key, field]) => {\n              return {\n                key,\n                ...field,\n                required: (schema.required || []).includes(key)\n              };\n            });\n            \n            // 初始化表单数据\n            formData.value = {};\n            formFields.value.forEach(field => {\n              if (field.type === \'array\') {\n                formData.value[field.key] = [];\n              } else if (field.type === \'object\') {\n                formData.value[field.key] = {};\n                Object.keys(field.properties).forEach(prop => {\n                  formData.value[field.key][prop] = field.properties[prop].default || \'\';\n                });\n              } else {\n                formData.value[field.key] = field.default || \'\';\n                \n                // 设置默认日期\n                if (field.format === \'date\' && !field.default) {\n                  const today = new Date();\n                  formData.value[field.key] = today.toISOString().split(\'T\')[0];\n                }\n              }\n            });\n            \n            // 计算总金额\n            if (formData.value.quantity && formData.value.unitPrice) {\n              formData.value.totalAmount = formData.value.quantity * formData.value.unitPrice;\n            }\n            \n            status.type = \'success\';\n            status.message = `表单生成成功！共 ${formFields.value.length} 个字段`;\n            \n            // 更新可用变量\n            availableVariables.value = Object.keys(schema.properties);\n            \n            // 更新模板预览\n            updateTemplatePreview();\n          } catch (error) {\n            status.type = \'error-status\';\n            status.message = `错误: ${error.message}`;\n            formFields.value = [];\n          }\n        };\n        \n        // 加载示例Schema\n        const loadExample = () => {\n          schemaJson.value = JSON.stringify(exampleSchema, null, 2);\n          status.type = \'success\';\n          status.message = \'示例Schema已加载，点击\"生成表单\"按钮创建表单\';\n          generateForm();\n        };\n        \n        // 重置表单\n        const resetForm = () => {\n          schemaJson.value = \'\';\n          formFields.value = [];\n          formData.value = {};\n          status.type = \'success\';\n          status.message = \'表单已重置，请输入新的JSON Schema\';\n        };\n        \n        // 保存表单数据\n        const saveFormData = () => {\n          console.log(\'表单数据已保存:\', formData.value);\n          status.type = \'success\';\n          status.message = \'表单数据已保存！\';\n          \n          // 更新模板预览\n          updateTemplatePreview();\n        };\n        \n        // 更新模板预览\n        const updateTemplatePreview = () => {\n          try {\n            // 使用[[ ]]作为模板语法，避免与Vue冲突\n            let result = contractTemplate.value;\n            \n            // 替换所有变量\n            Object.keys(formData.value).forEach(key => {\n              const regex = new RegExp(`\\\\[\\\\[\\\\.${key}\\\\]\\\\]`, \'g\');\n              const value = formData.value[key] || \'\';\n              result = result.replace(regex, value);\n            });\n            \n            compiledTemplate.value = result;\n          } catch (error) {\n            compiledTemplate.value = `<div class=\"error\">模板渲染错误: ${error.message}</div>`;\n          }\n        };\n        \n        // 保存模板\n        const saveTemplate = () => {\n          console.log(\'合同模板已保存:\', contractTemplate.value);\n          status.type = \'success\';\n          status.message = \'合同模板已保存到数据库！\';\n          updateTemplatePreview();\n        };\n        \n        // 加载模板示例\n        const loadTemplateExample = () => {\n          contractTemplate.value = templateExample;\n          updateTemplatePreview();\n          status.type = \'success\';\n          status.message = \'已加载合同模板示例\';\n        };\n        \n        // 生成PDF\n        const generatePDF = () => {\n          // 模拟PDF生成过程\n          pdfGenerated.value = true;\n          \n          // 在实际应用中，这里会调用后端API：\n          // 1. 将contractTemplate和formData发送到后端\n          // 2. 后端将[[ ]]替换为{{ }}以符合Golang模板语法\n          // 3. 使用Golang模板引擎填充数据\n          // 4. 使用go-wkhtmltopdf生成PDF\n          // 5. 返回PDF文件URL\n          \n          console.log(\'生成PDF请求已发送\', {\n            template: contractTemplate.value,\n            data: formData.value\n          });\n          \n          status.type = \'success\';\n          status.message = \'PDF生成成功！在实际应用中，PDF将在后端生成并存储。\';\n        };\n        \n        // 下载PDF\n        const downloadPDF = () => {\n          alert(\'在实际应用中，这里会提供生成的PDF文件下载\');\n        };\n        \n        // 初始化\n        generateForm();\n        \n        // 监听表单数据变化\n        watch(formData, (newVal) => {\n          // 自动计算总金额\n          if (newVal.quantity && newVal.unitPrice) {\n            newVal.totalAmount = newVal.quantity * newVal.unitPrice;\n          }\n          updateTemplatePreview();\n        }, { deep: true });\n        \n        return {\n          currentTab,\n          schemaJson,\n          formFields,\n          formData,\n          status,\n          contractTemplate,\n          compiledTemplate,\n          availableVariables,\n          pdfGenerated,\n          generateForm,\n          loadExample,\n          resetForm,\n          saveFormData,\n          updateTemplatePreview,\n          saveTemplate,\n          loadTemplateExample,\n          generatePDF,\n          downloadPDF\n        };\n      }\n    }).mount(\'#app\');\n  </script>\n</body>\n</html>','2025-08-22 00:53:31','2025-08-22 00:53:31');
/*!40000 ALTER TABLE `pdf_gens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permissions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '权限标识',
  `code` varchar(255) NOT NULL COMMENT '权限标识',
  `type` int(11) NOT NULL COMMENT '权限类型: 1-菜单，2-按钮，3-API',
  `description` varchar(255) NOT NULL COMMENT '描述信息',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,'首页','workplace',1,'',1,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(2,'系统','system',1,'',2,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(3,'菜单管理','system:auth:menu',1,'',3,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(4,'角色管理','system:auth:role',1,'',4,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(5,'权限管理','system:auth:permission',1,'',5,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(6,'用户管理','system:user',1,'',6,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(7,'附件中心','system:netdisk',1,'',7,'2025-08-22 00:53:31','2025-08-22 00:53:31'),(8,'代码生成','system:crud:index',1,'',8,'2025-08-22 00:57:59','2025-08-22 00:57:59'),(9,'数据库表设计','system:crud:column',1,'',9,'2025-08-22 01:32:00','2025-08-22 01:32:00');
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_permissions`
--

DROP TABLE IF EXISTS `role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_permissions` (
  `role_id` bigint(20) unsigned NOT NULL,
  `permission_id` bigint(20) unsigned NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_permissions`
--

LOCK TABLES `role_permissions` WRITE;
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;
INSERT INTO `role_permissions` VALUES (1,9),(1,8),(1,7),(1,6),(1,5),(1,4),(1,3),(1,2),(1,1);
/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `label` varchar(255) NOT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `is_disable` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '0',
  `tenant_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `is_admin` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin','超级管理员','超级管理员',0,0,0,0,'2025-08-22 00:53:31','2025-08-22 01:50:22'),(2,'instructor','教练','教练',0,0,0,0,'2025-08-22 00:53:31','2025-08-22 00:53:31');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rooms`
--

DROP TABLE IF EXISTS `rooms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rooms` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rooms`
--

LOCK TABLES `rooms` WRITE;
/*!40000 ALTER TABLE `rooms` DISABLE KEYS */;
/*!40000 ALTER TABLE `rooms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
  `user_id` bigint(20) unsigned NOT NULL,
  `role_id` bigint(20) unsigned NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
INSERT INTO `user_roles` VALUES (1,1),(2,2);
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `sex` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `openid` varchar(255) NOT NULL,
  `unionid` varchar(255) NOT NULL,
  `realname` varchar(255) NOT NULL,
  `id_card_number` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL,
  `remark` text NOT NULL,
  `last_login` datetime DEFAULT CURRENT_TIMESTAMP,
  `status` enum('正常','暂停','关闭') NOT NULL DEFAULT '正常',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','','$2a$12$lnPfJT1Nku4jsTe.ImmLOeDmSRES6bt5o/zL8zfjC0nzW/yvOpELq','','','','Nolan5802','','','','2025-08-22 12:16:45','正常','2025-08-22 00:53:30','2025-08-22 04:16:45'),(2,'test','','$2a$12$0/D3KIZifnH8QKTu7SJFOeul8x7Ed4g5uAHWejJwISG/l1ngxHoE.','','','','Gerlach1557','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:31'),(3,'Holden Cummings','','$2a$12$8C0chA0bO3rAKyO.jl/xLedlg7j/hPuafKTmjTu/l4jvoxYjWokxW','','','','Leuschke7074','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(4,'Emma Swaniawski','','$2a$12$Z4KyRbm5hCeoax3jRGU66uTA2ME1yPGX.e2Wo9h4/Kzl0eJwtCb4O','','','','Leffler5651','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(5,'Octavia Schinner','','$2a$12$UQD7Eayn4/JQd7LwMG4JR.P6qeKj5NKSB5yqFPRBqW9/GWyTJ.Emu','','','','Wuckert9957','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(6,'Ida Hamill','','$2a$12$s5suKBuWY.G85fp5kyGml.bQnhxCUysiT/ENTDvYv0Ar5Vtn.sdUy','','','','Donnelly4510','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(7,'Astrid Kshlerin','','$2a$12$htl0QNTqdief/iz0Q2Li4.hvvqOD6voMralx80wAX2dRmKNpesy9i','','','','Rogahn4255','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(8,'Anna Hansen','','$2a$12$lBO8FaiFW.pb6q5mra2Ch.31QiOHhr.Hzmty2flZpbeiWZbQzUh8W','','','','Ondricka8870','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(9,'Cristobal Bauch','','$2a$12$HcUoNmRZ0vM2F3tpKqV81eRAYl3cxbX3IphBOWDyLPAFQ3ZNeOyee','','','','Gulgowski7939','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30'),(10,'Wilford Oberbrunner','','$2a$12$uEDQnAOAcVIVzBhxnWb6uu4kzS.SjUFstwKBUVGTm5YdwRrbiQZr.','','','','Renner7081','','','',NULL,'正常','2025-08-22 00:53:30','2025-08-22 00:53:30');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'goravel'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-22 12:29:24
