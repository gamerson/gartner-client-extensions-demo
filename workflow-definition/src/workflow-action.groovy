import com.liferay.portal.kernel.json.JSONFactoryUtil;
import java.net.URL;
import org.osgi.framework.FrameworkUtil;
import com.liferay.portal.workflow.kaleo.runtime.scripting.internal.util.KaleoScriptingEvaluator;

def bc = FrameworkUtil.getBundle(KaleoScriptingEvaluator.class).getBundleContext()

def oauth2ApplicationLS =bc.getService(bc.getServiceReference("com.liferay.oauth2.provider.service.OAuth2ApplicationLocalService"))

def oAuth2Application = oauth2ApplicationLS.fetchOAuth2ApplicationByExternalReferenceCode(
	Long.valueOf(companyId), "machine-workflow-action-headless-server")

def localOAuthClient = bc.getService(bc.getServiceReference("com.liferay.oauth.client.LocalOAuthClient"))

def rootJSON = JSONFactoryUtil.createJSONObject()

def token = ""

this.binding.variables.each {k,v ->
	println("$k = $v")

	if (k == "workflowTaskAssignees") {
		v.each {
			if (it.assigneeClassName == "com.liferay.portal.kernel.model.User") {
				String requestTokens = localOAuthClient.requestTokens(
					Long.valueOf(it.assigneeClassPK), oAuth2Application)

				if (requestTokens) {
					def requestTokensJSONObject = JSONFactoryUtil.createJSONObject(
						requestTokens);

					token = "Bearer ${requestTokensJSONObject.getString("access_token")}";
				}
			}
		}
	}

	if ((v instanceof Number) || (v instanceof String)) {
		rootJSON.put(k, v)

		return;
	}

	rootJSON.put(k, JSONFactoryUtil.createJSONObject(JSONFactoryUtil.serialize(v)))
}

def post = new URL("https://coupon-function-springboot.localdev.me/workflow/action").openConnection();
post.setRequestMethod("POST")
post.setDoOutput(true)
post.setRequestProperty("Content-Type", "application/json")

if (token) {
	post.setRequestProperty("Authorization", token)
}

post.getOutputStream().write(rootJSON.toJSONString().getBytes("UTF-8"));
def postRC = post.getResponseCode();
println(postRC);
if (postRC.equals(200)) {
    println(post.getInputStream().getText());
}