#include "LocalGetter.h"

LocalGetter::~LocalGetter() 
{
    cleanupCache();
}

void LocalGetter::getFile() 
{
    const std::string directoryPath = "../BackendAndFrontend/Files";
    for (const auto &entry : std::filesystem::directory_iterator(directoryPath)) 
    {
        if (entry.is_regular_file()) 
        {
            cachedFilePath = entry.path().string();
            fileStream.open(cachedFilePath, std::ios::in);
            break;  // Only get one file
        }
    }
}

std::string LocalGetter::getFileName() 
{
    // Check if cachedFilePath is not empty
    if (cachedFilePath.empty()) 
    {
        std::cerr << "Error: File path is empty!" << std::endl;
        return;
    }

    // Find the position of the last directory separator
    size_t lastSeparator = cachedFilePath.find_last_of("/\\");

    // Find the position of the last period
    size_t lastPeriod = cachedFilePath.rfind('.');

    if (lastPeriod == std::string::npos || (lastSeparator != std::string::npos && lastPeriod < lastSeparator))
    {
        std::cerr << "Error: File extension not found or incorrect!" << std::endl;
        return;
    }

    // Extract the file name
    std::string fileName;
    if (lastSeparator != std::string::npos) 
    {
        fileName = cachedFilePath.substr(lastSeparator + 1, lastPeriod - lastSeparator - 1);
    }
    else 
    {
        fileName = cachedFilePath.substr(0, lastPeriod);
    }

    // Output the file name
    return fileName;
}


std::ifstream& LocalGetter::retrieveFile() 
{
    return fileStream;
}

void LocalGetter::removeFile() 
{
    if (!cachedFilePath.empty()) 
    {
        fileStream.close();  // Close the stream before deleting the file
        std::filesystem::remove(cachedFilePath);
        cachedFilePath.clear();
    }
}

void LocalGetter::cleanupCache() 
{
    if (fileStream.is_open()) 
    {
        fileStream.close();
    }
    cachedFilePath.clear();
}
