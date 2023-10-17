#ifndef CSV_TO_JSON_H
#define CSV_TO_JSON_H

#include <fstream>
#include <sstream>
#include <vector>
#include <string>
#include <stdexcept>
#include <algorithm>

class CSVtoJSON 
{
private:
    std::ifstream csvFile;

public:
    // Constructor taking an ifstream by move semantics
    CSVtoJSON(std::ifstream&& file);

    // Convert the CSV file to JSON format
    std::string convert();

private:
    // Static utility function to remove newlines and carriage returns from a string
    static void removeNewlinesAndCarriageReturns(std::string &s);

    // Static utility function to split a string by a given delimiter
    static std::vector<std::string> split(const std::string &s, char delimiter);

    // Static utility function to trim whitespace from the start and end of a string
    static void trim(std::string &s);
};

#endif // CSV_TO_JSON_H