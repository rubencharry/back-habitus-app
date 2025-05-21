-- Datos de prueba: hábitos
INSERT INTO habits (title, description, frequency) VALUES
('Leer 30 minutos', 'Dedicar media hora a la lectura diaria', 'daily'),
('Ejercicio', 'Realizar actividad física', 'daily'),
('Meditar', '5 minutos de meditación', 'daily'),
('Beber 2 litros de agua', 'Mantenerse hidratado', 'daily'),
('Escribir en el diario', 'Reflexionar al final del día', 'daily');

-- Datos de prueba: tareas
INSERT INTO tasks (title, description, due_date, completed) VALUES
('Preparar parcial', 'Estudiar los temas vistos en clase', '2025-05-12', false),
('Sacar al perro', 'Caminar al perro en la tarde', '2025-05-10', false),
('Comprar víveres', 'Ir al supermercado', '2025-05-08', false),
('Enviar informe de progreso', 'Enviar avance al equipo', '2025-05-14', false),
('Llamar a mamá', 'Hacer una llamada', '2025-05-09', false);