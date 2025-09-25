# Proyecto 1 - Sistemas Operativos 1  
**Desarrollo de Contenedores y Gestión de Imágenes en Entornos Virtualizados**  

Estudiante: **Kevin Castañeda**  
Carnet: **201901801**  

---

## Descripción del Proyecto
Este proyecto implementa un entorno contenerizado con **3 APIs en Go**, distribuidas en diferentes VMs y gestionadas con **Docker y Containerd**, además de un **registro privado Zot**.  

Cada API expone endpoints REST y responde con mensajes en formato **JSON** que identifican la API, la VM donde se ejecuta y los datos del estudiante.  

---

## Estructura del Proyecto
![Estructura del proyecto](/201901801_LAB_SO1_2S2025/proyecto1/img/image1.png)


---

## Tecnologías utilizadas
- **Lenguaje**: Go 1.24  
- **Contenerización**: Docker, Containerd  
- **Registry privado**: Zot  
- **Sistema operativo base**: Ubuntu 22.04 LTS  

---

## Endpoints de las APIs
Cada API expone **dos endpoints** que permiten comunicación cruzada.  

### API1 (VM1)
- `GET /api1/201901801/llamar-api2`  
- `GET /api1/201901801/llamar-api3`  

### API2 (VM1)
- `GET /api2/201901801/llamar-api1`  
- `GET /api2/201901801/llamar-api3`  

### API3 (VM2)
- `GET /api3/201901801/llamar-api1`  
- `GET /api3/201901801/llamar-api2`  

---

## Guía de Instalación y Ejecución

### 1. Clonar repositorio
```bash
git clone https://github.com/tuusuario/201901801_LAB_SO1_2S2025.git
cd 201901801_LAB_SO1_2S2025/proyecto1
```

### 2. Instalar dependencias (Ubuntu 22.04)
```
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```
### 3. Construir imagenes
```
cd Api1
docker build -t api1-vm1 .

cd ../Api2
docker build -t api2-vm1 .

cd ../Api3
docker build -t api3-vm2 .
```
### 4. Ejecutar contenedores
```
docker run -d -p 8081:8081 --name api1 api1-vm1
docker run -d -p 8082:8082 --name api2 api2-vm1
docker run -d -p 8083:8083 --name api3 api3-vm2
```
### 5. Probar Apis
```
curl http://localhost:8081/api1/201901801/llamar-api2
curl http://localhost:8082/api2/201901801/llamar-api3
curl http://localhost:8083/api3/201901801/llamar-api1
```

### 6. Respuestas
```
{
  "mensaje": "Hola, responde la API: API1 en la VM1, desarrollada por el estudiante Kevin Castañeda con carnet: 201901801"
}
```
---
## Publicación en Zot (VM3)
### 1. Etiquetar imagenes antes de subir
```
docker tag api1-vm1 localhost:5000/API1-VM1:1.0
docker tag api2-vm1 localhost:5000/API2-VM1:1.0
docker tag api3-vm2 localhost:5000/API3-VM2:1.0
```
### 2. Subir imagenes al registro privado Zot
```
docker push localhost:5000/API1-VM1:1.0
docker push localhost:5000/API2-VM1:1.0
docker push localhost:5000/API3-VM2:1.0
```