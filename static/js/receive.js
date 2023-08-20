
const qrcode1 = QRCreator('Привет, Мир!',
{ mode: 4,
  eccl: 0,
  version: 3,
  mask: -1,
  image: 'html',
  modsize: -1,
  margin: 0
});

const content = (qrcode) =>{
  return qrcode.error ?
    `недопустимые исходные данные ${qrcode.error}`:
     qrcode.result;
};

document.getElementById('qrcode').append(QRCreator('15PbLQR1M2Hbvo2gYVrN4BzDeN9yXLMuFb').result);