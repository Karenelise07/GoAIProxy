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
  const requestData = {
    model: "qwen1.5-chat",
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
    body: JSON.stringify(requestData),
  }).then(response => {
    console.log('Received response from server');
    const reader = response.body.getReader();
    const decoder = new TextDecoder();
    let accumulatedContent = ''; // 用于累积对话内容

    function read() {
      reader.read().then(({done, value}) => {
        if (done) {
          console.log('Stream completed');
          // 在这里处理累积的内容 accumulatedContent
          // 因为流结束了，我们需要处理最后一次累积的数据
          processAccumulatedContent(accumulatedContent);
          return;
        }
        const chunk = decoder.decode(value, {stream: true});
        accumulatedContent += chunk; // 累积从流中读取的内容
        read(); // 继续读取下一个数据块
      }).catch(error => {
        console.error('Error reading from stream:', error);
      });
    }
    read();
  }).catch(error => {
    console.error('Error sending message to bot:', error);
  });
}

function processAccumulatedContent(accumulatedContent) {
  // 首先，移除结束标记之后的所有内容（如果存在）
  const endIndex = accumulatedContent.indexOf('data: [DONE]');
  if (endIndex !== -1) {
    // 如果找到结束标记，只保留结束标记之前的内容
    accumulatedContent = accumulatedContent.substring(0, endIndex);
  }
 
  let jsonStrings = accumulatedContent.split(/data: (?={)/);
  jsonStrings = jsonStrings.filter(str => str.trim().startsWith("{"));

  let fullMessage = ''; // 用于累积完整的消息内容

  jsonStrings.forEach(jsonStr => {
    console.log(jsonStr); // 输出当前尝试解析的字符串
    try {
      let jsonObj = JSON.parse(jsonStr);
      // 检查是否有内容需要累积
      if (jsonObj.choices && jsonObj.choices.length > 0) {
        const choice = jsonObj.choices[0];
        if (choice.delta && choice.delta.content) {
          // 累积内容
          fullMessage += choice.delta.content;
        }
        // 检查是否是对话结束的信号
        if (choice.finish_reason === 'stop') {
          // 如果是结束信号，展示累积的完整消息
          const botReply = { sender: 'bot', text: fullMessage };
          botConversations[currentBot].push(botReply);
          const messagesContainer = document.getElementById('messages');
          appendMessage(messagesContainer, 'bot', botReply.text, currentBotAvatar);
          messagesContainer.scrollTop = messagesContainer.scrollHeight;

          // 重置fullMessage以准备下一轮对话
          fullMessage = '';
        }
      }
    } catch (error) {
      console.error('Error parsing JSON chunk:', error);
    }
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
