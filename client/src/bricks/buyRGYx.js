import {
  Box,
  Button,
  Form,
  FormField,
  Heading,
  Select,
  TextInput
} from "grommet";
import React, { useState } from "react";

export default props => {
  const [form, setForm] = useState({
    investors: props.investors,
    rgys: props.rgys
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
        <FormField label="Investor" name="investor" required>
          <Select
            options={form.investors}
            value={form.investor}
            onChange={e => setForm({ ...form, investor: e.option })}
          />
        </FormField>
        <FormField label="RGYx" name="rgy" required>
          <Select
            options={form.rgys}
            value={form.rgy}
            onChange={e => setForm({ ...form, rgy: e.option })}
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
