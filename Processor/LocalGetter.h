#ifndef LOCALGETTER_H
#define LOCALGETTER_H

#include "IFileGetter.h"
#include <filesystem>
#include <iostream>
#include <string>

class LocalGetter : public IFileGetter 
{
private:
    std::ifstream fileStream;  // Cached file stream
    std::string cachedFilePath;

public:
    ~LocalGetter() override;

    void getFile() override;
    std::string getFileName() override;
    std::ifstream& retrieveFile() override;
    void removeFile() override;
    void cleanupCache() override;
};

#endif // LOCALGETTER_H
