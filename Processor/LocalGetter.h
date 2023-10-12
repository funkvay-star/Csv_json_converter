#ifndef LOCALGETTER_H
#define LOCALGETTER_H

#include "IFileGetter.h"
#include <fstream>
#include <string>

class LocalGetter : public IFileGetter {
private:
    std::ifstream fileStream;  // Cached file stream
    std::string cachedFilePath;  // Path of the cached file

public:
    ~LocalGetter() override;

    void getFile() override;
    std::ifstream& retrieveFile() override;
    void removeFile() override;
    void cleanupCache() override;
};

#endif // LOCALGETTER_H
