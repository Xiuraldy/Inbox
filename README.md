# **Inbox - Email Indexing & Search**

## 📌 Descripción

Este proyecto permite indexar y buscar correos electrónicos desde una base de datos de correos electrónicos de Enron utilizando **Go**, **Vue 3** y **OpenObserve**.

## 🚀 Tecnologías utilizadas

- **Backend:** Go con Chi
- **Base de Datos:** OpenObserve
- **Frontend:** Vue 3 + Vite
- **Estilos:** TailwindCSS
- **Almacenamiento Local:** OpenObserve

## ⚡ Instalación

Clona el repositorio e instala las dependencias necesarias:

```sh
git clone https://github.com/Xiuraldy/Inbox.git
cd Inbox
```

### 🔧 Configurar la API

Antes de iniciar, asegúrate de tener **Go** instalado. Luego, ejecuta:

```sh
cd api
go mod tidy
```

### 🔧 Configurar la App

Navega al directorio de la aplicación e instala las dependencias de Node.js:

```sh
cd app
npm install
```

## 🚀 Ejecución

### **1️⃣ Levantar OpenObserve**

Ejecuta **OpenObserve** de forma local con permisos de administrador:

```sh
openobserve.exe
```

### **2️⃣ Iniciar la API**

Corre el servidor de Go con:

```sh
go run main.go
```

### **3️⃣ Iniciar la App**

Corre la aplicación de Vue con:

```sh
npm run dev
```

## 📡 API Endpoints

Estos son los principales endpoints del proyecto:

| Método   | Endpoint                       | Descripción                                 |
| -------- | ------------------------------ | ------------------------------------------- |
| **POST** | `/email?from=0&search=keyword` | Retorna los emails filtrados.               |
| **POST** | `/total?from=0&search=keyword` | Retorna el número total de emails filtrados |
| **POST** | `/indexing`                    | Sube los correos a OpenObserve.             |

---

Si necesitas más información, contáctame en **[LinkedIn](https://www.linkedin.com/in/xiuraldy/)**. 🚀
