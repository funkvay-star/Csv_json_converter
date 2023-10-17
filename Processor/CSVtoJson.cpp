#include "CSVtoJson.h"

CSVtoJSON::CSVtoJSON(std::ifstream&& file) 
    : csvFile(std::move(file)) 
{
    if (!csvFile.is_open()) 
    {
        throw std::runtime_error("Failed to open CSV file.");
    }
}

std::string CSVtoJSON::convert() 
{
    std::string line;
    std::vector<std::string> headers;
    bool isHeader = true;

    std::ostringstream jsonOutput;
    jsonOutput << "[\n";

    while (std::getline(csvFile, line)) 
    {
        std::vector<std::string> values = split(line, ',');

        if (isHeader) 
        {
            headers = values;
            isHeader = false;
            continue;
        }

        jsonOutput << "  {\n";
        for (size_t i = 0; i < values.size(); i++) 
        {
            if (i != 0) 
                jsonOutput << ",\n";

            std::string key = headers[i];
            std::string value = values[i];

            removeNewlinesAndCarriageReturns(key);
            trim(key);
            removeNewlinesAndCarriageReturns(value);
            trim(value);

            jsonOutput << "    \"" << key << "\": \"" << value << "\"";
        }   
        jsonOutput << "\n  },\n";
    }

    std::string result = jsonOutput.str();
    if (result.substr(result.length() - 3) == ",\n") 
    {
        result = result.substr(0, result.length() - 2);
    }

    result += "\n]";
    return result;
}

void CSVtoJSON::removeNewlinesAndCarriageReturns(std::string &s) 
{
    s.erase(std::remove(s.begin(), s.end(), '\n'), s.end());
    s.erase(std::remove(s.begin(), s.end(), '\r'), s.end());
}

std::vector<std::string> CSVtoJSON::split(const std::string &s, char delimiter) 
{
    std::vector<std::string> tokens;
    std::string item;
    bool insideQuotes = false;

    for (char ch : s) 
    {
        if (ch == '\"')
            insideQuotes = !insideQuotes;
        else if (ch == delimiter && !insideQuotes) 
        {
            tokens.push_back(item);
            item.clear();
        } 
        else 
        {
            item += ch;
        }
    }
    if (!item.empty()) 
        tokens.push_back(item); // Add the last token

    return tokens;
}

void CSVtoJSON::trim(std::string &s) 
{
    size_t start = s.find_first_not_of(" \t");
    size_t end = s.find_last_not_of(" \t");

    if (start == std::string::npos || end == std::string::npos) 
    {
        s = "";
        return;
    }

    s = s.substr(start, end - start + 1);
}
