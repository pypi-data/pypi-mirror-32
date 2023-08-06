<?php

/*
	PHP合并DOCX文档脚本
*/

error_reporting(0);

require "vendor/autoload.php";
use DocxMerge\DocxMerge;

$dm = new DocxMerge();

// 使用方法 php merge_docx.php 需要合并的文档 输出的文档
unset($argv[0]);
$saved_docx = array_pop($argv);
// print_r($argv);
// print_r($saved_docx);

$dm->merge(array_values($argv), $saved_docx);
echo 'DONE'
// $dm->merge(['template.docx', 'template1.docx', 'template2.docx'], 'result.docx');
?>
