#ifndef IFILEGETTER_H
#define IFILEGETTER_H

#include <fstream>

class IFileGetter 
{
public:
    virtual ~IFileGetter() 
    { 
        cleanupCache(); 
    }

    // Pure virtual method to open a file stream to a single file
    virtual void getFile() = 0;

    // Pure virtual method to retrieve the opened file stream
    virtual std::ifstream& retrieveFile() = 0;

    // Pure virtual method to remove the fetched file
    virtual void removeFile() = 0;

    // Pure virtual method to clean up any cached data (like closing the file stream)
    virtual void cleanupCache() = 0;
};

#endif // IFILEGETTER_H
