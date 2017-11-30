<?php

// ---------------------
// target PHP 7.2
// ---------------------

echo count([1]); // 1
echo count([1, 2, 3]); // 3
echo count([1, null]); // 2
echo count(["one" => 1, "two" => 2]); // 2

echo count(null); // Warning
echo count(new StdClass); // Warning
echo count(1); // Warning
echo count("1"); // Warning

class FullName implements Countable
{
    private $firstName;
    private $middleName;
    private $lastName;

    public function __construct(?string $firstName, ?string $middleName, ?string $lastName)
    {
        $this->firstName = $firstName;
        $this->middleName = $middleName;
        $this->lastName = $lastName;
    }

    public function count() : int
    {
        $count = 0;
        foreach (get_object_vars($this) as $val) {
            if ($val) {
                $count++;
            }
        }
        return $count;
    }
}

// 3
echo count(new FullName("hogeta", "hogehoge", "hogeo"));
