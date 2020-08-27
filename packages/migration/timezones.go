/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package migration

var sqlTimeZonesSQL = `
	{{head "1_time_zones"}}
		t.Column("id", "bigint", {"default": "0"})
		t.Column("name", "string", {"default": "", "size":255})
		t.Column("offset", "string", {"default": "", "size":6})
	{{footer "primary" "unique(name)"}}
`

var timeZonesSQL = `
INSERT INTO "1_time_zones" VALUES 
(next_id('1_time_zones'), 'Africa/Abidjan', 'UTC'),
(next_id('1_time_zones'), 'Africa/Accra', 'UTC'),
(next_id('1_time_zones'), 'Africa/Addis_Ababa', '+03:00'),
(next_id('1_time_zones'), 'Africa/Algiers', '+01:00'),
(next_id('1_time_zones'), 'Africa/Asmara', '+03:00'),
(next_id('1_time_zones'), 'Africa/Bamako', 'UTC'),
(next_id('1_time_zones'), 'Africa/Bangui', '+01:00'),
(next_id('1_time_zones'), 'Africa/Banjul', 'UTC'),
(next_id('1_time_zones'), 'Africa/Bissau', 'UTC'),
(next_id('1_time_zones'), 'Africa/Blantyre', '+02:00'),
(next_id('1_time_zones'), 'Africa/Brazzaville', '+01:00'),
(next_id('1_time_zones'), 'Africa/Bujumbura', '+02:00'),
(next_id('1_time_zones'), 'Africa/Cairo', '+02:00'),
(next_id('1_time_zones'), 'Africa/Casablanca', '+01:00'),
(next_id('1_time_zones'), 'Africa/Ceuta', '+02:00'),
(next_id('1_time_zones'), 'Africa/Conakry', 'UTC'),
(next_id('1_time_zones'), 'Africa/Dakar', 'UTC'),
(next_id('1_time_zones'), 'Africa/Dar_es_Salaam', '+03:00'),
(next_id('1_time_zones'), 'Africa/Djibouti', '+03:00'),
(next_id('1_time_zones'), 'Africa/Douala', '+01:00'),
(next_id('1_time_zones'), 'Africa/El_Aaiun', '+01:00'),
(next_id('1_time_zones'), 'Africa/Freetown', 'UTC'),
(next_id('1_time_zones'), 'Africa/Gaborone', '+02:00'),
(next_id('1_time_zones'), 'Africa/Harare', '+02:00'),
(next_id('1_time_zones'), 'Africa/Johannesburg', '+02:00'),
(next_id('1_time_zones'), 'Africa/Juba', '+03:00'),
(next_id('1_time_zones'), 'Africa/Kampala', '+03:00'),
(next_id('1_time_zones'), 'Africa/Khartoum', '+02:00'),
(next_id('1_time_zones'), 'Africa/Kigali', '+02:00'),
(next_id('1_time_zones'), 'Africa/Kinshasa', '+01:00'),
(next_id('1_time_zones'), 'Africa/Lagos', '+01:00'),
(next_id('1_time_zones'), 'Africa/Libreville', '+01:00'),
(next_id('1_time_zones'), 'Africa/Lome', 'UTC'),
(next_id('1_time_zones'), 'Africa/Luanda', '+01:00'),
(next_id('1_time_zones'), 'Africa/Lubumbashi', '+02:00'),
(next_id('1_time_zones'), 'Africa/Lusaka', '+02:00'),
(next_id('1_time_zones'), 'Africa/Malabo', '+01:00'),
(next_id('1_time_zones'), 'Africa/Maputo', '+02:00'),
(next_id('1_time_zones'), 'Africa/Maseru', '+02:00'),
(next_id('1_time_zones'), 'Africa/Mbabane', '+02:00'),
(next_id('1_time_zones'), 'Africa/Mogadishu', '+03:00'),
(next_id('1_time_zones'), 'Africa/Monrovia', 'UTC'),
(next_id('1_time_zones'), 'Africa/Nairobi', '+03:00'),
(next_id('1_time_zones'), 'Africa/Ndjamena', '+01:00'),
(next_id('1_time_zones'), 'Africa/Niamey', '+01:00'),
(next_id('1_time_zones'), 'Africa/Nouakchott', 'UTC'),
(next_id('1_time_zones'), 'Africa/Ouagadougou', 'UTC'),
(next_id('1_time_zones'), 'Africa/Porto-Novo', '+01:00'),
(next_id('1_time_zones'), 'Africa/Sao_Tome', 'UTC'),
(next_id('1_time_zones'), 'Africa/Tripoli', '+02:00'),
(next_id('1_time_zones'), 'Africa/Tunis', '+01:00'),
(next_id('1_time_zones'), 'Africa/Windhoek', '+02:00'),
(next_id('1_time_zones'), 'America/Adak', '-09:00'),
(next_id('1_time_zones'), 'America/Anchorage', '-08:00'),
(next_id('1_time_zones'), 'America/Anguilla', '-04:00'),
(next_id('1_time_zones'), 'America/Antigua', '-04:00'),
(next_id('1_time_zones'), 'America/Araguaina', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Buenos_Aires', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Catamarca', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Cordoba', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Jujuy', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/La_Rioja', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Mendoza', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Rio_Gallegos', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Salta', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/San_Juan', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/San_Luis', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Tucuman', '-03:00'),
(next_id('1_time_zones'), 'America/Argentina/Ushuaia', '-03:00'),
(next_id('1_time_zones'), 'America/Aruba', '-04:00'),
(next_id('1_time_zones'), 'America/Asuncion', '-04:00'),
(next_id('1_time_zones'), 'America/Atikokan', '-05:00'),
(next_id('1_time_zones'), 'America/Bahia', '-03:00'),
(next_id('1_time_zones'), 'America/Bahia_Banderas', '-05:00'),
(next_id('1_time_zones'), 'America/Barbados', '-04:00'),
(next_id('1_time_zones'), 'America/Belem', '-03:00'),
(next_id('1_time_zones'), 'America/Belize', '-06:00'),
(next_id('1_time_zones'), 'America/Blanc-Sablon', '-04:00'),
(next_id('1_time_zones'), 'America/Boa_Vista', '-04:00'),
(next_id('1_time_zones'), 'America/Bogota', '-05:00'),
(next_id('1_time_zones'), 'America/Boise', '-06:00'),
(next_id('1_time_zones'), 'America/Cambridge_Bay', '-06:00'),
(next_id('1_time_zones'), 'America/Campo_Grande', '-04:00'),
(next_id('1_time_zones'), 'America/Cancun', '-05:00'),
(next_id('1_time_zones'), 'America/Caracas', '-04:00'),
(next_id('1_time_zones'), 'America/Cayenne', '-03:00'),
(next_id('1_time_zones'), 'America/Cayman', '-05:00'),
(next_id('1_time_zones'), 'America/Chicago', '-05:00'),
(next_id('1_time_zones'), 'America/Chihuahua', '-06:00'),
(next_id('1_time_zones'), 'America/Costa_Rica', '-06:00'),
(next_id('1_time_zones'), 'America/Creston', '-07:00'),
(next_id('1_time_zones'), 'America/Cuiaba', '-04:00'),
(next_id('1_time_zones'), 'America/Curacao', '-04:00'),
(next_id('1_time_zones'), 'America/Danmarkshavn', 'UTC'),
(next_id('1_time_zones'), 'America/Dawson', '-07:00'),
(next_id('1_time_zones'), 'America/Dawson_Creek', '-07:00'),
(next_id('1_time_zones'), 'America/Denver', '-06:00'),
(next_id('1_time_zones'), 'America/Detroit', '-04:00'),
(next_id('1_time_zones'), 'America/Dominica', '-04:00'),
(next_id('1_time_zones'), 'America/Edmonton', '-06:00'),
(next_id('1_time_zones'), 'America/Eirunepe', '-05:00'),
(next_id('1_time_zones'), 'America/El_Salvador', '-06:00'),
(next_id('1_time_zones'), 'America/Fort_Nelson', '-07:00'),
(next_id('1_time_zones'), 'America/Fortaleza', '-03:00'),
(next_id('1_time_zones'), 'America/Glace_Bay', '-03:00'),
(next_id('1_time_zones'), 'America/Godthab', '-02:00'),
(next_id('1_time_zones'), 'America/Goose_Bay', '-03:00'),
(next_id('1_time_zones'), 'America/Grand_Turk', '-04:00'),
(next_id('1_time_zones'), 'America/Grenada', '-04:00'),
(next_id('1_time_zones'), 'America/Guadeloupe', '-04:00'),
(next_id('1_time_zones'), 'America/Guatemala', '-06:00'),
(next_id('1_time_zones'), 'America/Guayaquil', '-05:00'),
(next_id('1_time_zones'), 'America/Guyana', '-04:00'),
(next_id('1_time_zones'), 'America/Halifax', '-03:00'),
(next_id('1_time_zones'), 'America/Havana', '-04:00'),
(next_id('1_time_zones'), 'America/Hermosillo', '-07:00'),
(next_id('1_time_zones'), 'America/Indiana/Indianapolis', '-04:00'),
(next_id('1_time_zones'), 'America/Indiana/Knox', '-05:00'),
(next_id('1_time_zones'), 'America/Indiana/Marengo', '-04:00'),
(next_id('1_time_zones'), 'America/Indiana/Petersburg', '-04:00'),
(next_id('1_time_zones'), 'America/Indiana/Tell_City', '-05:00'),
(next_id('1_time_zones'), 'America/Indiana/Vevay', '-04:00'),
(next_id('1_time_zones'), 'America/Indiana/Vincennes', '-04:00'),
(next_id('1_time_zones'), 'America/Indiana/Winamac', '-04:00'),
(next_id('1_time_zones'), 'America/Inuvik', '-06:00'),
(next_id('1_time_zones'), 'America/Iqaluit', '-04:00'),
(next_id('1_time_zones'), 'America/Jamaica', '-05:00'),
(next_id('1_time_zones'), 'America/Juneau', '-08:00'),
(next_id('1_time_zones'), 'America/Kentucky/Louisville', '-04:00'),
(next_id('1_time_zones'), 'America/Kentucky/Monticello', '-04:00'),
(next_id('1_time_zones'), 'America/Kralendijk', '-04:00'),
(next_id('1_time_zones'), 'America/La_Paz', '-04:00'),
(next_id('1_time_zones'), 'America/Lima', '-05:00'),
(next_id('1_time_zones'), 'America/Los_Angeles', '-07:00'),
(next_id('1_time_zones'), 'America/Lower_Princes', '-04:00'),
(next_id('1_time_zones'), 'America/Maceio', '-03:00'),
(next_id('1_time_zones'), 'America/Managua', '-06:00'),
(next_id('1_time_zones'), 'America/Manaus', '-04:00'),
(next_id('1_time_zones'), 'America/Marigot', '-04:00'),
(next_id('1_time_zones'), 'America/Martinique', '-04:00'),
(next_id('1_time_zones'), 'America/Matamoros', '-05:00'),
(next_id('1_time_zones'), 'America/Mazatlan', '-06:00'),
(next_id('1_time_zones'), 'America/Menominee', '-05:00'),
(next_id('1_time_zones'), 'America/Merida', '-05:00'),
(next_id('1_time_zones'), 'America/Metlakatla', '-08:00'),
(next_id('1_time_zones'), 'America/Mexico_City', '-05:00'),
(next_id('1_time_zones'), 'America/Miquelon', '-02:00'),
(next_id('1_time_zones'), 'America/Moncton', '-03:00'),
(next_id('1_time_zones'), 'America/Monterrey', '-05:00'),
(next_id('1_time_zones'), 'America/Montevideo', '-03:00'),
(next_id('1_time_zones'), 'America/Montserrat', '-04:00'),
(next_id('1_time_zones'), 'America/Nassau', '-04:00'),
(next_id('1_time_zones'), 'America/New_York', '-04:00'),
(next_id('1_time_zones'), 'America/Nipigon', '-04:00'),
(next_id('1_time_zones'), 'America/Nome', '-08:00'),
(next_id('1_time_zones'), 'America/Noronha', '-02:00'),
(next_id('1_time_zones'), 'America/North_Dakota/Beulah', '-05:00'),
(next_id('1_time_zones'), 'America/North_Dakota/Center', '-05:00'),
(next_id('1_time_zones'), 'America/North_Dakota/New_Salem', '-05:00'),
(next_id('1_time_zones'), 'America/Ojinaga', '-06:00'),
(next_id('1_time_zones'), 'America/Panama', '-05:00'),
(next_id('1_time_zones'), 'America/Pangnirtung', '-04:00'),
(next_id('1_time_zones'), 'America/Paramaribo', '-03:00'),
(next_id('1_time_zones'), 'America/Phoenix', '-07:00'),
(next_id('1_time_zones'), 'America/Port-au-Prince', '-04:00'),
(next_id('1_time_zones'), 'America/Port_of_Spain', '-04:00'),
(next_id('1_time_zones'), 'America/Porto_Velho', '-04:00'),
(next_id('1_time_zones'), 'America/Puerto_Rico', '-04:00'),
(next_id('1_time_zones'), 'America/Punta_Arenas', '-03:00'),
(next_id('1_time_zones'), 'America/Rainy_River', '-05:00'),
(next_id('1_time_zones'), 'America/Rankin_Inlet', '-05:00'),
(next_id('1_time_zones'), 'America/Recife', '-03:00'),
(next_id('1_time_zones'), 'America/Regina', '-06:00'),
(next_id('1_time_zones'), 'America/Resolute', '-05:00'),
(next_id('1_time_zones'), 'America/Rio_Branco', '-05:00'),
(next_id('1_time_zones'), 'America/Santarem', '-03:00'),
(next_id('1_time_zones'), 'America/Santiago', '-04:00'),
(next_id('1_time_zones'), 'America/Santo_Domingo', '-04:00'),
(next_id('1_time_zones'), 'America/Sao_Paulo', '-03:00'),
(next_id('1_time_zones'), 'America/Scoresbysund', 'UTC'),
(next_id('1_time_zones'), 'America/Sitka', '-08:00'),
(next_id('1_time_zones'), 'America/St_Barthelemy', '-04:00'),
(next_id('1_time_zones'), 'America/St_Johns', '-02:30'),
(next_id('1_time_zones'), 'America/St_Kitts', '-04:00'),
(next_id('1_time_zones'), 'America/St_Lucia', '-04:00'),
(next_id('1_time_zones'), 'America/St_Thomas', '-04:00'),
(next_id('1_time_zones'), 'America/St_Vincent', '-04:00'),
(next_id('1_time_zones'), 'America/Swift_Current', '-06:00'),
(next_id('1_time_zones'), 'America/Tegucigalpa', '-06:00'),
(next_id('1_time_zones'), 'America/Thule', '-03:00'),
(next_id('1_time_zones'), 'America/Thunder_Bay', '-04:00'),
(next_id('1_time_zones'), 'America/Tijuana', '-07:00'),
(next_id('1_time_zones'), 'America/Toronto', '-04:00'),
(next_id('1_time_zones'), 'America/Tortola', '-04:00'),
(next_id('1_time_zones'), 'America/Vancouver', '-07:00'),
(next_id('1_time_zones'), 'America/Whitehorse', '-07:00'),
(next_id('1_time_zones'), 'America/Winnipeg', '-05:00'),
(next_id('1_time_zones'), 'America/Yakutat', '-08:00'),
(next_id('1_time_zones'), 'America/Yellowknife', '-06:00'),
(next_id('1_time_zones'), 'Antarctica/Casey', '+08:00'),
(next_id('1_time_zones'), 'Antarctica/Davis', '+07:00'),
(next_id('1_time_zones'), 'Antarctica/DumontDUrville', '+10:00'),
(next_id('1_time_zones'), 'Antarctica/Macquarie', '+11:00'),
(next_id('1_time_zones'), 'Antarctica/Mawson', '+05:00'),
(next_id('1_time_zones'), 'Antarctica/McMurdo', '+12:00'),
(next_id('1_time_zones'), 'Antarctica/Palmer', '-03:00'),
(next_id('1_time_zones'), 'Antarctica/Rothera', '-03:00'),
(next_id('1_time_zones'), 'Antarctica/Syowa', '+03:00'),
(next_id('1_time_zones'), 'Antarctica/Troll', '+02:00'),
(next_id('1_time_zones'), 'Antarctica/Vostok', '+06:00'),
(next_id('1_time_zones'), 'Arctic/Longyearbyen', '+02:00'),
(next_id('1_time_zones'), 'Asia/Aden', '+03:00'),
(next_id('1_time_zones'), 'Asia/Almaty', '+06:00'),
(next_id('1_time_zones'), 'Asia/Amman', '+03:00'),
(next_id('1_time_zones'), 'Asia/Anadyr', '+12:00'),
(next_id('1_time_zones'), 'Asia/Aqtau', '+05:00'),
(next_id('1_time_zones'), 'Asia/Aqtobe', '+05:00'),
(next_id('1_time_zones'), 'Asia/Ashgabat', '+05:00'),
(next_id('1_time_zones'), 'Asia/Atyrau', '+05:00'),
(next_id('1_time_zones'), 'Asia/Baghdad', '+03:00'),
(next_id('1_time_zones'), 'Asia/Bahrain', '+03:00'),
(next_id('1_time_zones'), 'Asia/Baku', '+04:00'),
(next_id('1_time_zones'), 'Asia/Bangkok', '+07:00'),
(next_id('1_time_zones'), 'Asia/Barnaul', '+07:00'),
(next_id('1_time_zones'), 'Asia/Beirut', '+03:00'),
(next_id('1_time_zones'), 'Asia/Bishkek', '+06:00'),
(next_id('1_time_zones'), 'Asia/Brunei', '+08:00'),
(next_id('1_time_zones'), 'Asia/Chita', '+09:00'),
(next_id('1_time_zones'), 'Asia/Choibalsan', '+08:00'),
(next_id('1_time_zones'), 'Asia/Chongqing', '+08:00'),
(next_id('1_time_zones'), 'Asia/Colombo', '+05:30'),
(next_id('1_time_zones'), 'Asia/Damascus', '+03:00'),
(next_id('1_time_zones'), 'Asia/Dhaka', '+06:00'),
(next_id('1_time_zones'), 'Asia/Dili', '+09:00'),
(next_id('1_time_zones'), 'Asia/Dubai', '+04:00'),
(next_id('1_time_zones'), 'Asia/Dushanbe', '+05:00'),
(next_id('1_time_zones'), 'Asia/Famagusta', '+03:00'),
(next_id('1_time_zones'), 'Asia/Gaza', '+03:00'),
(next_id('1_time_zones'), 'Asia/Hebron', '+03:00'),
(next_id('1_time_zones'), 'Asia/Ho_Chi_Minh', '+07:00'),
(next_id('1_time_zones'), 'Asia/Hong_Kong', '+08:00'),
(next_id('1_time_zones'), 'Asia/Hovd', '+07:00'),
(next_id('1_time_zones'), 'Asia/Irkutsk', '+08:00'),
(next_id('1_time_zones'), 'Asia/Jakarta', '+07:00'),
(next_id('1_time_zones'), 'Asia/Jayapura', '+09:00'),
(next_id('1_time_zones'), 'Asia/Jerusalem', '+03:00'),
(next_id('1_time_zones'), 'Asia/Kabul', '+04:30'),
(next_id('1_time_zones'), 'Asia/Kamchatka', '+12:00'),
(next_id('1_time_zones'), 'Asia/Karachi', '+05:00'),
(next_id('1_time_zones'), 'Asia/Kathmandu', '+05:45'),
(next_id('1_time_zones'), 'Asia/Khandyga', '+09:00'),
(next_id('1_time_zones'), 'Asia/Kolkata', '+05:30'),
(next_id('1_time_zones'), 'Asia/Krasnoyarsk', '+07:00'),
(next_id('1_time_zones'), 'Asia/Kuala_Lumpur', '+08:00'),
(next_id('1_time_zones'), 'Asia/Kuching', '+08:00'),
(next_id('1_time_zones'), 'Asia/Kuwait', '+03:00'),
(next_id('1_time_zones'), 'Asia/Macau', '+08:00'),
(next_id('1_time_zones'), 'Asia/Magadan', '+11:00'),
(next_id('1_time_zones'), 'Asia/Makassar', '+08:00'),
(next_id('1_time_zones'), 'Asia/Manila', '+08:00'),
(next_id('1_time_zones'), 'Asia/Muscat', '+04:00'),
(next_id('1_time_zones'), 'Asia/Nicosia', '+03:00'),
(next_id('1_time_zones'), 'Asia/Novokuznetsk', '+07:00'),
(next_id('1_time_zones'), 'Asia/Novosibirsk', '+07:00'),
(next_id('1_time_zones'), 'Asia/Omsk', '+06:00'),
(next_id('1_time_zones'), 'Asia/Pontianak', '+07:00'),
(next_id('1_time_zones'), 'Asia/Pyongyang', '+09:00'),
(next_id('1_time_zones'), 'Asia/Qatar', '+03:00'),
(next_id('1_time_zones'), 'Asia/Qostanay', '+06:00'),
(next_id('1_time_zones'), 'Asia/Qyzylorda', '+05:00'),
(next_id('1_time_zones'), 'Asia/Riyadh', '+03:00'),
(next_id('1_time_zones'), 'Asia/Sakhalin', '+11:00'),
(next_id('1_time_zones'), 'Asia/Samarkand', '+05:00'),
(next_id('1_time_zones'), 'Asia/Seoul', '+09:00'),
(next_id('1_time_zones'), 'Asia/Shanghai', '+08:00'),
(next_id('1_time_zones'), 'Asia/Singapore', '+08:00'),
(next_id('1_time_zones'), 'Asia/Srednekolymsk', '+11:00'),
(next_id('1_time_zones'), 'Asia/Taipei', '+08:00'),
(next_id('1_time_zones'), 'Asia/Tashkent', '+05:00'),
(next_id('1_time_zones'), 'Asia/Tbilisi', '+04:00'),
(next_id('1_time_zones'), 'Asia/Tehran', '+04:30'),
(next_id('1_time_zones'), 'Asia/Thimphu', '+06:00'),
(next_id('1_time_zones'), 'Asia/Tokyo', '+09:00'),
(next_id('1_time_zones'), 'Asia/Tomsk', '+07:00'),
(next_id('1_time_zones'), 'Asia/Ulaanbaatar', '+08:00'),
(next_id('1_time_zones'), 'Asia/Urumqi', '+06:00'),
(next_id('1_time_zones'), 'Asia/Ust-Nera', '+10:00'),
(next_id('1_time_zones'), 'Asia/Vientiane', '+07:00'),
(next_id('1_time_zones'), 'Asia/Vladivostok', '+10:00'),
(next_id('1_time_zones'), 'Asia/Yakutsk', '+09:00'),
(next_id('1_time_zones'), 'Asia/Yangon', '+06:30'),
(next_id('1_time_zones'), 'Asia/Yekaterinburg', '+05:00'),
(next_id('1_time_zones'), 'Asia/Yerevan', '+04:00'),
(next_id('1_time_zones'), 'Atlantic/Azores', 'UTC'),
(next_id('1_time_zones'), 'Atlantic/Bermuda', '-03:00'),
(next_id('1_time_zones'), 'Atlantic/Canary', '+01:00'),
(next_id('1_time_zones'), 'Atlantic/Cape_Verde', '-01:00'),
(next_id('1_time_zones'), 'Atlantic/Faroe', '+01:00'),
(next_id('1_time_zones'), 'Atlantic/Madeira', '+01:00'),
(next_id('1_time_zones'), 'Atlantic/Reykjavik', 'UTC'),
(next_id('1_time_zones'), 'Atlantic/South_Georgia', '-02:00'),
(next_id('1_time_zones'), 'Atlantic/St_Helena', 'UTC'),
(next_id('1_time_zones'), 'Atlantic/Stanley', '-03:00'),
(next_id('1_time_zones'), 'Australia/Adelaide', '+09:30'),
(next_id('1_time_zones'), 'Australia/Brisbane', '+10:00'),
(next_id('1_time_zones'), 'Australia/Broken_Hill', '+09:30'),
(next_id('1_time_zones'), 'Australia/Currie', '+10:00'),
(next_id('1_time_zones'), 'Australia/Darwin', '+09:30'),
(next_id('1_time_zones'), 'Australia/Eucla', '+08:45'),
(next_id('1_time_zones'), 'Australia/Hobart', '+10:00'),
(next_id('1_time_zones'), 'Australia/Lindeman', '+10:00'),
(next_id('1_time_zones'), 'Australia/Lord_Howe', '+10:30'),
(next_id('1_time_zones'), 'Australia/Melbourne', '+10:00'),
(next_id('1_time_zones'), 'Australia/Perth', '+08:00'),
(next_id('1_time_zones'), 'Australia/Sydney', '+10:00'),
(next_id('1_time_zones'), 'Europe/Amsterdam', '+02:00'),
(next_id('1_time_zones'), 'Europe/Andorra', '+02:00'),
(next_id('1_time_zones'), 'Europe/Astrakhan', '+04:00'),
(next_id('1_time_zones'), 'Europe/Athens', '+03:00'),
(next_id('1_time_zones'), 'Europe/Belgrade', '+02:00'),
(next_id('1_time_zones'), 'Europe/Berlin', '+02:00'),
(next_id('1_time_zones'), 'Europe/Bratislava', '+02:00'),
(next_id('1_time_zones'), 'Europe/Brussels', '+02:00'),
(next_id('1_time_zones'), 'Europe/Bucharest', '+03:00'),
(next_id('1_time_zones'), 'Europe/Budapest', '+02:00'),
(next_id('1_time_zones'), 'Europe/Busingen', '+02:00'),
(next_id('1_time_zones'), 'Europe/Chisinau', '+03:00'),
(next_id('1_time_zones'), 'Europe/Copenhagen', '+02:00'),
(next_id('1_time_zones'), 'Europe/Dublin', '+01:00'),
(next_id('1_time_zones'), 'Europe/Gibraltar', '+02:00'),
(next_id('1_time_zones'), 'Europe/Guernsey', '+01:00'),
(next_id('1_time_zones'), 'Europe/Helsinki', '+03:00'),
(next_id('1_time_zones'), 'Europe/Isle_of_Man', '+01:00'),
(next_id('1_time_zones'), 'Europe/Istanbul', '+03:00'),
(next_id('1_time_zones'), 'Europe/Jersey', '+01:00'),
(next_id('1_time_zones'), 'Europe/Kaliningrad', '+02:00'),
(next_id('1_time_zones'), 'Europe/Kiev', '+03:00'),
(next_id('1_time_zones'), 'Europe/Kirov', '+03:00'),
(next_id('1_time_zones'), 'Europe/Lisbon', '+01:00'),
(next_id('1_time_zones'), 'Europe/Ljubljana', '+02:00'),
(next_id('1_time_zones'), 'Europe/London', '+01:00'),
(next_id('1_time_zones'), 'Europe/Luxembourg', '+02:00'),
(next_id('1_time_zones'), 'Europe/Madrid', '+02:00'),
(next_id('1_time_zones'), 'Europe/Malta', '+02:00'),
(next_id('1_time_zones'), 'Europe/Mariehamn', '+03:00'),
(next_id('1_time_zones'), 'Europe/Minsk', '+03:00'),
(next_id('1_time_zones'), 'Europe/Monaco', '+02:00'),
(next_id('1_time_zones'), 'Europe/Moscow', '+03:00'),
(next_id('1_time_zones'), 'Europe/Oslo', '+02:00'),
(next_id('1_time_zones'), 'Europe/Paris', '+02:00'),
(next_id('1_time_zones'), 'Europe/Podgorica', '+02:00'),
(next_id('1_time_zones'), 'Europe/Prague', '+02:00'),
(next_id('1_time_zones'), 'Europe/Riga', '+03:00'),
(next_id('1_time_zones'), 'Europe/Rome', '+02:00'),
(next_id('1_time_zones'), 'Europe/Samara', '+04:00'),
(next_id('1_time_zones'), 'Europe/San_Marino', '+02:00'),
(next_id('1_time_zones'), 'Europe/Sarajevo', '+02:00'),
(next_id('1_time_zones'), 'Europe/Saratov', '+04:00'),
(next_id('1_time_zones'), 'Europe/Simferopol', '+03:00'),
(next_id('1_time_zones'), 'Europe/Skopje', '+02:00'),
(next_id('1_time_zones'), 'Europe/Sofia', '+03:00'),
(next_id('1_time_zones'), 'Europe/Stockholm', '+02:00'),
(next_id('1_time_zones'), 'Europe/Tallinn', '+03:00'),
(next_id('1_time_zones'), 'Europe/Tirane', '+02:00'),
(next_id('1_time_zones'), 'Europe/Ulyanovsk', '+04:00'),
(next_id('1_time_zones'), 'Europe/Uzhgorod', '+03:00'),
(next_id('1_time_zones'), 'Europe/Vaduz', '+02:00'),
(next_id('1_time_zones'), 'Europe/Vatican', '+02:00'),
(next_id('1_time_zones'), 'Europe/Vienna', '+02:00'),
(next_id('1_time_zones'), 'Europe/Vilnius', '+03:00'),
(next_id('1_time_zones'), 'Europe/Volgograd', '+04:00'),
(next_id('1_time_zones'), 'Europe/Warsaw', '+02:00'),
(next_id('1_time_zones'), 'Europe/Zagreb', '+02:00'),
(next_id('1_time_zones'), 'Europe/Zaporozhye', '+03:00'),
(next_id('1_time_zones'), 'Europe/Zurich', '+02:00'),
(next_id('1_time_zones'), 'Indian/Antananarivo', '+03:00'),
(next_id('1_time_zones'), 'Indian/Chagos', '+06:00'),
(next_id('1_time_zones'), 'Indian/Christmas', '+07:00'),
(next_id('1_time_zones'), 'Indian/Cocos', '+06:30'),
(next_id('1_time_zones'), 'Indian/Comoro', '+03:00'),
(next_id('1_time_zones'), 'Indian/Kerguelen', '+05:00'),
(next_id('1_time_zones'), 'Indian/Mahe', '+04:00'),
(next_id('1_time_zones'), 'Indian/Maldives', '+05:00'),
(next_id('1_time_zones'), 'Indian/Mauritius', '+04:00'),
(next_id('1_time_zones'), 'Indian/Mayotte', '+03:00'),
(next_id('1_time_zones'), 'Indian/Reunion', '+04:00'),
(next_id('1_time_zones'), 'Pacific/Apia', '+13:00'),
(next_id('1_time_zones'), 'Pacific/Auckland', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Bougainville', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Chatham', '+12:45'),
(next_id('1_time_zones'), 'Pacific/Chuuk', '+10:00'),
(next_id('1_time_zones'), 'Pacific/Easter', '-06:00'),
(next_id('1_time_zones'), 'Pacific/Efate', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Enderbury', '+13:00'),
(next_id('1_time_zones'), 'Pacific/Fakaofo', '+13:00'),
(next_id('1_time_zones'), 'Pacific/Fiji', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Funafuti', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Galapagos', '-06:00'),
(next_id('1_time_zones'), 'Pacific/Gambier', '-09:00'),
(next_id('1_time_zones'), 'Pacific/Guadalcanal', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Guam', '+10:00'),
(next_id('1_time_zones'), 'Pacific/Honolulu', '-10:00'),
(next_id('1_time_zones'), 'Pacific/Kiritimati', '+14:00'),
(next_id('1_time_zones'), 'Pacific/Kosrae', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Kwajalein', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Majuro', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Marquesas', '-09:30'),
(next_id('1_time_zones'), 'Pacific/Midway', '-11:00'),
(next_id('1_time_zones'), 'Pacific/Nauru', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Niue', '-11:00'),
(next_id('1_time_zones'), 'Pacific/Norfolk', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Noumea', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Pago_Pago', '-11:00'),
(next_id('1_time_zones'), 'Pacific/Palau', '+09:00'),
(next_id('1_time_zones'), 'Pacific/Pitcairn', '-08:00'),
(next_id('1_time_zones'), 'Pacific/Pohnpei', '+11:00'),
(next_id('1_time_zones'), 'Pacific/Port_Moresby', '+10:00'),
(next_id('1_time_zones'), 'Pacific/Rarotonga', '-10:00'),
(next_id('1_time_zones'), 'Pacific/Saipan', '+10:00'),
(next_id('1_time_zones'), 'Pacific/Tahiti', '-10:00'),
(next_id('1_time_zones'), 'Pacific/Tarawa', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Tongatapu', '+13:00'),
(next_id('1_time_zones'), 'Pacific/Wake', '+12:00'),
(next_id('1_time_zones'), 'Pacific/Wallis', '+12:00');
`