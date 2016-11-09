FROM hyperledger/fabric-peer:latest
WORKDIR $GOPATH/src/github.com/
RUN mkdir branciard
WORKDIR $GOPATH/src/github.com/branciard
RUN git clone https://github.com/branciard/paza.git
WORKDIR $GOPATH/src/github.com/branciard/paza/chaincode/sample
RUN go build
WORKDIR $GOPATH/src/github.com/branciard/paza/chaincode/asset
RUN go build
WORKDIR /user/branciard/paza/
RUN npm install hfc
