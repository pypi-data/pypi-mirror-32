
# Manual Gluu Server Clustering

# This documentation is a developmental work in progress to optimize the process. Please refer to the official [Gluu Clustering Docs](https://gluu.org/docs/ce/3.0.2/installation-guide/cluster/) for clustering.

## Introduction
If you have requirements for high availability (HA) or failover, you can configure your Gluu Server for multi-master replication by following the instructions below.

## Prerequisites

Some prerequisites are necessary for setting up Gluu with delta-syncrepl MMR:

- A minimum of three (3) servers or VMs--two (2) for Gluu Servers and one (1) for load balancing (in our example, NGINX);
- To create the following instructions we used Ubuntu 14 Trusty, but the process should not be OS specific;
- To create the following instructions we used an Nginx load balancer/proxy, however if you have your own load balancer, like F5 or Cisco, you should use that instead and disregard the bottom instructions about configuring Nginx.
- Gluu Server 3.x using OpenLDAP.

## Concept

Multi-master replication with OpenLDAP through delta-syncrepl by creating an accesslog database and configuring synchronization by means the slapd.conf file. The ldap.conf file for all the servers will allow the self-signed certs that Gluu creates and configuring the symas-openldap.conf to allow external connections for LDAP. There are also some additional steps that are required to persist Gluu functionality across servers. This is where a load-balancer/proxy are required.

## Instructions

### 1. [Install Gluu](https://gluu.org/docs/ce/3.0.2/installation-guide/install/) on ALL of the servers making sure to use a separate NGINX server FQDN as hostname and not the current servers hostname.

- A separate NGINX server is recommended, but not necessary, since replicating a Gluu server to a different hostname breaks the functionality of the Gluu web page, when using a hostname other than what is in the certificates. For example, if I used c1.gluu.info as my host and copied that to a second server (e.g. c2.gluu.info), the process of accessing the site on c2.gluu.info, even with replication, will fail authentication, due to hostname conflict. So if c1 failed, you couldn't access the Gluu web GUI anymore.

### 2. Copy the necessary files to every other server

- Inside the Gluu chroot, navigate to `/etc/gluu/conf/` and edit `ox-ldap.properties` replacing:

`servers: localhost:1636`

With:

`servers: localhost:1636,{insert server1 FQDN here}:1636,{insert server2 FQDN here}:1636,...`

Placing all servers in your cluster topology in this config portion.

- Now tar the `/etc/gluu/conf/` folder to the other servers and replace the existing `/etc/gluu/conf/`

```
Gluu.Root # cd /etc/gluu/
Gluu.Root # tar -cvf gluu.gz conf/
Gluu.Root # scp gluu.gz root@server2.com:/opt/gluu-server-3.0.2/etc/gluu/
```

Server 2 (Other servers, as well, for your replication strategy)

```
Gluu.Root # cd /etc/gluu/
Gluu.Root # rm -rf conf/
Gluu.Root # tar -xvf gluu.gz
```

- Make sure the directory structure here is `/etc/gluu/conf`


### 3. There needs to be primary server to replicate from initially for delta-syncrepl to inject data from. After the initial sync, all servers will be exactly the same, as delta-syncrepl will fill the newly created database.

- So choose one server as a base and then on every other server:

```
Gluu.Root # service solserver stop
Gluu.Root # rm /opt/gluu/data/main_db/*.mdb
Gluu.Root # rm /opt/gluu/data/site_db/*.mdb
```

- Now make accesslog directories on **every** servers and give ldap ownership:

```
Gluu.Root # mkdir /opt/gluu/data/accesslog_db
Gluu.Root # chown -R ldap. /opt/gluu/data/
```

### 4. Now is where we will set servers to associate with each other for MMR by editing the slapd.conf, ldap.conf and symas-openldap.conf files.

- Creating the slapd.conf file is relatively easy, but can be prone to errors if done manually. Attached is a script and template files for creating multiple slapd.conf files for every server. Download git and clone the necessary files on one server:

```
Gluu.Root # apt-get update && apt-get install git && cd /tmp/ && git clone https://github.com/GluuFederation/cluster-mgr.git && cd /tmp/cluster-mgr/manual_install/slapd_conf_script/
```

- We need to change the configuration file for our own specific needs:

```
Gluu.Root # vi syncrepl.cfg
```

- Here we want to change the `ip_address`, `fqn_hostname`, `ldap_password` to our specific server instances. For example:

```
[server_1]
ip_address = 192.168.30.133
fqn_hostname = server1.com
ldap_password = (your password)
enable = Yes

[server_2]
ip_address = 192.168.30.130
fqn_hostname = server2.com
ldap_password = (your password)
enable = Yes

[server_3]
...

[nginx]
fqn_hostname = nginx.server.com
```

- The `fqn_hostname` under the `[server_x]` needs to be the FQDN of the Gluu server. 

- The `create_slapd_conf.py` script now creates the NGINX configuration file for you as nginx.conf

- If required, you can change the `/tmp/cluster-mgr/manual_install/slapd_conf_script/ldap_templates/slapd.conf` to fit your specific needs to include different schemas, indexes, etc. Avoid changing any of the `{#variables#}`.

- Now run the python script `create_slapd_conf.py` (Built with python 2.7) in the `/tmp/cluster-mgr/manual_install/slapd_conf_script/` directory :

```
Gluu.Root # python create_slapd_conf.py
```

- There is also a 2.6 Python script included.

- This will output multiple `.conf` files in `/tmp/cluster-mgr/manual_install/slapd_conf_script/` named to match your server FQDN:

```
Gluu.Root #  ls
... server1_com.conf  server2_com.conf ...
```

- Move each .conf file to their respective server replacing the slapd.conf:

```
Gluu.Root # mv server1_com.conf /opt/symas/etc/openldap/slapd.conf
```

- and for the other servers

```
Gluu.Root # scp server_example_com.conf root@server.example.com:/opt/gluu-server-3.0.2/opt/symas/etc/openldap/slapd.conf
```

- Now create and modify the ldap.conf **on every server**:

```
Gluu.Root # vi /opt/symas/etc/openldap/ldap.conf
```

- Add these lines

```
TLS_CACERT /etc/certs/openldap.pem
TLS_REQCERT never
```

- Modify the HOST_LIST entry of symas-openldap.conf **on every server**:

```
vi /opt/symas/etc/openldap/symas-openldap.conf
```

- Replace:

```
HOST_LIST="ldaps://127.0.0.1:1636/"
```

- With:

```
HOST_LIST="ldaps://0.0.0.0:1636/ ldaps:///"
```

### 5. It is important that our servers times are synchronized so we must install ntp outside of the Gluu chroot and set ntp to update by the minute (necessary for delta-sync log synchronization). If time gets out of sync, the entries will conflict and their could be issues with replication.

```
GLUU.root@host:/ # logout
# apt install ntp
# crontab -e
```

- Select your preferred editor and add this to the bottom of the file:

```
* * * * * /usr/sbin/ntpdate -s time.nist.gov
```
 
- This synchronizes the time every minute.

- Force-reload solserver on every server

```
# service gluu-server-3.0.2 login
# service solserver force-reload
```

- Delta-sync multi-master replication should be initializing and running. Check the logs for confirmation. It might take a moment for them to sync, but you should end up see something like the following:

```
# tail -f /var/log/openldap/ldap.log | grep sync

Aug 23 22:40:29 dc4 slapd[79544]: do_syncrep2: rid=001 cookie=rid=001,sid=001,csn=20170823224029.216104Z#000000#001#000000
Aug 23 22:40:29 dc4 slapd[79544]: syncprov_matchops: skipping original sid 001
Aug 23 22:40:29 dc4 slapd[79544]: syncrepl_message_to_op: rid=001 be_modify
```

### 6. **If you have your own load balancer, you are done here.** If not, let's configure our NGINX server for oxTrust and oxAuth web failover.

- We need the httpd.crt and httpd.key certs from one of the Gluu servers.   

- From the NGINX server:  

```
mkdir /etc/nginx/ssl/
scp root@server1.com:/opt/gluu-server-3.0.2/etc/certs/httpd.key /etc/nginx/ssl/
scp root@server1.com:/opt/gluu-server-3.0.2/etc/certs/httpd.crt /etc/nginx/ssl/
```

### 7. Next we install, clear the nginx.conf file and configure NGINX to proxy-pass connections.  

```
apt-get install nginx -y
cd /etc/nginx/
>nginx.conf
vi nginx.conf
```

- Put the following template in it's place. Make sure to change the `{serverX_ip_or_FQDN}` portion to your servers IP addresses or FQDN under the upstream section. Add as many servers as exist in your replication setup. The `server_name` needs to be your NGINX servers FQDN.    

```
events {
        worker_connections 768;
}

http {
  upstream backend_id {
    ip_hash;
    server {server1_ip_or_FQDN}:443;
    server {server2_ip_or_FQDN}:443;
  }
  upstream backend {
    server {server1_ip_or_FQDN}:443;
    server {server2_ip_or_FQDN}:443;
        
  }
  server {
    listen       80;
    server_name  {NGINX_server_FQDN};
    return       301 https://{NGINX_server_FQDN}$request_uri;
   }
  server {
    listen 443;
    server_name {NGINX_server_FQDN};

    ssl on;
    ssl_certificate         /etc/nginx/ssl/httpd.crt;
    ssl_certificate_key     /etc/nginx/ssl/httpd.key;

    location ~ ^(/)$ {
      proxy_pass https://backend;
    }
    location /.well-known {
        proxy_pass https://backend/.well-known;
    }
    location /oxauth {
        proxy_pass https://backend/oxauth;
    }
    location /identity {
        proxy_pass https://backend_id/identity;
    }
    location /cas {
        proxy_pass https://backend/cas;
    }
    location /asimba {
        proxy_pass https://backend/asimba;
    }
    location /passport {
        proxy_pass https://backend/passport;
    }
  }
}

```

- Now all traffic for the Gluu web GUI will route through one address e.g. `nginx.gluu.info`. This gives us failover redundancy for our Gluu web GUI if any server goes down, as NGINX automatically does passive health checks.   

## Support
If you have any questions or run into any issues, please open a ticket on [Gluu Support](https://support.gluu.org).
