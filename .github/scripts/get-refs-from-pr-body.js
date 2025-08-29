/**
 * Configuration for repositories and their default refs
 * Add new repos here to automatically include them in the ref extraction
 */
const REPO_CONFIG = {
  core: {
    pattern: /core ref:\s*(\S+)/i,
    defaultRef: "develop",
    outputKey: "core-ref",
  },
  // Add new config here:
  // <name>: {
  //   pattern: /<name> ref:\s*(\S+)/i,
  //   defaultRef: "develop",
  //   outputKey: "<name>-ref"
  // }
};

/**
 * Validates a Git reference (branch name, tag, or commit SHA).
 *
 * These git refs are extracted from PR body text and must be validated for safety.
 *
 * @param {string} ref - The Git reference to validate
 * @returns {string} - The validated reference
 * @throws {Error} - If the reference is invalid
 */
function validateGitRef(ref) {
  if (!ref) return null;

  // Check length (Git refs can't be longer than 255 chars in practice)
  if (ref.length > 255) {
    throw new Error(`Git ref too long (${ref.length} chars): ${ref}`);
  }

  // Check if the ref is a valid commit SHA
  const isCommitSHA = /^[0-9a-f]{4,40}$/i.test(ref);
  if (isCommitSHA) {
    return ref;
  }

  // Check if the ref is a valid SemVer tag
  const isSemVer =
    /^v?(?:0|[1-9]\d*)\.(?:0|[1-9]\d*)\.(?:0|[1-9]\d*)(?:-(?:[0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*)?$/.test(
      ref,
    );
  if (isSemVer) {
    return ref;
  }

  // Finally, check if it's a valid branch/tag name, which is the fallback
  const isBranchOrTag =
    /^(?!.*\.lock$)(?!.*\.\.)(?!.*\/\/)(?!.*@\{)(?!.*[ ~^:\?\*\[\\])(?!\.)(?!.*\.$)(?!.*\/\.)[A-Za-z0-9][A-Za-z0-9._\/-]*[A-Za-z0-9]$/.test(
      ref,
    );
  if (isBranchOrTag) {
    // We add an extra check here: if it looks like a version but failed the SemVer check, reject it.
    if (/^v?\d+\.\d+/.test(ref)) {
      throw new Error(
        `Invalid SemVer format for version-like tag: ${ref}. It contains leading zeros or other formatting errors.`,
      );
    }
    return ref;
  }

  // If none of the patterns match, it's invalid.
  throw new Error(
    `Invalid Git reference format: ${ref}. Must be a valid branch name, semver tag, or commit SHA.`,
  );

  // Additional safety checks for dangerous sequences not covered by patterns above
  if (ref.includes("\\")) {
    throw new Error(`Git reference contains backslashes: ${ref}`);
  }

  return ref;
}

/**
 * Extracts Git references from PR body text using the configured patterns
 * @param {string} body - The PR body text
 * @returns {Object} - Object containing matched references for each repo
 */
function extractRefsFromBody(body) {
  const refs = {};

  for (const [repoName, config] of Object.entries(REPO_CONFIG)) {
    const match = body.match(config.pattern);
    refs[repoName] = match?.[1];
  }

  return refs;
}

/**
 * Validates and processes a single git ref
 * @param {string} refName - Name of the ref type (for error messages)
 * @param {string|undefined} refValue - The ref value to validate (may be undefined)
 * @param {string} defaultRef - Default value if ref is not provided
 * @param {Object} core - GitHub Actions core object
 * @returns {string|null} - Validated ref or null if validation failed
 */
function processRef(refName, refValue, defaultRef, core) {
  // If no ref was specified in the PR body, use the default
  if (!refValue) {
    core.info(`No ${refName} ref specified, using default: ${defaultRef}`);
    return defaultRef;
  }

  try {
    const validatedRef = validateGitRef(refValue);
    core.info(`Using custom ${refName} ref: ${validatedRef}`);
    return validatedRef;
  } catch (error) {
    core.setFailed(`Invalid ${refName} ref: ${error.message}`);
    return null;
  }
}

/**
 * Main function for the GitHub Action
 * @param {object} { github, context, core }
 */
module.exports = async ({ github, context, core }) => {
  try {
    const pr = await github.rest.pulls.get({
      owner: context.repo.owner,
      repo: context.repo.repo,
      pull_number: context.issue.number,
    });

    const body = pr.data.body || "";
    core.info(`PR Body: ${body}`);

    // If the PR body is empty, use defaults for all repos and exit early
    if (!body.trim()) {
      core.info("PR body is empty. Using default refs for all repositories.");
      for (const [, config] of Object.entries(REPO_CONFIG)) {
        core.setOutput(config.outputKey, config.defaultRef);
      }
      core.info("All default refs set. Exiting.");
      return;
    }

    const extractedRefs = extractRefsFromBody(body);
    core.info(`Extracted refs: ${JSON.stringify(extractedRefs)}`);

    // Log what was found in the PR body
    const foundRefs = [];
    for (const [repoName, refValue] of Object.entries(extractedRefs)) {
      if (refValue) {
        foundRefs.push(`${repoName}: ${refValue}`);
      }
    }

    if (foundRefs.length > 0) {
      core.info(`Found custom refs in PR body: ${foundRefs.join(", ")}`);
    } else {
      core.info(
        "No custom refs found in PR body, using defaults for all repos",
      );
    }

    // Process and validate each ref dynamically
    const processedRefs = {};
    let hasValidationError = false;

    for (const [repoName, config] of Object.entries(REPO_CONFIG)) {
      const refValue = extractedRefs[repoName];
      const processedRef = processRef(
        repoName,
        refValue,
        config.defaultRef,
        core,
      );

      if (processedRef === null) {
        hasValidationError = true;
        break;
      }

      processedRefs[repoName] = processedRef;
    }

    // If any ref validation failed, the processRef function will have called core.setFailed
    if (hasValidationError) {
      return;
    }

    // Set outputs dynamically using core.setOutput
    for (const [repoName, config] of Object.entries(REPO_CONFIG)) {
      core.setOutput(config.outputKey, processedRefs[repoName]);
    }

    // Log final refs
    const finalRefsList = Object.entries(processedRefs)
      .map(([repo, ref]) => `${repo}: ${ref}`)
      .join(", ");
    core.info(`Final refs - ${finalRefsList}`);

    core.info("All refs processed successfully.");
  } catch (error) {
    core.setFailed(`Action failed with error: ${error.message}`);
  }
};
