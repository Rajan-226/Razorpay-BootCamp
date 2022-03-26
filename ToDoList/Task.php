<?php

class Task
{
    //importance value of a task
    public $criticalValue;

    //description of a todo
    public $description;

    function __construct($description, $criticalValue = 0)
    {
        $this->description = $description;
        $this->criticalValue = $criticalValue;
    }

    function getDescription(): string
    {
        return $this->description;
    }

    function setDescription(string $description): void
    {
        $this->description = $description;
    }

    function getCriticalValue()
    {
        return $this->criticalValue;
    }

    function setCriticalValue($criticalValue): void
    {
        $this->criticalValue = $criticalValue;
    }

}
