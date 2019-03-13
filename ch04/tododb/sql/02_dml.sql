INSERT INTO todo(`title`, `content`, `status`, `created`, `updated`) VALUES
('MySQL 도커 이미지 만들기','MySQL 마스터와 슬레이브를 환경 변수로 설정할 수 있는 MySQL 이미지 생성','DONE', NOW(), NOW()),
('MySQL 스택 만들기','MySQL 마스터 및 슬레이브 서비스로 구성된 스택을 스웜 클러스터에 구축한다', 'DONE', NOW(), NOW()),
('API구현하기', 'Go 언어로 TODO를 확인, 수정할 수 있는 API 구현', 'PROGRESS', NOW(), NOW()),
('Nginx 도커 이미지 만들기','HTTP 요청을 백엔드로 전달하는 Nginx 이미지 만들기', 'PROGRESS', NOW(), NOW()),
('API 스택 구축하기','스웜에 Nginx와 API로 구성된 스택을 구축', 'PROGRESS', NOW(), NOW()),
('웹 앱 구현하기','Nuxt.js를 통해 API와 연동되는 웹 애플리케이션 구현', 'PROGRESS', NOW(), NOW()),
('웹 앱 스택 구축','스웜에 Nginx와 웹 앱으로 구성되는 스택을 구축', 'PROGRESS', NOW(), NOW()),
('인그레스 구축하기','외부에서 스웜 클러스터에 접근하게 해주는 인그레스 구축', 'TODO', NOW(), NOW());

