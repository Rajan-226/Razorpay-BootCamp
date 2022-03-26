<?php
include_once "BasicUser.php";
include_once "StandardUser.php";
include_once "ProUser.php";

function dummyDataTest($user)
{
    $user->createNewList("Grocery");

    for ($i = 0; $i < 5; $i++) {
        $user->addTask(0, "Task", rand(1, 20));
    }

    $user->displayAllLists();

    $user->sortByImportance(0);

    print "After Sorting\n\n";
    $user->displayALlLists();

    for ($i = 0; $i < 100; $i++) {
        if (!$user->createNewList("Something")) {
            return;
        }
    }
}

$proUser = new ProUser("Ketan");
dummyDataTest($proUser);

//Turn on to test them

//$proUser = new StandardUser("Anish");
//dummyDataTest($proUser);

//$proUser = new BasicUser("Rajan");
//dummyDataTest($proUser);

