import {
  Box,
  Button,
  Form,
  FormField,
  Grid,
  Heading,
  Layer,
  Text,
  TextInput
} from "grommet";
import React, { useState } from "react";

export default props => {
  const [form, setForm] = useState({});
  return (
    <Box margin="medium">
      <Box margin={{ bottom: "small" }}>
        <Heading level="2">
          {props.title ? props.title : "CREATE NEW DEVELOPER"}
        </Heading>
      </Box>
      <Form
        onSubmit={e => (props.onSubmit ? props.onSubmit(e.value) : null)}
        value={form}
      >
        <FormField label="Name" name="name" required>
          <TextInput
            placeholder="Type your Name"
            value={form.name}
            onChange={e => setForm({ ...form, name: e.target.value })}
          />
        </FormField>
        <FormField label="Address" name="address" required>
          <TextInput
            placeholder="Type your ETH Address"
            value={form.address}
            onChange={e => setForm({ ...form, address: e.target.value })}
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
