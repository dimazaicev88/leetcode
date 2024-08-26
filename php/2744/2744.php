<?php

class Solution
{

    /**
     * @param String[] $words
     * @return Integer
     */
    function maximumNumberOfStringPairs($words)
    {
        $count = 0;
        $mapWords = [];
        foreach ($words as $word) {
            if (isset($mapWords[$word])) {
                $count++;
            }
            $mapWords[strrev($word)] = 1;
        }

        return $count;
    }
}

echo (new Solution())->maximumNumberOfStringPairs(["cd","ac","dc","ca","zz"]);
