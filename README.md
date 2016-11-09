# paza
pAza chain


Setup environnement
===================

- prerequis
    -   docker installed
    -   understand  hyperledger starter kit here :
    http://hyperledger-fabric.readthedocs.io/en/latest/starter/fabric-starter-kit/
    
- run it
    - go into docker directory
    - create the sample image by tapping :
    docker build --no-cache=true -t branciard/paza-sample .
    - go into docker/sample directory et type :
   docker-compose up -d
    - check the the 3 containers with 
    docker ps
     - launch on one terminal to check peer logs:
     docker logs -f peer
     - launch on one terminal to check membersrvc logs:
      docker logs -f paza-sample
     - launch on one terminal :
      docker exec -it paza-sample /bin/bash
      node app-sample
     it works if :Query Response:{"Name":"b","Amount":"21"}
    
