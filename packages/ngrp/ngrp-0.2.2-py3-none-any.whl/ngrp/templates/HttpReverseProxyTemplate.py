from ngrp.templates.Template import Template


class HttpReverseProxyTemplate(Template):
    """\
    server {
        listen 80;
        server_name <domain>;

        location / {
            proxy_pass_header Authorization;
            proxy_pass http://127.0.0.1:<port>;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_http_version 1.1;
            proxy_set_header Connection "";
            proxy_buffering off;
            client_max_body_size 0;
            proxy_read_timeout 36000s;
            proxy_redirect off;
        }
    }
    """
