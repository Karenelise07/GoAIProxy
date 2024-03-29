let currentBot = '';
let currentBotAvatar = '';
const userAvatar = './user.jpg'; // 假设用户头像URL
const botConversations = {
  'Bot 1': [],
  'Bot 2': [],
  'Bot 3': []
};

function selectBot(botName, botAvatar, index) {
  currentBot = botName;
  currentBotAvatar = botAvatar;

  // 更新选中的机器人样式
  const botListItems = document.getElementById('botList').children;
  for (let i = 0; i < botListItems.length; i++) {
    botListItems[i].classList.remove('selected');
  }
  botListItems[index].classList.add('selected');

  // 更新当前机器人信息区域
  document.getElementById('currentBotName').textContent = botName;
  document.getElementById('currentBotAvatar').src = botAvatar;

  // 清空对话并加载当前选中的机器人对话历史
  const messagesContainer = document.getElementById('messages');
  messagesContainer.innerHTML = ''; // 清空当前对话
  const conversations = botConversations[currentBot] || [];
  conversations.forEach(message => {
    if (message.sender === 'user') {
      messagesContainer.innerHTML += `<div class="message user"><div class="text">${message.text}</div><img class="avatar" src="${userAvatar}" alt="User"></div>`;
    } else {
      messagesContainer.innerHTML += `<div class="message bot"><img class="avatar" src="${currentBotAvatar}" alt="${currentBot}"><div class="text">${message.text}</div></div>`;
    }
  });

  // 滚动到最新消息
  messagesContainer.scrollTop = messagesContainer.scrollHeight;
}


function sendMessage() {
  const input = document.getElementById('userInput');
  const message = input.value.trim();
  if (message && currentBot) {
    const userMessage = { sender: 'user', text: message };
    botConversations[currentBot].push(userMessage);
    
    // 显示用户消息
    const messagesContainer = document.getElementById('messages');
    appendMessage(messagesContainer, 'user', message, userAvatar);
    
    // 发送消息到服务器并处理回复
    sendToBotServer(message);

    input.value = ''; // 清空输入框
  }
}

function sendToBotServer(userMessage) {
  // 定义请求的数据结构
  const requestData = {
    model: "qwen1.5-chat", // 假设您要使用的模型名称
    // stream: true,
    // frequency_penalty: 0,
    // presence_penalty: 0,
    // temperature: 0.6,
    // top_p: 1,
    messages: [{
      content: userMessage,
      role: "user"
    }]
  };

  fetch('http://localhost:8083/v1/chat/completions', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(requestData), // 使用新的请求数据结构
  })
  .then(response => response.json())
  .then(data => {
    const botReply = { sender: 'bot', text: data.response };
    botConversations[currentBot].push(botReply);
    
    // 显示机器人回复
    const messagesContainer = document.getElementById('messages');
    appendMessage(messagesContainer, 'bot', botReply.text, currentBotAvatar);

    // 滚动到最新消息
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
  })
  .catch(error => {
    console.error('Error sending message to bot:', error);
  });
}

function appendMessage(container, sender, text, avatar) {
  let messageHTML = '';
  if (sender === 'user') {
    messageHTML = `<div class="message user"><div class="text">${text}</div><img class="avatar" src="${userAvatar}" alt="User"></div>`;
  } else { // 假设sender为'bot'
    messageHTML = `<div class="message bot"><img class="avatar" src="${avatar}" alt="Bot"><div class="text">${text}</div></div>`;
  }
  container.innerHTML += messageHTML;
}


function init() {
  // 监听回车键发送消息
  document.getElementById('userInput').addEventListener('keypress', function(event) {
    if (event.key === 'Enter') {
      event.preventDefault(); // 阻止默认行为
      sendMessage();
    }
  });
}

// 在页面加载完成后执行init函数
window.onload = init;
