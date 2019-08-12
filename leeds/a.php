<?PHP
$d = "2019-01-01";
$ts = strtotime($d);
var_dump($d, $ts, date("Y-m=d", $ts));
