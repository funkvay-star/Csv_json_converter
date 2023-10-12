#ifndef GCPGETTER_H
#define GCPGETTER_H

#include "IFileGetter.h"

class GCPGetter : public IFileGetter 
{
private:
    std::ifstream fileStream;  // Cached file stream

public:
    ~GCPGetter() override;

    void getFile() override;
    std::ifstream& retrieveFile() override;
    void removeFile() override;
    void cleanupCache() override;
};

#endif // GCPGETTER_H
