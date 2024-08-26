function maximumNumberOfStringPairs(words: string[]): number {
    const map = new Map<string, boolean>();
    let count = 0;
    for (const word of words) {
        if (map.has(word)) {
            count++;
        }
        map.set(word.split("").reverse().join(""), true)
    }

    return count;
};

console.log(maximumNumberOfStringPairs(["cd","ac","dc","ca","zz"]));