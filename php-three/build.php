<?php
$pagesDir = "./pages";
$outputDir = "./output";
$frameDir = "./frame";

// Create output directory if it doesn't exist
if (!is_dir($outputDir)) {
    mkdir($outputDir);
}

// Process all files in the pages directory
$pages = array_filter(scandir($pagesDir), fn($file) => pathinfo($file, PATHINFO_EXTENSION) === "html");

function create_navigation($pages)
{
    $nav = "<nav>\n";
    foreach ($pages as $page) {
        $nav .= "<a href=\"$page\">$page</a>\n";
    }
    $nav .= "</nav>\n";
    return $nav;
}
foreach ($pages as $page) {
    $pageContent = file_get_contents("$pagesDir/$page");
    ob_start();
    include "$frameDir/head.html";
    include "$frameDir/header.html";

    echo create_navigation($pages);
    echo $pageContent;
    include "$frameDir/footer.html";
    $fullPage = ob_get_clean();

    // Write the result to the output folder
    file_put_contents("$outputDir/$page", $fullPage);
    echo "Generated $page\n";
}
