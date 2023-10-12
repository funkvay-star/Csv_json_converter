#include "LocalGetter.h"
#include <filesystem>

LocalGetter::~LocalGetter() 
{
    cleanupCache();
}

void LocalGetter::getFile() {
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
