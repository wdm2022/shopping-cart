FROM mongo:4.4

RUN openssl rand -base64 756 > /keyFile.key \
    && chown mongodb:mongodb /keyFile.key \
    && chmod 400 /keyFile.key
