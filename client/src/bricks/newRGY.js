import { Box, Button, Form, FormField, Heading, TextInput } from "grommet";
import React, { useState } from "react";

export default props => {
  const [form, setForm] = useState({
    devID: props.devID
  });
  return (
    <Box margin="medium">
      <Box margin={{ bottom: "small" }} pad={{ horizontal: "small" }}>
        <Heading level="2">CREATE NEW DEVELOPER</Heading>
      </Box>
      <Form
        onSubmit={e => (props.onSubmit ? props.onSubmit(e.value) : null)}
        value={form}
      >
        <FormField label="Dev ID" name="devID" required>
          <TextInput
            placeholder="Developer ID"
            value={form.devID}
            onChange={e => setForm({ ...form, devID: e.target.value })}
          />
        </FormField>
        <FormField label="Name" name="name" required>
          <TextInput
            placeholder="Your Token name (e.g. RGYX)"
            value={form.name}
            onChange={e => setForm({ ...form, name: e.target.value })}
          />
        </FormField>
        <FormField label="Number of shares" name="shares" required>
          <TextInput
            placeholder="How shares are you need?"
            value={form.share}
            type="number"
            onChange={e => setForm({ ...form, shares: e.target.value })}
          />
        </FormField>
        <FormField label="Cost by share" name="cost" required>
          <TextInput
            placeholder="Set your cost in cent dollars"
            value={form.cost}
            type="number"
            onChange={e => setForm({ ...form, cost: e.target.value })}
          />
        </FormField>
        <Box margin={{ top: "large" }} direction="row" justify="between">
          <Button label="close" onClick={props.onClose} />
          <Button
            type="submit"
            primary
            label="Submit"
            margin={{ left: "large" }}
          />
        </Box>
      </Form>
    </Box>
  );
};
