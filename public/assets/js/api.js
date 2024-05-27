async function getTeamCollection() {
    try {
        const response = await fetch('/team/collections', {
          method: 'GET',
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        return await response.text();
      } catch (error) {
        console.error('Error at fetch user collection:', error);
        // Handle errors
    }
}

async function getUserCollection(userId) {
    try {
        const response = await fetch('/my/collection/'+userId, {
          method: 'GET',
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        return await response.text();
      } catch (error) {
        console.error('Error at fetch user collection:', error);
        // Handle errors
    }
}

function authProvider(provider) {
  const searchParams = new URLSearchParams();
  searchParams.set("provider", provider)
  window.location.href = "/v1/auth/signin?" + searchParams.toString()
}