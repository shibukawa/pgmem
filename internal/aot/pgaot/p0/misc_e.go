package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EOH_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	m.T0[v5].(func(*base.Module, int32, int32, int32))(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ER_get_flat_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v8 != int32(2249) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27&int32(17) == int32(1) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(0) <= v11 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = F_expanded_record_fetch_tupdesc(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v21 = v14
	goto L6
L6:
	;
	F_assign_record_type_typmod(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v21 = v17
	goto L6
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v24
	goto L1
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	return v33
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v35 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v27&int32(4) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v158 = v35
	goto L15
L15:
	;
	return v158
L16:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	v45 = v27
	goto L18
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v45&int32(16) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v45 = v44
	goto L18
L20:
	;
	if int32(0) < v46 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v107 = v46
	goto L22
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v113 = int32(24)
	if v107 <= int32(0) {
		v145 = v113
		v147 = v2
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v56 = int32(0)
	v57 = v46
	goto L26
L24:
	;
	v96 = v46
	v101 = v45
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v101 & int32(-17)
	v107 = v96
	goto L22
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v56))))
	if v64 != 0 {
		v88 = v57
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v96 = v88
	v101 = v93
	goto L25
L28:
	;
	v91 = v56 + int32(1)
	if v91 < v88 {
		v56 = v91
		v57 = v88
		goto L26
	} else {
		goto L34
	}
L29:
	;
	v67 = v47 + int32(20) + v56<<(uint(int32(4))%32)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+6)))
	if v68 != 0 {
		v88 = v57
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v69 != int32(65535) {
		v88 = v57
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v56<<(uint(int32(2))%32))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v77 != int32(1) {
		v88 = v57
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v80 = int32(1)
	v82 = int32(0)
	F_expanded_record_set_field_internal(m, l0, v56+v80, v76, v82, v80, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v88 = v87
	goto L28
L34:
	;
	goto L27
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v149 = F_heap_compute_data_size(m, v47, v148, v112)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L43
	}
L36:
	;
	v118 = int32(0)
	goto L37
L37:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v112))))
	if v125 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v135 = base.I32_div_s(v131+int32(7), int32(8))
	v145 = (v135 + int32(30)) & int32(-8)
	v147 = int32(1)
	goto L35
L39:
	;
	v129 = v118 + int32(1)
	if v107 != v129 {
		v118 = v129
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v145 = v113
	v147 = v2
	goto L35
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)) = uint8(v147)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v149
	v154 = v149 + v145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v154
	v158 = v154
	goto L15
}
func F_ER_mc_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v3 == int32(0) {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v6
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		if v8 <= v6 {
			return
		} else {
			v12 = v8 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = v12
			if v12 != 0 {
				return
			} else {
				F_FreeTupleDesc(m, v3)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_ExecARDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l4 == int32(0) {
		if v8 != 0 {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
			if v17 != 0 {
				v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if l3 == int32(0) {
						v27 = int32(0)
						v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v38 = int32(0)
							v40 = int32(1)
							F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = int32(0)
							v40 = int32(1)
							F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			} else {
				if l4 == int32(0) {
					return
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
					if v20 != int32(1) {
						return
					} else {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			if l4 == int32(0) {
				return
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
				if v20 != int32(1) {
					return
				} else {
					v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						if l3 == int32(0) {
							v27 = int32(0)
							v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
		if v11 == int32(0) {
			if v8 != 0 {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
				if v17 != 0 {
					v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						if l3 == int32(0) {
							v27 = int32(0)
							v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = int32(0)
								v40 = int32(1)
								F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				} else {
					if l4 == int32(0) {
						return
					} else {
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						if v20 != int32(1) {
							return
						} else {
							v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								if l3 == int32(0) {
									v27 = int32(0)
									v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				if l4 == int32(0) {
					return
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
					if v20 != int32(1) {
						return
					} else {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
			if v14 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						F_errmsg(m, int32(176232), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(519958), int32(2817), int32(143411))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v8 != 0 {
					v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+19)))
					if v17 != 0 {
						v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if l3 == int32(0) {
								v27 = int32(0)
								v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							} else {
								F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = int32(0)
									v40 = int32(1)
									F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					} else {
						if l4 == int32(0) {
							return
						} else {
							v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
							if v20 != int32(1) {
								return
							} else {
								v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
								mBase = m.M
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									if l3 == int32(0) {
										v27 = int32(0)
										v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return
										} else {
											v38 = int32(0)
											v40 = int32(1)
											F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												return
											}
										}
									} else {
										F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											v38 = int32(0)
											v40 = int32(1)
											F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l4 == int32(0) {
						return
					} else {
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
						if v20 != int32(1) {
							return
						} else {
							v23 = F_ExecGetTriggerOldSlot(m, l0, l1)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return
							} else {
								if l3 == int32(0) {
									v27 = int32(0)
									v33 = F_GetTupleForTrigger(m, l0, v27, l1, l2, int32(3), v23, v27, v27, v27, v27)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								} else {
									F_ExecForceStoreHeapTuple(m, l3, v23, int32(0))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = int32(0)
										v40 = int32(1)
										F_AfterTriggerSaveEvent(m, l0, l1, v38, v38, v40, v40, v23, v38, v38, v38, l4, l5)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecARUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L33
	}
L2:
	;
	if v13 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v16 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+1)))
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+2)))
	if v20 == int32(1) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	return
L8:
	;
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)))
	if v23 != 0 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if l8 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+1)))
	if v26 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l8)+2)))
	if v27 != int32(1) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	v30 = l2
	goto L18
L17:
	;
	v30 = l1
	goto L18
L18:
	;
	v31 = F_ExecGetTriggerOldSlot(m, l0, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	if l5 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v57 = F_ExecGetAllUpdatedCols(m, l1, l0)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L31
	}
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	m.T0[v52].(func(*base.Module, int32))(m, v31)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L30
	}
L23:
	;
	if l4 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_ExecForceStoreHeapTuple(m, l5, v31, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L29
	}
L26:
	;
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
	if v37 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v40 = int32(0)
	v46 = F_GetTupleForTrigger(m, l0, v40, v30, l4, int32(3), v31, v40, v40, v40, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	goto L21
L29:
	;
	goto L21
L30:
	;
	goto L21
L31:
	;
	F_AfterTriggerSaveEvent(m, l0, l1, l2, l3, int32(2), int32(1), v31, l6, l7, v57, l8, l9)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	goto L7
L33:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(176232), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(519958), int32(3164), int32(143453))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecBSDeleteTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v19 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L22
	}
L2:
	;
	m.G0 = v9 + int32(48)
	return
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+21)))
	if v22 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	v28 = F_before_stmt_triggers_fired(m, v26, int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v28 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(38654706106)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v34 <= int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v41 = v3
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v46 = v43 + v41*int32(60)
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
	if v47&int32(75) != int32(10) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L2
L11:
	;
	v74 = v41 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v74 < v75 {
		v41 = v74
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v53 = int32(0)
	v56 = F_TriggerEnabled(m, l0, l1, v46, v52, v53, v53, v53)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v56 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v46
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v65 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v68 = v65
	goto L17
L16:
	;
	v66 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v70 = F_ExecCallTriggerFunc(m, v9+int32(4), v41, v63, v64, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	v68 = v66
	goto L17
L19:
	;
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	goto L10
L22:
	;
	F_errcode(m, int32(16908867))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(364445), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(519958), int32(2677), int32(143390))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecCloseIndices(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v20 = v13 << (uint(int32(2)) % 32)
	v21 = v11 + v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10+v20)))
	F_index_insert_cleanup(m, v22, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_relation_close(m, v27, int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(0)
	v34 = v13 + int32(1)
	if v34 != v7 {
		v13 = v34
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_ExecGetRootToChildMap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
	if v7 == int32(0) {
		v10 = int32(4562080)
		v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
		v26 = F_build_attrmap_by_name_if_req(m, v16, v13, (v21^int32(-1))&int32(1))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v26 != 0 {
				v30 = F_convert_tuples_by_name_attrmap(m, v16, v13, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v30
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v35)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					return v42
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
				v35 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v35)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
				return v42
			}
		}
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		return v42
	}
}
func F_ExecIRDeleteTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v14 = F_ExecGetTriggerOldSlot(m, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(90194313658)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v28
	F_ExecForceStoreHeapTuple(m, l2, v14, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v34 <= int32(0) {
		v91 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v11 + int32(48)
	return v91
L5:
	;
	v44 = int32(0)
	goto L6
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v48 = v45 + v44*int32(60)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+12)))
	if v49&int32(75) != int32(73) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v91 = v82
	goto L4
L8:
	;
	v82 = int32(1)
	v84 = v44 + v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v84 < v85 {
		v44 = v84
		goto L6
	} else {
		goto L22
	}
L9:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v55 = int32(0)
	v57 = F_TriggerEnabled(m, l0, l1, v48, v54, v55, v14, v55)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v57 == int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v14
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v68 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = v68
	goto L14
L13:
	;
	v69 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	v73 = F_ExecCallTriggerFunc(m, v11+int32(4), v44, v66, v67, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v71 = v69
	goto L14
L16:
	;
	if v73 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v91 = int32(0)
	goto L4
L18:
	;
	goto L19
L19:
	;
	if l2 == v73 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v73)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L8
L22:
	;
	goto L7
}
func F_ExecInitGenerated(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v4 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L14
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+17)))
	if v27 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+18)))
	if v30 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l2 != int32(2) {
		v42 = int32(0)
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v43 = int32(4562080)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
	v50 = F_palloc0(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L16
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+13)))
	if v38 != 0 {
		v42 = int32(0)
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v39 = F_ExecGetUpdatedCols(m, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	return
L15:
	;
	v42 = v39
	goto L8
L16:
	;
	if int32(0) < v26 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l2 == int32(2) {
		goto L56
	} else {
		goto L57
	}
L18:
	;
	v60 = int32(0)
	v62 = v4
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_pfree(m, v50)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L54
	}
L21:
	;
	v74 = v60 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(110)+v75<<(uint(int32(4))%32)+v60*int32(100)))))
	if v82 == int32(0) {
		v163 = v62
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v163 != 0 {
		v191 = v163
		v192 = v50
		goto L17
	} else {
		goto L53
	}
L23:
	;
	if v74 != v26 {
		v60 = v74
		v62 = v163
		goto L21
	} else {
		goto L52
	}
L24:
	;
	v85 = F_build_column_default(m, v21, v74)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	if v85 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v42 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
	F_pull_varattnos(m, v85, int32(1), v19+int32(12))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v82 == int32(115) {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v97 = int32(0)
	if v42 == v97 {
		v138 = v97
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v138 == int32(0) {
		v163 = v62
		goto L23
	} else {
		goto L45
	}
L32:
	;
	goto L31
L33:
	;
	if v96 == int32(0) {
		v138 = v97
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v106 < v107 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v109 = v106
	goto L37
L36:
	;
	v109 = v107
	goto L37
L37:
	;
	if v109 <= int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v112 = int32(1)
	goto L40
L39:
	;
	v112 = v109
	goto L40
L40:
	;
	v113 = int32(8)
	v118 = int32(0)
	goto L41
L41:
	;
	v125 = v118 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v96+v113+v125)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v42+v113))))
	v130 = v127 & v129
	v132 = base.B2i32(v130 != int32(0))
	if v130 != 0 {
		v138 = v132
		goto L32
	} else {
		goto L43
	}
L42:
	;
	v138 = v132
	goto L32
L43:
	;
	v134 = v118 + int32(1)
	if v134 != v112 {
		v118 = v134
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L29
L46:
	;
	v149 = F_ExecPrepareExpr(m, v85, l1)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L49
	}
L47:
	;
	v154 = v62
	goto L48
L48:
	;
	if l2 != int32(2) {
		v163 = v154
		goto L23
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50+v60<<(uint(int32(2))%32)))) = v149
	v154 = v62 + int32(1)
	goto L48
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v160 = F_bms_add_member(m, v157, v60+int32(8))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v160
	v163 = v154
	goto L23
L52:
	;
	goto L22
L53:
	;
	goto L20
L54:
	;
	v184 = int32(0)
	v191 = v184
	v192 = v184
	goto L17
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v44
	goto L2
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v192
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v206)
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v192
	goto L55
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v235 + int32(4)
	F_errmsg_internal(m, int32(753974), v19)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(524080), int32(479), int32(469303))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecInitNullTupleSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = F_MakeTupleTableSlot(m, l1, int32(1652372))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
		v11 = F_lappend(m, v10, v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
			m.T0[v15].(func(*base.Module, int32))(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v21 = v19 << (uint(int32(2)) % 32)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				if v22&int32(3) != 0 {
					v40 = v21
					v44 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v40)
					mBase = m.M
				} else {
					if base.Ui32(int32(1024)) < base.Ui32(v21) {
						v40 = v21
						v44 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v40)
						mBase = m.M
					} else {
						v27 = v21 + v22
						if base.Ui32(v27) <= base.Ui32(v22) {
						} else {
							v32 = v22 + int32(4)
							if base.Ui32(v32) < base.Ui32(v27) {
								v34 = v27
							} else {
								v34 = v32
							}
							v40 = (v22^int32(-1)+v34)&int32(-4) + int32(4)
							v44 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(int32(0)), v40)
							mBase = m.M
						}
					}
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
				v52 = F__emscripten_memset_bulkmem(m, v47, base.I32_extend8_s(int32(1)), v50)
				mBase = m.M
				v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
				v55 = v53 & int32(65533)
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v55)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v58)
				return v6
			}
		}
	}
}
func F_ExecMaterial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v15 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	if v14 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v42 = v15
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v47*int32(24))+4)))
	goto L20
L10:
	;
	v21 = int32(1)
	v53 = v21
	v54 = v2
	v55 = v21
	goto L6
L11:
	;
	goto L12
L12:
	;
	v23 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[323]))
	v28 = F_tuplestore_begin_heap(m, v23, int32(0), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_tuplestore_set_eflags(m, v28, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v33&int32(16) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v36 = F_tuplestore_alloc_read_pointer(m, v28, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v28
	if v28 == int32(0) {
		v53 = int32(1)
		v54 = v2
		v55 = v23
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v42 = v28
	goto L9
L20:
	;
	v53 = v51
	v54 = v42
	v55 = int32(0)
	goto L6
L21:
	;
	return v112
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	m.T0[v109].(func(*base.Module, int32))(m, v107)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L53
	}
L23:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v81 != 0 {
		v107 = v80
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v76 = F_tuplestore_gettupleslot(m, v54, base.B2i32(v14 == int32(1)), int32(0), v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v53 != 0 {
		v80 = v71
		goto L23
	} else {
		goto L33
	}
L26:
	;
	if v53 == int32(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v60 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v63 = int32(0)
	v65 = F_tuplestore_advance(m, v54, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v72 = v70
	goto L24
L31:
	;
	if v65 == int32(0) {
		v112 = v63
		goto L21
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v72 = v71
	goto L24
L34:
	;
	if v76 != 0 {
		v112 = v72
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v14 != int32(1) {
		v107 = v72
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v80 = v72
	goto L23
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	if v83 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_ExecReScan(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v87 = m.T0[v86].(func(*base.Module, int32) int32)(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	if v55 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v87 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v89&int32(2) == int32(0) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v94)
	return int32(0)
L47:
	;
	goto L46
L48:
	;
	F_tuplestore_puttupleslot(m, v54, v87)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+32))
	m.T0[v103].(func(*base.Module, int32, int32))(m, v80, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	return v80
L53:
	;
	v112 = v107
	goto L21
}
func F_ExecMaterializesOutput(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = l0 - int32(348)
	return base.B2i32(base.Ui32(v3) < base.Ui32(int32(15))) & int32(base.Ui32(int32(20541))>>(uint(v3)%32))
}
func F_ExecOpenIndices(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+116)))
	if v13 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_RelationGetIndexList(m, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v16 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = v20 << (uint(int32(2)) % 32)
	v25 = F_palloc(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v27 = F_palloc(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v32 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_list_free(m, v16)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L22
	}
L12:
	;
	v45 = v36 << (uint(int32(2)) % 32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)))
	v50 = F_index_open(m, v48, int32(3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v52 = F_BuildIndexInfo(m, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+v25))) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45+v27))) = v52
	v68 = v36 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v68 < v69 {
		v36 = v68
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+116)))
	if v56 != int32(1) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+192))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+15)))
	if v60 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_BuildSpeculativeIndexInfo(m, v50, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L13
L22:
	;
	goto L1
}
func F_ExecPendingInserts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v10 = int32(0)
	goto L1
L1:
	;
	v15 = int32(0)
	if v8 == v15 {
		v25 = v15
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v7 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v19 <= v10 {
		v25 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = v21 + v10<<(uint(int32(2))%32)
	goto L3
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+108))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+108)))
	F_ExecBatchInsert(m, v45, v46, v47, v48, v49, l0, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L15
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	F_list_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v28 <= v10 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v25 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v35 = v32 + v10<<(uint(int32(2))%32)
	if v35 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	return
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_list_free(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	return
L15:
	;
	v10 = v10 + int32(1)
	goto L1
}
func F_ExecProcessReturning(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v17 = int32(2)
	if base.Ui32(v17) <= base.Ui32(l2-v17) {
		if l2 == int32(4) {
			if l3 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l3
				v54 = int32(0)
				v55 = l3
				*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
				if l4 != 0 {
					v67 = int32(0)
					v68 = l4
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
					v74 = v70&int32(231) | v54 | v67
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
					m.T0[v79].(func(*base.Module, int32))(m, v77)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = int32(4562080)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
							m.G0 = v12 + int32(16)
							return v77
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(16)
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
					if v60&int32(4) == v58 {
						v67 = v59
						v68 = v58
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						v74 = v70&int32(231) | v54 | v67
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
						m.T0[v79].(func(*base.Module, int32))(m, v77)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = int32(4562080)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
								m.G0 = v12 + int32(16)
								return v77
							}
						}
					} else {
						v65 = F_ExecGetAllNullSlot(m, v14, l1)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = v59
							v68 = v65
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							v74 = v70&int32(231) | v54 | v67
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
							m.T0[v79].(func(*base.Module, int32))(m, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(4562080)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
									m.G0 = v12 + int32(16)
									return v77
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l5
				v44 = int32(8)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
				if v45&int32(2) != 0 {
					v52 = F_ExecGetAllNullSlot(m, v14, l1)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = v44
						v55 = v52
						*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
						if l4 != 0 {
							v67 = int32(0)
							v68 = l4
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							v74 = v70&int32(231) | v54 | v67
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
							m.T0[v79].(func(*base.Module, int32))(m, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(4562080)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
									m.G0 = v12 + int32(16)
									return v77
								}
							}
						} else {
							v58 = int32(0)
							v59 = int32(16)
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							if v60&int32(4) == v58 {
								v67 = v59
								v68 = v58
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
								v74 = v70&int32(231) | v54 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
								m.T0[v79].(func(*base.Module, int32))(m, v77)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = int32(4562080)
									v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(65533)
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
										m.G0 = v12 + int32(16)
										return v77
									}
								}
							} else {
								v65 = F_ExecGetAllNullSlot(m, v14, l1)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = v59
									v68 = v65
									*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
									v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
									v74 = v70&int32(231) | v54 | v67
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
									m.T0[v79].(func(*base.Module, int32))(m, v77)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = int32(4562080)
										v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
										v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
											v97 = v95 & int32(65533)
											*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
											*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
											m.G0 = v12 + int32(16)
											return v77
										}
									}
								}
							}
						}
					}
				} else {
					v54 = v44
					v55 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
					if l4 != 0 {
						v67 = int32(0)
						v68 = l4
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						v74 = v70&int32(231) | v54 | v67
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
						m.T0[v79].(func(*base.Module, int32))(m, v77)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = int32(4562080)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
								m.G0 = v12 + int32(16)
								return v77
							}
						}
					} else {
						v58 = int32(0)
						v59 = int32(16)
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						if v60&int32(4) == v58 {
							v67 = v59
							v68 = v58
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							v74 = v70&int32(231) | v54 | v67
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
							m.T0[v79].(func(*base.Module, int32))(m, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(4562080)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
									m.G0 = v12 + int32(16)
									return v77
								}
							}
						} else {
							v65 = F_ExecGetAllNullSlot(m, v14, l1)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = v59
								v68 = v65
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
								v74 = v70&int32(231) | v54 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
								m.T0[v79].(func(*base.Module, int32))(m, v77)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = int32(4562080)
									v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(65533)
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
										m.G0 = v12 + int32(16)
										return v77
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
				F_errmsg_internal(m, int32(509328), v12)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(524080), int32(316), int32(351680))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if l4 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l4
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l5
		if l3 == int32(0) {
			v44 = int32(8)
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
			if v45&int32(2) != 0 {
				v52 = F_ExecGetAllNullSlot(m, v14, l1)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = v44
					v55 = v52
					*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
					if l4 != 0 {
						v67 = int32(0)
						v68 = l4
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						v74 = v70&int32(231) | v54 | v67
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
						m.T0[v79].(func(*base.Module, int32))(m, v77)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = int32(4562080)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
								m.G0 = v12 + int32(16)
								return v77
							}
						}
					} else {
						v58 = int32(0)
						v59 = int32(16)
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						if v60&int32(4) == v58 {
							v67 = v59
							v68 = v58
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							v74 = v70&int32(231) | v54 | v67
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
							m.T0[v79].(func(*base.Module, int32))(m, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(4562080)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
									m.G0 = v12 + int32(16)
									return v77
								}
							}
						} else {
							v65 = F_ExecGetAllNullSlot(m, v14, l1)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = v59
								v68 = v65
								*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
								v74 = v70&int32(231) | v54 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
								m.T0[v79].(func(*base.Module, int32))(m, v77)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = int32(4562080)
									v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(65533)
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
										*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
										m.G0 = v12 + int32(16)
										return v77
									}
								}
							}
						}
					}
				}
			} else {
				v54 = v44
				v55 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
				if l4 != 0 {
					v67 = int32(0)
					v68 = l4
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
					v74 = v70&int32(231) | v54 | v67
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
					m.T0[v79].(func(*base.Module, int32))(m, v77)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = int32(4562080)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
							m.G0 = v12 + int32(16)
							return v77
						}
					}
				} else {
					v58 = int32(0)
					v59 = int32(16)
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
					if v60&int32(4) == v58 {
						v67 = v59
						v68 = v58
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						v74 = v70&int32(231) | v54 | v67
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
						m.T0[v79].(func(*base.Module, int32))(m, v77)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = int32(4562080)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
								m.G0 = v12 + int32(16)
								return v77
							}
						}
					} else {
						v65 = F_ExecGetAllNullSlot(m, v14, l1)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = v59
							v68 = v65
							*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
							v74 = v70&int32(231) | v54 | v67
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
							m.T0[v79].(func(*base.Module, int32))(m, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = int32(4562080)
								v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(65533)
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
									m.G0 = v12 + int32(16)
									return v77
								}
							}
						}
					}
				}
			}
		} else {
			v54 = int32(0)
			v55 = l3
			*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v55
			if l4 != 0 {
				v67 = int32(0)
				v68 = l4
				*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
				v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
				v74 = v70&int32(231) | v54 | v67
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
				m.T0[v79].(func(*base.Module, int32))(m, v77)
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					v82 = int32(4562080)
					v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
					v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
						v97 = v95 & int32(65533)
						*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
						*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
						m.G0 = v12 + int32(16)
						return v77
					}
				}
			} else {
				v58 = int32(0)
				v59 = int32(16)
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
				if v60&int32(4) == v58 {
					v67 = v59
					v68 = v58
					*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
					v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
					v74 = v70&int32(231) | v54 | v67
					*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
					m.T0[v79].(func(*base.Module, int32))(m, v77)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = int32(4562080)
						v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(65533)
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
							v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
							*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
							m.G0 = v12 + int32(16)
							return v77
						}
					}
				} else {
					v65 = F_ExecGetAllNullSlot(m, v14, l1)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = v59
						v68 = v65
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
						v74 = v70&int32(231) | v54 | v67
						*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v74)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
						m.T0[v79].(func(*base.Module, int32))(m, v77)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = int32(4562080)
							v83 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(65533)
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v97)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
								*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v100)
								m.G0 = v12 + int32(16)
								return v77
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecSecLabelStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v15 == v3 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L30
	} else {
		goto L156
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L30
	} else {
		goto L152
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L30
	} else {
		goto L148
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L30
	} else {
		goto L144
	}
L5:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v86 - int32(1) {
	case 0, 5, 8, 11, 13, 17, 18, 20, 21, 22, 28, 29, 32, 33, 35, 36, 37, 40, 41, 48, 50:
		goto L28
	default:
		goto L29
	}
L6:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v14 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v20 != int32(1) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v80 = v24
	goto L5
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v27 <= int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v30 = int32(0)
	if v30 < v27 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = v27
	goto L15
L14:
	;
	v33 = v30
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v38 = v3
	goto L16
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34+v38<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v51 == int32(0) {
		v70 = v50
		v71 = v51
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L2
L18:
	;
	if v71-v70 == int32(0) {
		v80 = v46
		goto L5
	} else {
		goto L26
	}
L19:
	;
	goto L18
L20:
	;
	if v50 != v51 {
		v70 = v50
		v71 = v51
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v55 = v15
	v56 = v47
	goto L22
L22:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v59
		v71 = v60
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v70 = v59
	v71 = v60
	goto L19
L24:
	;
	v63 = int32(1)
	if v59 == v60 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v76 = v38 + int32(1)
	if v33 != v76 {
		v38 = v76
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_get_object_address(m, l0, v86, v105, v11+int32(44), int32(4), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L30
	} else {
		goto L35
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return
L31:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(118723), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(522450), int32(160), int32(104266))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v118
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_check_object_ownership(m, v113, v115, v11+int32(16), v114, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v125 == int32(6) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+119)))
	v132 = v130 - int32(99)
	if base.Ui32(int32(19)) < base.Ui32(v132) {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	m.T0[v143].(func(*base.Module, int32, int32))(m, l0, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L30
	} else {
		goto L42
	}
L40:
	;
	if int32(1)<<(uint(v132)%32)&int32(566281) == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v148 = int32(0)
	v149 = m.G0
	v151 = v149 - int32(240)
	m.G0 = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v156 = int32(1)
	if v153 <= int32(3591) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	m.G0 = v151 + int32(240)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v440 != 0 {
		goto L140
	} else {
		goto L141
	}
L44:
	;
	if v224 != 0 {
		goto L72
	} else {
		goto L73
	}
L45:
	;
	goto L44
L46:
	;
	v224 = int32(0)
	goto L45
L47:
	;
	if base.Ui32(v153-int32(2964)) < base.Ui32(int32(4)) {
		v224 = v156
		goto L45
	} else {
		goto L70
	}
L48:
	;
	if v153 <= int32(2670) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v153 <= int32(5999) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	switch v153 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v224 = v156
		goto L45
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L46
	default:
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v168 = v153 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v168) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v153-int32(2396)) {
		goto L46
	} else {
		goto L55
	}
L55:
	;
	v224 = v156
	goto L45
L56:
	;
	if int32(1)<<(uint(v168)%32)&int32(226492515) == int32(0) {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v224 = v156
	goto L45
L58:
	;
	if base.Ui32(v153-int32(3592)) < base.Ui32(int32(2)) {
		v224 = v156
		goto L45
	} else {
		goto L68
	}
L59:
	;
	v180 = v153 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v180) {
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	switch v153 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v224 = v156
		goto L45
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L46
	default:
		goto L64
	}
L62:
	;
	if int32(1)<<(uint(v180)%32)&int32(963) == int32(0) {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v224 = v156
	goto L45
L64:
	;
	if base.Ui32(v153-int32(6000)) < base.Ui32(int32(3)) {
		v224 = v156
		goto L45
	} else {
		goto L65
	}
L65:
	;
	v196 = v153 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v196) {
		goto L46
	} else {
		goto L66
	}
L66:
	;
	if int32(1)<<(uint(v196)%32)&int32(49153) != 0 {
		v224 = v156
		goto L45
	} else {
		goto L67
	}
L67:
	;
	goto L46
L68:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v153-int32(4060)) {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	v224 = v156
	goto L45
L70:
	;
	if base.Ui32(v153-int32(2846)) < base.Ui32(int32(2)) {
		v224 = v156
		goto L45
	} else {
		goto L71
	}
L71:
	;
	goto L46
L72:
	;
	v225 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v225
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+16)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v231
	v233 = F_cstring_to_text(m, v146)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L30
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+12)) = uint8(v322)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)) = uint8(v322)
	*(*int32)(unsafe.Add(mBase, uint32(v151)+8)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = v322
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+16)) = v330
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+20)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+24)) = v334
	v336 = F_cstring_to_text(m, v146)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L30
	} else {
		goto L107
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+24)) = v233
	if v147 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v236 = F_cstring_to_text(m, v147)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L30
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v151+int32(48), int32(1), int32(3), int32(184), v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L30
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v236
	goto L78
L80:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v151+int32(96), int32(2), int32(3), int32(184), v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L30
	} else {
		goto L81
	}
L81:
	;
	v257 = int32(3)
	v260 = F_cstring_to_text(m, v146)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L30
	} else {
		goto L82
	}
L82:
	;
	F_ScanKeyInit(m, v151+int32(144), v257, v257, int32(67), v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L30
	} else {
		goto L83
	}
L83:
	;
	v266 = F_table_open(m, int32(3592), int32(3))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L30
	} else {
		goto L85
	}
L84:
	;
	F_sequence_close(m, v266, int32(3))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L30
	} else {
		goto L106
	}
L85:
	;
	v274 = F_systable_beginscan(m, v266, int32(3593), int32(1), int32(0), int32(3), v151+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L30
	} else {
		goto L86
	}
L86:
	;
	v276 = F_systable_getnext(m, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L30
	} else {
		goto L87
	}
L87:
	;
	if v276 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v147 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v299 = v148
	goto L90
L90:
	;
	F_systable_endscan(m, v274)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L30
	} else {
		goto L98
	}
L91:
	;
	F_CatalogTupleDelete(m, v266, v276+int32(4))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L30
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)) = uint8(v286)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v266)+52))
	v295 = F_heap_modify_tuple(m, v276, v290, v151+int32(16), v151+int32(8), v151)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L30
	} else {
		goto L96
	}
L94:
	;
	F_systable_endscan(m, v274)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L30
	} else {
		goto L95
	}
L95:
	;
	goto L84
L96:
	;
	F_CatalogTupleUpdate(m, v266, v276+int32(4), v295)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L30
	} else {
		goto L97
	}
L97:
	;
	v299 = v295
	goto L90
L98:
	;
	if v147 == int32(0) {
		v313 = v299
		goto L99
	} else {
		goto L100
	}
L99:
	;
	if v313 == int32(0) {
		goto L84
	} else {
		goto L104
	}
L100:
	;
	if v299 != 0 {
		v313 = v299
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v266)+52))
	v309 = F_heap_form_tuple(m, v304, v151+int32(16), v151+int32(8))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L30
	} else {
		goto L102
	}
L102:
	;
	F_CatalogTupleInsert(m, v266, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L30
	} else {
		goto L103
	}
L103:
	;
	v313 = v309
	goto L99
L104:
	;
	F_pfree(m, v313)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L30
	} else {
		goto L105
	}
L105:
	;
	goto L84
L106:
	;
	goto L43
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+28)) = v336
	if v147 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v339 = F_cstring_to_text(m, v147)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L30
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v151+int32(48), int32(1), int32(3), int32(184), v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L30
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+32)) = v339
	goto L110
L112:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v151+int32(96), int32(2), int32(3), int32(184), v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L30
	} else {
		goto L113
	}
L113:
	;
	v360 = int32(3)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ScanKeyInit(m, v151+int32(144), v360, v360, int32(65), v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L30
	} else {
		goto L114
	}
L114:
	;
	v371 = F_cstring_to_text(m, v146)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L30
	} else {
		goto L115
	}
L115:
	;
	F_ScanKeyInit(m, v151+int32(192), int32(4), int32(3), int32(67), v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L30
	} else {
		goto L116
	}
L116:
	;
	v377 = F_table_open(m, int32(3596), int32(3))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L30
	} else {
		goto L118
	}
L117:
	;
	F_sequence_close(m, v377, int32(3))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L30
	} else {
		goto L139
	}
L118:
	;
	v385 = F_systable_beginscan(m, v377, int32(3597), int32(1), int32(0), int32(4), v151+int32(48))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L30
	} else {
		goto L119
	}
L119:
	;
	v387 = F_systable_getnext(m, v385)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L30
	} else {
		goto L120
	}
L120:
	;
	if v387 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v147 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	v410 = v148
	goto L123
L123:
	;
	F_systable_endscan(m, v385)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L30
	} else {
		goto L131
	}
L124:
	;
	F_CatalogTupleDelete(m, v377, v387+int32(4))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L30
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v397 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)) = uint8(v397)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v377)+52))
	v406 = F_heap_modify_tuple(m, v387, v401, v151+int32(16), v151+int32(8), v151)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L30
	} else {
		goto L129
	}
L127:
	;
	F_systable_endscan(m, v385)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L30
	} else {
		goto L128
	}
L128:
	;
	goto L117
L129:
	;
	F_CatalogTupleUpdate(m, v377, v387+int32(4), v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L30
	} else {
		goto L130
	}
L130:
	;
	v410 = v406
	goto L123
L131:
	;
	if v147 == int32(0) {
		v424 = v410
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v424 == int32(0) {
		goto L117
	} else {
		goto L137
	}
L133:
	;
	if v410 != 0 {
		v424 = v410
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v377)+52))
	v420 = F_heap_form_tuple(m, v415, v151+int32(16), v151+int32(8))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L30
	} else {
		goto L135
	}
L135:
	;
	F_CatalogTupleInsert(m, v377, v420)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L30
	} else {
		goto L136
	}
L136:
	;
	v424 = v420
	goto L132
L137:
	;
	F_pfree(m, v424)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L30
	} else {
		goto L138
	}
L138:
	;
	goto L117
L139:
	;
	goto L43
L140:
	;
	F_relation_close(m, v440, int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L30
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	m.G0 = v11 + int32(48)
	return
L143:
	;
	goto L142
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L30
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(484640), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L30
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(522450), int32(131), int32(104266))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L30
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L30
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(484685), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L30
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(522450), int32(135), int32(104266))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L30
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L30
	} else {
		goto L153
	}
L153:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v494
	F_errmsg(m, int32(484534), v11+int32(32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L30
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(522450), int32(154), int32(104266))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L30
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L30
	} else {
		goto L157
	}
L157:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v514 + int32(4)
	F_errmsg(m, int32(739573), v11)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L30
	} else {
		goto L158
	}
L158:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+48))
	v523 = int32(*(*int8)(unsafe.Add(mBase, uint32(v522)+119)))
	F_errdetail_relkind_not_supported(m, v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L30
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(522450), int32(195), int32(104266))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L30
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExpireTreeKnownAssignedTransactionIds(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v7 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v11 = F_LWLockAcquire(m, v7+int32(512), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		F_KnownAssignedXidsRemoveTree(m, l0, l1, l2)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[64]))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
			if v18 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v18)) == int32(0) {
					v30 = base.B2i32(base.Ui32(v18) < base.Ui32(l3))
				} else {
					v30 = int32(base.Ui32(v18-l3) >> (uint(int32(31)) % 32))
				}
				v32 = *(*int32)(unsafe.Add(mBase, _consts[64]))
				if v30 == int32(0) {
					v43 = v32
				} else {
					v36 = v32
					*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
					v43 = v36
				}
			} else {
				v36 = v16
				*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
				v43 = v36
			}
			v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v43)+56)) = v44 + int64(1)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v49+int32(512))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F___errno_location(m *base.Module) int32 {
	return int32(4735052)
}
func F___extenddftf2(m *base.Module, l0 int32, l1 float64) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v63 int64
	_ = v63
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	v3 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = base.I64_reinterpret_f64(l1)
	v15 = v13 & int64(4503599627370495)
	v19 = int64(base.Ui64(v13)>>(uint(int64(52))%64)) & int64(2047)
	if v19 != v3 {
		if v19 != int64(2047) {
			v83 = v19 + int64(15360)
			v84 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v86 = v15 << (uint(int64(60)) % 64)
		} else {
			v83 = int64(32767)
			v84 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v86 = v15 << (uint(int64(60)) % 64)
		}
	} else {
		if v15 == int64(0) {
			v37 = int64(0)
			v83 = v37
			v84 = v3
			v86 = v37
		} else {
			v39 = int64(0)
			if base.Ui64(v15) < base.Ui64(int64(4294967296)) {
				v50 = base.I32_clz(base.I32_wrap_i64(v13)) | int32(32)
			} else {
				v50 = base.I32_clz(base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(int64(32)) % 64))))
			}
			v52 = v50 + int32(49)
			if v52&int32(64) != 0 {
				v71 = int64(0)
				v72 = v15 << (uint(base.I64_extend_i32_u(v50+int32(-15))) % 64)
			} else {
				if v52 == int32(0) {
					v71 = v15
					v72 = v39
				} else {
					v63 = base.I64_extend_i32_u(v52)
					v71 = v15 << (uint(v63) % 64)
					v72 = v39<<(uint(v63)%64) | int64(base.Ui64(v15)>>(uint(base.I64_extend_i32_u(int32(64)-v52))%64))
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v71
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v72
			v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
			v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v83 = base.I64_extend_i32_u(int32(15372) - v50)
			v84 = v79 ^ int64(281474976710656)
			v86 = v82
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v86
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v13&int64(-9223372036854775807-1) | v83<<(uint(int64(48))%64) | v84
	m.G0 = v11 + int32(16)
	return
}
func F__equalA_Expr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != v5 {
		v25 = v3
		return v25
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v9 = F_equal(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v25 = v3
				return v25
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v17 = F_equal(m, v15, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						v25 = v3
						return v25
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						v23 = F_equal(m, v21, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							v25 = v23
							return v25
						}
					}
				}
			}
		}
	}
}
func F__equalA_Indices(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(0)
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v4 != v5 {
		v19 = v3
		return v19
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v9 = F_equal(m, v7, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				v19 = v3
				return v19
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v17 = F_equal(m, v15, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					return v19
				}
			}
		}
	}
}
func F__equalCoerceViaIO(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v18 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v18 = v3
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v18 = base.B2i32(v15 == v16)
			}
		}
		return v18
	}
}
func F__equalNullTest(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v18 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v18 = v3
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
				v18 = base.B2i32(v15 == v16)
			}
		}
		return v18
	}
}
func F__equalRelabelType(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v21 = v3
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v12 != v13 {
				v21 = v3
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v15 != v16 {
					v21 = v3
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					v21 = base.B2i32(v18 == v19)
				}
			}
		}
		return v21
	}
}
func F_each_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+32))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(118291), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518949), int32(2176), int32(88506))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_each_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 != 0 {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
		if v6 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
		} else {
		}
		return int32(0)
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(241532), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518949), int32(2190), int32(241405))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_ec_member_matches_indexcol(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v9 = v7 << (uint(int32(2)) % 32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
	if v14 != int32(403) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v13 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18+v9)))
	v21 = int32(0)
	if v17 == v21 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v59 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= int32(0) {
		v52 = v21
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v59 = v52
	goto L3
L8:
	;
	v30 = int32(0)
	if v30 < v27 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = v27
	goto L11
L10:
	;
	v33 = v30
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v36 = int32(0)
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34+v36<<(uint(int32(2))%32))))
	v45 = base.B2i32(v44 == v20)
	if v44 == v20 {
		v52 = v45
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v52 = v45
	goto L7
L14:
	;
	v47 = v36 + int32(1)
	if v47 != v33 {
		v36 = v47
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	return int32(0)
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v69 = F_match_index_to_operand(m, v68, v7, v10)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v13 == v64 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	return int32(0)
L20:
	;
	return int32(0)
L21:
	;
	return v69
}
func F_elements_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v3)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v12 == int32(1) {
		v15 = int32(4562080)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v18
		if l1 == int32(0) {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
			if v29 == int32(1) {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v33 = F_cstring_to_text(m, v32)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v33
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v38)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v53 = F_heap_form_tuple(m, v48, v7+int32(12), v7+int32(11))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_tuplestore_puttuple(m, v55, v53)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_MemoryContextReset(m, v60)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
				v44 = F_cstring_to_text_with_len(m, v40, v42-v40)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v44
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v53 = F_heap_form_tuple(m, v48, v7+int32(12), v7+int32(11))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_tuplestore_puttuple(m, v55, v53)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							F_MemoryContextReset(m, v60)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v22 != int32(1) {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
				if v29 == int32(1) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v33 = F_cstring_to_text(m, v32)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v33
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v38)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v53 = F_heap_form_tuple(m, v48, v7+int32(12), v7+int32(11))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_tuplestore_puttuple(m, v55, v53)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								F_MemoryContextReset(m, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(0)
								}
							}
						}
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
					v44 = F_cstring_to_text_with_len(m, v40, v42-v40)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v44
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v53 = F_heap_form_tuple(m, v48, v7+int32(12), v7+int32(11))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_tuplestore_puttuple(m, v55, v53)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								F_MemoryContextReset(m, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v27 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+11)) = uint8(v27)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v53 = F_heap_form_tuple(m, v48, v7+int32(12), v7+int32(11))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_tuplestore_puttuple(m, v55, v53)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						F_MemoryContextReset(m, v60)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return int32(0)
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
	}
}
func F_elements_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20
				F_errmsg(m, int32(25201), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518949), int32(2427), int32(88594))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return int32(0)
	}
}
func F_elements_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 != 0 {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
		if v10 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l1
		} else {
		}
		m.G0 = v6 + int32(16)
		return int32(0)
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v28
				F_errmsg(m, int32(241560), v6)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518949), int32(2442), int32(241337))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_eq_s(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6-v7 < l1 {
		v77 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v77
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = v10 + v7
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v73 != 0 {
		v77 = v4
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v73 = int32(0)
	goto L3
L5:
	;
	v47 = v42
	v48 = v43
	v49 = v44
	goto L15
L6:
	;
	if (v11|l2)&int32(3) != 0 {
		v42 = v11
		v43 = l2
		v44 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v35 = v11
	v36 = l2
	v37 = l1
	goto L8
L8:
	;
	if v37 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v19 = v11
	v20 = l2
	v21 = l1
	goto L10
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v24 != v25 {
		v42 = v19
		v43 = v20
		v44 = v21
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v35 = v30
	v36 = v28
	v37 = v32
	goto L8
L12:
	;
	v27 = int32(4)
	v28 = v20 + v27
	v30 = v19 + v27
	v32 = v21 - v27
	if base.Ui32(int32(3)) < base.Ui32(v32) {
		v19 = v30
		v20 = v28
		v21 = v32
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v42 = v35
	v43 = v36
	v44 = v37
	goto L5
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 == v53 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v73 = v52 - v53
	goto L3
L17:
	;
	v55 = int32(1)
	v60 = v49 - v55
	if v60 != 0 {
		v47 = v47 + v55
		v48 = v48 + v55
		v49 = v60
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1 + v7
	v77 = int32(1)
	goto L1
}
func F_eqjoinsel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 float64
	_ = v20
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 float32
	_ = v67
	var v69 float32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v115 int32
	_ = v115
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v131 float64
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 float32
	_ = v145
	var v147 float32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v174 float64
	_ = v174
	var v175 float64
	_ = v175
	var v179 int32
	_ = v179
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 float64
	_ = v190
	var v193 int32
	_ = v193
	var v200 float64
	_ = v200
	var v203 float64
	_ = v203
	var v204 int32
	_ = v204
	var v209 float64
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int64
	_ = v214
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 float32
	_ = v341
	var v342 float32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v389 float64
	_ = v389
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 float32
	_ = v484
	var v485 int32
	_ = v485
	var v487 float32
	_ = v487
	var v501 int32
	_ = v501
	var v512 float64
	_ = v512
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 float64
	_ = v531
	var v539 float64
	_ = v539
	var v540 float64
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v574 float64
	_ = v574
	var v575 float64
	_ = v575
	var v592 float32
	_ = v592
	var v593 float64
	_ = v593
	var v596 int32
	_ = v596
	var v597 float64
	_ = v597
	var v599 int32
	_ = v599
	var v603 float32
	_ = v603
	var v604 float64
	_ = v604
	var v607 int32
	_ = v607
	var v608 float64
	_ = v608
	var v610 float64
	_ = v610
	var v612 float64
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v637 float64
	_ = v637
	var v638 float64
	_ = v638
	var v655 float32
	_ = v655
	var v656 float64
	_ = v656
	var v659 int32
	_ = v659
	var v660 float64
	_ = v660
	var v662 float64
	_ = v662
	var v664 float64
	_ = v664
	var v665 float64
	_ = v665
	var v667 float64
	_ = v667
	var v675 float64
	_ = v675
	var v678 float64
	_ = v678
	var v704 float64
	_ = v704
	var v705 float64
	_ = v705
	var v708 float64
	_ = v708
	var v709 float64
	_ = v709
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 float64
	_ = v726
	var v731 float64
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v754 float64
	_ = v754
	var v755 float64
	_ = v755
	var v772 float32
	_ = v772
	var v773 float64
	_ = v773
	var v776 int32
	_ = v776
	var v777 float64
	_ = v777
	var v779 int32
	_ = v779
	var v783 float32
	_ = v783
	var v784 float64
	_ = v784
	var v787 int32
	_ = v787
	var v788 float64
	_ = v788
	var v790 float64
	_ = v790
	var v792 float64
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v817 float64
	_ = v817
	var v818 float64
	_ = v818
	var v835 float32
	_ = v835
	var v836 float64
	_ = v836
	var v839 int32
	_ = v839
	var v840 float64
	_ = v840
	var v842 float64
	_ = v842
	var v844 float64
	_ = v844
	var v845 float64
	_ = v845
	var v847 float64
	_ = v847
	var v856 float64
	_ = v856
	var v859 float64
	_ = v859
	var v883 float64
	_ = v883
	var v891 float64
	_ = v891
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 float64
	_ = v900
	var v903 float64
	_ = v903
	var v904 float64
	_ = v904
	var v909 float64
	_ = v909
	var v917 float64
	_ = v917
	var v925 float64
	_ = v925
	var v926 int32
	_ = v926
	var v927 float64
	_ = v927
	var v933 float64
	_ = v933
	var v940 float64
	_ = v940
	var v941 int32
	_ = v941
	var v942 float64
	_ = v942
	var v948 float64
	_ = v948
	var v949 int32
	_ = v949
	var v957 float64
	_ = v957
	var v959 float64
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 float32
	_ = v998
	var v1000 float64
	_ = v1000
	var v1003 float32
	_ = v1003
	var v1005 float64
	_ = v1005
	var v1009 float64
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1045 float64
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1079 int32
	_ = v1079
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1119 int32
	_ = v1119
	var v1127 float64
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1145 float64
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 float64
	_ = v1148
	var v1149 float64
	_ = v1149
	var v1150 float64
	_ = v1150
	var v1152 float64
	_ = v1152
	var v1155 float64
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 float64
	_ = v1173
	var v1181 float64
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	v2 = int32(0)
	v20 = float64(0)
	v35 = m.G0
	v37 = v35 - int32(256)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_get_join_variables(m, v41, v42, v43, v37+int32(160), v37+int32(128), v37+int32(47))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v55 = v37 + int32(160)
	v57 = v37 + int32(127)
	v58 = int32(0)
	v59 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v63 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v133 = v37 + int32(128)
	v135 = v37 + int32(126)
	v136 = int32(0)
	v137 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v136)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	if v141 != 0 {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+28)))
	if v101 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)))
	v66 = v64 + v65
	v67 = *(*float32)(unsafe.Add(mBase, uint32(v66)+8))
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v66)+16))
	v96 = base.F64_promote_f32(v69)
	v97 = base.F64_promote_f32(v67)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if v71 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v96 = float64(2)
	v97 = v59
	goto L4
L9:
	;
	goto L10
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v75 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v82 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+76))
	if v78 != int32(5) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v96 = float64(-1)
	v97 = v59
	goto L4
L14:
	;
	v96 = float64(0)
	v97 = v59
	goto L4
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v85 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)))
	switch v89 - int32(65530) {
	case 0:
		goto L17
	default:
		goto L14
	case 5:
		v96 = float64(-1)
		v97 = v59
		goto L4
	}
L17:
	;
	v96 = float64(1)
	v97 = v59
	goto L4
L18:
	;
	v102 = base.F64_neg(base.F64_sub(float64(1), v97))
	goto L20
L19:
	;
	v102 = v96
	goto L20
L20:
	;
	if base.F64_gt(v102, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v105 = F_clamp_row_est(m, v102)
	mBase = m.M
	v131 = v105
	goto L3
L22:
	;
	goto L23
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v106 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v109)
	v131 = float64(200)
	goto L3
L25:
	;
	goto L26
L26:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v106)+120))
	if base.F64_le(v112, float64(0)) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v115)
	v131 = float64(200)
	goto L3
L28:
	;
	goto L29
L29:
	;
	if base.F64_lt(v102, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = F_clamp_row_est(m, base.F64_mul(v112, base.F64_neg(v102)))
	mBase = m.M
	v131 = v122
	goto L3
L31:
	;
	goto L32
L32:
	;
	if base.F64_lt(v112, float64(200)) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v125 = F_clamp_row_est(m, v112)
	mBase = m.M
	v131 = v125
	goto L3
L34:
	;
	goto L35
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v126)
	v131 = float64(200)
	goto L3
L36:
	;
	v210 = F_get_opcode(m, v39)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L69
	}
L37:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+28)))
	if v179 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
	v144 = v142 + v143
	v145 = *(*float32)(unsafe.Add(mBase, uint32(v144)+8))
	v147 = *(*float32)(unsafe.Add(mBase, uint32(v144)+16))
	v174 = base.F64_promote_f32(v147)
	v175 = base.F64_promote_f32(v145)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	if v149 == int32(16) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = float64(2)
	v175 = v137
	goto L37
L42:
	;
	goto L43
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v153 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v160 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+76))
	if v156 != int32(5) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v174 = float64(-1)
	v175 = v137
	goto L37
L47:
	;
	v174 = float64(0)
	v175 = v137
	goto L37
L48:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v163 != int32(6) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+8)))
	switch v167 - int32(65530) {
	case 0:
		goto L50
	default:
		goto L47
	case 5:
		v174 = float64(-1)
		v175 = v137
		goto L37
	}
L50:
	;
	v174 = float64(1)
	v175 = v137
	goto L37
L51:
	;
	v180 = base.F64_neg(base.F64_sub(float64(1), v175))
	goto L53
L52:
	;
	v180 = v174
	goto L53
L53:
	;
	if base.F64_gt(v180, float64(0)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v183 = F_clamp_row_est(m, v180)
	mBase = m.M
	v209 = v183
	goto L36
L55:
	;
	goto L56
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v184 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v187)
	v209 = float64(200)
	goto L36
L58:
	;
	goto L59
L59:
	;
	v190 = *(*float64)(unsafe.Add(mBase, uint32(v184)+120))
	if base.F64_le(v190, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v193)
	v209 = float64(200)
	goto L36
L61:
	;
	goto L62
L62:
	;
	if base.F64_lt(v180, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v200 = F_clamp_row_est(m, base.F64_mul(v190, base.F64_neg(v180)))
	mBase = m.M
	v209 = v200
	goto L36
L64:
	;
	goto L65
L65:
	;
	if base.F64_lt(v190, float64(200)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v203 = F_clamp_row_est(m, v190)
	mBase = m.M
	v209 = v203
	goto L36
L67:
	;
	goto L68
L68:
	;
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v204)
	v209 = float64(200)
	goto L36
L69:
	;
	v212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+120)) = v212
	v214 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+112)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37)+104)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37)+96)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37)+56)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37-int32(-64)))) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37)+72)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v37)+80)) = v212
	*(*int64)(unsafe.Add(mBase, uint32(v37)+88)) = v214
	*(*int64)(unsafe.Add(mBase, uint32(v37)+48)) = v214
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if v235 == v212 {
		v258 = v212
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if v259 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v238 == int32(0) {
		v258 = v212
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v244 = int32(0)
	v246 = F_get_attstatsslot(m, v37+int32(88), v235, int32(1), v244, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v246 == int32(0) {
		v258 = v212
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	v254 = int32(0)
	v256 = F_get_attstatsslot(m, v37+int32(48), v252, int32(1), v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v258 = v256
	goto L70
L76:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	switch v1046 {
	case 0, 1, 2:
		v1155 = v1045
		goto L250
	default:
		goto L248
	case 4, 5:
		goto L251
	}
L77:
	;
	if v993 != 0 {
		goto L239
	} else {
		goto L240
	}
L78:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+126)))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+127)))
	v988 = int32(0)
	v989 = v980
	v992 = v983
	v993 = v984
	v994 = v986
	v996 = v987
	goto L77
L79:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v306)+16))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+22)))
	v317 = v315 + v316
	if v258 == int32(0) {
		v980 = v304
		v983 = v317
		v984 = v305
		goto L78
	} else {
		goto L105
	}
L80:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v307 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L81:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v306 != 0 {
		goto L79
	} else {
		goto L101
	}
L82:
	;
	v304 = v2
	v305 = v2
	goto L81
L83:
	;
	goto L84
L84:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+22)))
	v264 = v262 + v263
	if v258 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+189)))
	if v269 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v284 = int32(0)
	v287 = F_errstart(m, int32(13), v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L96
	}
L87:
	;
	v278 = v259
	goto L89
L88:
	;
	if v210 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v282 = F_get_attstatsslot(m, v37+int32(88), v278, int32(1), int32(0), int32(3))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L95
	}
L90:
	;
	v304 = int32(0)
	v305 = v264
	goto L81
L91:
	;
	goto L92
L92:
	;
	v273 = F_get_func_leakproof(m, v210)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v273 == int32(0) {
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	v278 = v277
	goto L89
L95:
	;
	v304 = v282
	v305 = v264
	goto L81
L96:
	;
	if v287 == int32(0) {
		v304 = v284
		v305 = v264
		goto L81
	} else {
		goto L97
	}
L97:
	;
	v291 = F_get_func_name(m, v210)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v291
	F_errmsg_internal(m, int32(355508), v37+int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(519114), int32(6242), int32(333732))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v304 = v284
	v305 = v264
	goto L81
L101:
	;
	v980 = v304
	v983 = v2
	v984 = v305
	goto L78
L102:
	;
	v980 = int32(0)
	v983 = v2
	v984 = v264
	goto L78
L103:
	;
	goto L104
L104:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+22)))
	v980 = int32(0)
	v983 = v311 + v312
	v984 = v264
	goto L78
L105:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+157)))
	if v322 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v962 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L234
	}
L107:
	;
	v330 = v306
	goto L109
L108:
	;
	if v210 == int32(0) {
		v980 = v304
		v983 = v317
		v984 = v305
		goto L78
	} else {
		goto L110
	}
L109:
	;
	v334 = F_get_attstatsslot(m, v37+int32(48), v330, int32(1), int32(0), int32(3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L113
	}
L110:
	;
	v325 = F_get_func_leakproof(m, v210)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	if v325 == int32(0) {
		goto L106
	} else {
		goto L112
	}
L112:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	v330 = v329
	goto L109
L113:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+126)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+127)))
	if v334&v304 != int32(1) {
		v988 = v334
		v989 = v304
		v992 = v317
		v993 = v305
		v994 = v336
		v996 = v337
		goto L77
	} else {
		goto L114
	}
L114:
	;
	v341 = *(*float32)(unsafe.Add(mBase, uint32(v305)+8))
	v342 = *(*float32)(unsafe.Add(mBase, uint32(v317)+8))
	F_fmgr_info(m, v210, v37+int32(192))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v347 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+252)) = uint8(v347)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+244)) = uint8(v347)
	v351 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+238)) = uint16(v351)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+236)) = uint8(v347)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+232)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v37)+224)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+220)) = v37 + int32(192)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v37)+104))
	v362 = F_palloc0(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	v365 = F_palloc0(m, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v37)+104))
	if v367 <= int32(0) {
		v704 = v20
		v705 = v20
		v708 = v20
		v709 = v20
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	if v718 <= int32(0) {
		v883 = float64(0)
		v891 = v20
		goto L176
	} else {
		goto L177
	}
L119:
	;
	v374 = v2
	v378 = v2
	v389 = v20
	goto L120
L120:
	;
	v405 = v374 << (uint(int32(2)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v37)+100))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405+v406)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+240)) = v408
	v410 = int32(0)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	if v411 <= v410 {
		v501 = v378
		v512 = v389
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v531 = float64(0)
	if base.F64_lt(v512, v531) != 0 {
		v539 = v531
		goto L137
	} else {
		goto L138
	}
L122:
	;
	v528 = v374 + int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v37)+104))
	if v528 < v529 {
		v374 = v528
		v378 = v501
		v389 = v512
		goto L120
	} else {
		goto L136
	}
L123:
	;
	v414 = v410
	v416 = v411
	goto L124
L124:
	;
	v448 = v414 + v365
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	if v449 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v477 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v448))) = uint8(v477)
	*(*uint8)(unsafe.Add(mBase, uint32(v374+v362))) = uint8(v477)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v37)+108))
	v484 = *(*float32)(unsafe.Add(mBase, uint32(v482+v405)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	v487 = *(*float32)(unsafe.Add(mBase, uint32(v485+v453)))
	v501 = v378 + v477
	v512 = base.F64_add(v389, base.F64_promote_f32(base.F32_mul(v484, v487)))
	goto L122
L126:
	;
	goto L125
L127:
	;
	v453 = v414 << (uint(int32(2)) % 32)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v37)+60))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453+v454)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+248)) = v456
	v458 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+236)) = uint8(v458)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v37)+220))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v464 = m.T0[v463].(func(*base.Module, int32) int32)(m, v37+int32(220))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	v472 = v416
	goto L129
L129:
	;
	v475 = v414 + int32(1)
	if v475 < v472 {
		v414 = v475
		v416 = v472
		goto L124
	} else {
		goto L135
	}
L130:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+236)))
	if v464 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v468 = v466
	goto L133
L132:
	;
	v468 = int32(1)
	goto L133
L133:
	;
	if v468 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L134:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	v472 = v471
	goto L129
L135:
	;
	v501 = v378
	v512 = v389
	goto L122
L136:
	;
	goto L121
L137:
	;
	v540 = base.F64_convert_i32_s(v501)
	if v529 <= int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	if base.F64_gt(v512, float64(1)) == int32(0) {
		v539 = v512
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v539 = float64(1)
	goto L137
L140:
	;
	v704 = v20
	v705 = v539
	v708 = v20
	v709 = v540
	goto L118
L141:
	;
	goto L142
L142:
	;
	v543 = int32(1)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v37)+108))
	if v529 == v543 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v529&v543 != 0 {
		goto L162
	} else {
		goto L163
	}
L144:
	;
	v618 = int32(0)
	v637 = float64(0)
	v638 = v20
	goto L143
L145:
	;
	goto L146
L146:
	;
	v552 = int32(0)
	v555 = v552
	v560 = v552
	v574 = float64(0)
	v575 = v20
	goto L147
L147:
	;
	v592 = *(*float32)(unsafe.Add(mBase, uint32(v545+v555<<(uint(int32(2))%32))))
	v593 = base.F64_promote_f32(v592)
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v362))))
	if v596 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v618 = v614
	v637 = v608
	v638 = v612
	goto L143
L149:
	;
	v597 = base.F64_add(v574, v593)
	goto L151
L150:
	;
	v597 = v574
	goto L151
L151:
	;
	v599 = v555 | int32(1)
	v603 = *(*float32)(unsafe.Add(mBase, uint32(v545+v599<<(uint(int32(2))%32))))
	v604 = base.F64_promote_f32(v603)
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v362))))
	if v607 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v608 = base.F64_add(v597, v604)
	goto L154
L153:
	;
	v608 = v597
	goto L154
L154:
	;
	if v596 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v610 = v575
	goto L157
L156:
	;
	v610 = base.F64_add(v575, v593)
	goto L157
L157:
	;
	if v607 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v612 = v610
	goto L160
L159:
	;
	v612 = base.F64_add(v610, v604)
	goto L160
L160:
	;
	v613 = int32(2)
	v614 = v555 + v613
	v616 = v560 + v613
	if v616 != v529&int32(2147483646) {
		v555 = v614
		v560 = v616
		v574 = v608
		v575 = v612
		goto L147
	} else {
		goto L161
	}
L161:
	;
	goto L148
L162:
	;
	v655 = *(*float32)(unsafe.Add(mBase, uint32(v545+v618<<(uint(int32(2))%32))))
	v656 = base.F64_promote_f32(v655)
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v362))))
	if v659 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v664 = v637
	v665 = v638
	goto L164
L164:
	;
	v667 = float64(0)
	if base.F64_lt(v664, v667) != 0 {
		v675 = v667
		goto L171
	} else {
		goto L172
	}
L165:
	;
	v660 = base.F64_add(v637, v656)
	goto L167
L166:
	;
	v660 = v637
	goto L167
L167:
	;
	if v659 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v662 = v638
	goto L170
L169:
	;
	v662 = base.F64_add(v638, v656)
	goto L170
L170:
	;
	v664 = v660
	v665 = v662
	goto L164
L171:
	;
	if base.F64_lt(v665, float64(0)) != 0 {
		v704 = v675
		v705 = v539
		v708 = v20
		v709 = v540
		goto L118
	} else {
		goto L174
	}
L172:
	;
	if base.F64_gt(v664, float64(1)) == int32(0) {
		v675 = v664
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v675 = float64(1)
	goto L171
L174:
	;
	v678 = float64(1)
	if base.F64_gt(v665, v678) != 0 {
		v704 = v675
		v705 = v539
		v708 = v678
		v709 = v540
		goto L118
	} else {
		goto L175
	}
L175:
	;
	v704 = v675
	v705 = v539
	v708 = v665
	v709 = v540
	goto L118
L176:
	;
	F_pfree(m, v362)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L211
	}
L177:
	;
	v721 = int32(1)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v37)+68))
	if v718 == v721 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v718&v721 != 0 {
		goto L197
	} else {
		goto L198
	}
L179:
	;
	v726 = float64(0)
	v798 = int32(0)
	v817 = v726
	v818 = v726
	goto L178
L180:
	;
	goto L181
L181:
	;
	v731 = float64(0)
	v732 = int32(0)
	v735 = v732
	v740 = v732
	v754 = v731
	v755 = v731
	goto L182
L182:
	;
	v772 = *(*float32)(unsafe.Add(mBase, uint32(v723+v735<<(uint(int32(2))%32))))
	v773 = base.F64_promote_f32(v772)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735+v365))))
	if v776 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v798 = v794
	v817 = v792
	v818 = v788
	goto L178
L184:
	;
	v777 = v755
	goto L186
L185:
	;
	v777 = base.F64_add(v755, v773)
	goto L186
L186:
	;
	v779 = v735 | int32(1)
	v783 = *(*float32)(unsafe.Add(mBase, uint32(v723+v779<<(uint(int32(2))%32))))
	v784 = base.F64_promote_f32(v783)
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365+v779))))
	if v787 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v788 = v777
	goto L189
L188:
	;
	v788 = base.F64_add(v777, v784)
	goto L189
L189:
	;
	if v776 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v790 = base.F64_add(v754, v773)
	goto L192
L191:
	;
	v790 = v754
	goto L192
L192:
	;
	if v787 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v792 = base.F64_add(v790, v784)
	goto L195
L194:
	;
	v792 = v790
	goto L195
L195:
	;
	v793 = int32(2)
	v794 = v735 + v793
	v796 = v740 + v793
	if v796 != v718&int32(2147483646) {
		v735 = v794
		v740 = v796
		v754 = v792
		v755 = v788
		goto L182
	} else {
		goto L196
	}
L196:
	;
	goto L183
L197:
	;
	v835 = *(*float32)(unsafe.Add(mBase, uint32(v723+v798<<(uint(int32(2))%32))))
	v836 = base.F64_promote_f32(v835)
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798+v365))))
	if v839 != 0 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	v844 = v817
	v845 = v818
	goto L199
L199:
	;
	v847 = float64(0)
	if base.F64_lt(v844, v847) != 0 {
		v856 = v847
		goto L206
	} else {
		goto L207
	}
L200:
	;
	v840 = v818
	goto L202
L201:
	;
	v840 = base.F64_add(v818, v836)
	goto L202
L202:
	;
	if v839 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v842 = base.F64_add(v817, v836)
	goto L205
L204:
	;
	v842 = v817
	goto L205
L205:
	;
	v844 = v842
	v845 = v840
	goto L199
L206:
	;
	if base.F64_lt(v845, float64(0)) != 0 {
		v883 = v847
		v891 = v856
		goto L176
	} else {
		goto L209
	}
L207:
	;
	if base.F64_gt(v844, float64(1)) == int32(0) {
		v856 = v844
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v856 = float64(1)
	goto L206
L209:
	;
	v859 = float64(1)
	if base.F64_gt(v845, v859) != 0 {
		v883 = v859
		v891 = v856
		goto L176
	} else {
		goto L210
	}
L210:
	;
	v883 = v845
	v891 = v856
	goto L176
L211:
	;
	F_pfree(m, v365)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v900 = float64(1)
	v903 = base.F64_sub(base.F64_sub(base.F64_sub(v900, base.F64_promote_f32(v342)), v891), v883)
	v904 = float64(0)
	v909 = base.F64_sub(base.F64_sub(base.F64_sub(v900, base.F64_promote_f32(v341)), v704), v708)
	if base.F64_lt(v909, v904) != 0 {
		v917 = v904
		goto L213
	} else {
		goto L214
	}
L213:
	;
	if base.F64_lt(v903, float64(0)) != 0 {
		v925 = v904
		goto L216
	} else {
		goto L217
	}
L214:
	;
	if base.F64_gt(v909, float64(1)) == int32(0) {
		v917 = v909
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v917 = float64(1)
	goto L213
L216:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v37)+64))
	v927 = base.F64_convert_i32_s(v926)
	if base.F64_lt(v927, v209) != 0 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	if base.F64_gt(v903, float64(1)) == int32(0) {
		v925 = v903
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v925 = float64(1)
	goto L216
L219:
	;
	v933 = base.F64_add(v705, base.F64_div(base.F64_mul(v708, v925), base.F64_sub(v209, v927)))
	goto L221
L220:
	;
	v933 = v705
	goto L221
L221:
	;
	if base.F64_gt(v209, v709) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v940 = base.F64_add(base.F64_div(base.F64_mul(v917, base.F64_add(v883, v925)), base.F64_sub(v209, v709)), v933)
	goto L224
L223:
	;
	v940 = v933
	goto L224
L224:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v37)+104))
	v942 = base.F64_convert_i32_s(v941)
	if base.F64_lt(v942, v131) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v948 = base.F64_add(v705, base.F64_div(base.F64_mul(v883, v917), base.F64_sub(v131, v942)))
	goto L227
L226:
	;
	v948 = v705
	goto L227
L227:
	;
	v949 = int32(1)
	if base.F64_gt(v131, v709) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v957 = base.F64_add(base.F64_div(base.F64_mul(base.F64_add(v708, v917), v925), base.F64_sub(v131, v709)), v948)
	goto L230
L229:
	;
	v957 = v948
	goto L230
L230:
	;
	if base.F64_lt(v940, v957) != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v959 = v940
	goto L233
L232:
	;
	v959 = v957
	goto L233
L233:
	;
	v1011 = v949
	v1013 = v949
	v1018 = v317
	v1020 = v305
	v1026 = v336
	v1027 = v337
	v1045 = v959
	goto L76
L234:
	;
	if v962 == int32(0) {
		v980 = v304
		v983 = v317
		v984 = v305
		goto L78
	} else {
		goto L235
	}
L235:
	;
	v966 = F_get_func_name(m, v210)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v966
	F_errmsg_internal(m, int32(355508), v37+int32(16))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(519114), int32(6242), int32(333732))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v980 = v304
	v983 = v317
	v984 = v305
	goto L78
L239:
	;
	v998 = *(*float32)(unsafe.Add(mBase, uint32(v993)+8))
	v1000 = base.F64_promote_f32(v998)
	goto L241
L240:
	;
	v1000 = v20
	goto L241
L241:
	;
	if v992 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1003 = *(*float32)(unsafe.Add(mBase, uint32(v992)+8))
	v1005 = base.F64_promote_f32(v1003)
	goto L244
L243:
	;
	v1005 = v20
	goto L244
L244:
	;
	if base.F64_gt(v131, v209) != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1009 = v131
	goto L247
L246:
	;
	v1009 = v209
	goto L247
L247:
	;
	v1011 = v988
	v1013 = v989
	v1018 = v992
	v1020 = v993
	v1026 = v994
	v1027 = v996
	v1045 = base.F64_div(base.F64_mul(base.F64_sub(float64(1), v1000), base.F64_sub(float64(1), v1005)), v1009)
	goto L76
L248:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L307
	}
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L304
	}
L250:
	;
	F_free_attstatsslot(m, v37+int32(88))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L290
	}
L251:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v1047 == int32(0) {
		goto L249
	} else {
		goto L252
	}
L252:
	;
	v1052 = int32(0)
	if v1047 == v1052 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	if v1111 == int32(0) {
		goto L249
	} else {
		goto L275
	}
L254:
	;
	if v1105 != 0 {
		goto L270
	} else {
		goto L271
	}
L255:
	;
	v1105 = int32(0)
	goto L254
L256:
	;
	goto L257
L257:
	;
	v1060 = int32(1)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	if v1061 <= v1060 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1064 = v1060
	goto L260
L259:
	;
	v1064 = v1061
	goto L260
L260:
	;
	v1069 = int32(0)
	v1072 = int32(-1)
	goto L262
L261:
	;
	v1105 = v1097
	goto L254
L262:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1047+int32(8)+v1069<<(uint(int32(2))%32))))
	if v1079 != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(220)))) = v1089
	v1097 = int32(1)
	goto L261
L264:
	;
	if int32(0) <= v1072 {
		v1097 = v1052
		goto L261
	} else {
		goto L267
	}
L265:
	;
	v1089 = v1072
	goto L266
L266:
	;
	v1091 = v1069 + int32(1)
	if v1091 != v1064 {
		v1069 = v1091
		v1072 = v1089
		goto L262
	} else {
		goto L269
	}
L267:
	;
	if base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v1079)) {
		v1097 = v1052
		goto L261
	} else {
		goto L268
	}
L268:
	;
	v1089 = base.I32_ctz(v1079) | v1069<<(uint(int32(5))%32)
	goto L266
L269:
	;
	goto L263
L270:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v37)+220))
	v1107 = F_find_base_rel(m, v41, v1106)
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1109 = F_find_join_rel(m, v41, v1047)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	v1111 = v1107
	goto L253
L274:
	;
	v1111 = v1109
	goto L253
L275:
	;
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+47)))
	if v1114 == int32(0) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	v1149 = *(*float64)(unsafe.Add(mBase, uint32(v1111)+16))
	v1150 = base.F64_mul(v1045, v1149)
	if base.F64_gt(v1150, v1148) != 0 {
		goto L287
	} else {
		goto L288
	}
L277:
	;
	v1119 = int32(1)
	v1127 = F_eqjoinsel_semi(m, v210, v40, v37+int32(128), v131, v209, v1027&v1119, v1026&v1119, v37+int32(88), v37+int32(48), v1020, v1013, v1011, v1111)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1129 = F_get_commutator(m, v39)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L1
	} else {
		goto L281
	}
L280:
	;
	v1148 = v1127
	goto L276
L281:
	;
	if v1129 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1131 = F_get_opcode(m, v1129)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	v1134 = int32(0)
	goto L284
L284:
	;
	v1137 = int32(1)
	v1145 = F_eqjoinsel_semi(m, v1134, v40, v37+int32(160), v209, v131, v1026&v1137, v1027&v1137, v37+int32(48), v37+int32(88), v1018, v1011, v1013, v1111)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L286
	}
L285:
	;
	v1134 = v1131
	goto L284
L286:
	;
	v1148 = v1145
	goto L276
L287:
	;
	v1152 = v1148
	goto L289
L288:
	;
	v1152 = v1150
	goto L289
L289:
	;
	v1155 = v1152
	goto L250
L290:
	;
	F_free_attstatsslot(m, v37+int32(48))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v37)+168))
	if v1165 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v37)+172))
	m.T0[v1166].(func(*base.Module, int32))(m, v1165)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v37)+136))
	if v1169 != 0 {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	goto L294
L296:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v37)+140))
	m.T0[v1170].(func(*base.Module, int32))(m, v1169)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1173 = float64(0)
	if base.F64_lt(v1155, v1173) != 0 {
		v1181 = v1173
		goto L300
	} else {
		goto L301
	}
L299:
	;
	goto L298
L300:
	;
	v1182 = F_Float8GetDatum(m, v1181)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L303
	}
L301:
	;
	if base.F64_gt(v1155, float64(1)) == int32(0) {
		v1181 = v1155
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v1181 = float64(1)
	goto L300
L303:
	;
	m.G0 = v37 + int32(256)
	return v1182
L304:
	;
	F_errmsg_internal(m, int32(183989), int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(519114), int32(6950), int32(322429))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v1206
	F_errmsg_internal(m, int32(508206), v37)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(519114), int32(2422), int32(321851))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errcode_for_file_access(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	if int32(0) <= v4 {
		v9 = v4 * int32(100)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[900])))
		switch v12 - int32(2) {
		case 0, 61, 67:
			v24 = int32(16797828)
		default:
			v24 = int32(2600)
		case 18:
			v24 = int32(33686021)
		case 27:
			v24 = int32(786949)
		case 29, 52, 53:
			v24 = int32(151027844)
		case 31, 39:
			v24 = int32(197)
		case 35:
			v24 = int32(50463237)
		case 42:
			v24 = int32(16908805)
		case 46:
			v24 = int32(8389)
		case 49:
			v24 = int32(4293)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_consts[901]))) = v24
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[899])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(475542), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errfinish(m, int32(523098), int32(882), int32(138649))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_errhint_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4555004)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4562080)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[903])))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[899])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v40 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v45 = v40
	goto L10
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[907])))
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v58 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v58 != 0 {
		v45 = v58
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v70 = F_pstrdup(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[907]))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v78 = int32(4555004)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v80 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(475542), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(523098), int32(1346), int32(325935))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errmsg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4555004)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[899]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4562080)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[903])))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[904]))) = l0
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[899])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v41 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v47 = v41
	goto L10
L8:
	;
	goto L9
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[905])))
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[900])))
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v59 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v59 != 0 {
		v47 = v59
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v71 = F_pstrdup(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[905]))) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v79 = int32(4555004)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[902]))
	*(*int32)(unsafe.Add(mBase, _consts[902])) = v81 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(475542), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(523098), int32(1077), int32(343749))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errorConflictingDefElem(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errmsg(m, int32(146234), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_parser_errposition(m, l1, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(523833), int32(376), int32(305936))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_error_severity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v5 = l0 - int32(10)
	if base.Ui32(v5) <= base.Ui32(int32(13)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_consts[922])))
		v13 = v12
	} else {
		v13 = int32(572618)
	}
	return v13
}
func F_esc_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v89 int64
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v13) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v15 = l0
	v17 = l2
	v22 = v8
	goto L5
L3:
	;
	v89 = v8
	goto L4
L4:
	;
	m.G0 = v11 + int32(16)
	return v89
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v23 != int32(92) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v89 = v78
	goto L4
L7:
	;
	v78 = v22 + int64(1)
	if base.Ui32(v76) < base.Ui32(v13) {
		v15 = v76
		v17 = v17 + int32(1)
		v22 = v78
		goto L5
	} else {
		goto L18
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v23)
	v76 = v15 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v30 = v15 + int32(3)
	if base.Ui32(v13) <= base.Ui32(v30) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = v15 + int32(1)
	if base.Ui32(v13) <= base.Ui32(v64) {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v32&int32(252) != int32(48) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
	if v37&int32(248) != int32(48) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v42&int32(248) != int32(48) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v56 = v42 + (v37<<(uint(int32(3))%32)&int32(56) | v32<<(uint(int32(6))%32)) - int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v56)
	v76 = v15 + int32(4)
	goto L7
L16:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v66 != int32(92) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v69)
	v76 = v15 + int32(2)
	goto L7
L18:
	;
	goto L6
L19:
	;
	return int64(0)
L20:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(533208)
	F_errmsg(m, int32(199922), v11)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(524498), int32(513), int32(433866))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_esc_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int64
	_ = v114
	v6 = int64(0)
	if base.Ui32(l0+l1) <= base.Ui32(l0) {
		v114 = v6
	} else {
		v9 = int32(3)
		v10 = l1 & v9
		if base.Ui32(l1-int32(1)) < base.Ui32(v9) {
			v77 = l0
			v82 = v6
		} else {
			v18 = l0
			v19 = int32(0)
			v23 = v6
			for {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v27 == int32(92) {
					v30 = int64(2)
				} else {
					v30 = int64(1)
				}
				if base.I32_extend8_s(v27) <= int32(0) {
					v34 = int64(4)
				} else {
					v34 = v30
				}
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
				if v39 == int32(92) {
					v42 = int64(2)
				} else {
					v42 = int64(1)
				}
				if base.I32_extend8_s(v39) <= int32(0) {
					v46 = int64(4)
				} else {
					v46 = v42
				}
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+2)))
				if v51 == int32(92) {
					v54 = int64(2)
				} else {
					v54 = int64(1)
				}
				if base.I32_extend8_s(v51) <= int32(0) {
					v58 = int64(4)
				} else {
					v58 = v54
				}
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+3)))
				if v63 == int32(92) {
					v66 = int64(2)
				} else {
					v66 = int64(1)
				}
				if base.I32_extend8_s(v63) <= int32(0) {
					v70 = int64(4)
				} else {
					v70 = v66
				}
				v71 = v23 + v34 + v46 + v58 + v70
				v72 = int32(4)
				v73 = v18 + v72
				v75 = v19 + v72
				if v75 != l1&int32(-4) {
					v18 = v73
					v19 = v75
					v23 = v71
					continue
				} else {
					break
				}
				break
			}
			v77 = v73
			v82 = v71
		}
		if v10 == int32(0) {
			v114 = v82
		} else {
			v86 = v77
			v87 = int32(0)
			v91 = v82
			for {
				v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
				if v95 == int32(92) {
					v98 = int64(2)
				} else {
					v98 = int64(1)
				}
				if base.I32_extend8_s(v95) <= int32(0) {
					v102 = int64(4)
				} else {
					v102 = v98
				}
				v103 = v91 + v102
				v104 = int32(1)
				v107 = v87 + v104
				if v107 != v10 {
					v86 = v86 + v104
					v87 = v107
					v91 = v103
					continue
				} else {
					break
				}
				break
			}
			v114 = v103
		}
	}
	return v114
}
func F_estimate_ln_dweight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v373 int64
	_ = v373
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v405 int64
	_ = v405
	var v420 int32
	_ = v420
	var v422 int64
	_ = v422
	var v432 int64
	_ = v432
	var v436 int64
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 float64
	_ = v445
	var v447 float64
	_ = v447
	var v460 float64
	_ = v460
	var v463 float64
	_ = v463
	var v468 float64
	_ = v468
	var v469 float64
	_ = v469
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v476 float64
	_ = v476
	var v477 float64
	_ = v477
	var v478 float64
	_ = v478
	var v503 float64
	_ = v503
	var v515 float64
	_ = v515
	var v537 float64
	_ = v537
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 float64
	_ = v569
	var v571 float64
	_ = v571
	var v582 int64
	_ = v582
	var v597 int32
	_ = v597
	var v599 int64
	_ = v599
	var v609 int64
	_ = v609
	var v613 int64
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v622 float64
	_ = v622
	var v624 float64
	_ = v624
	var v637 float64
	_ = v637
	var v640 float64
	_ = v640
	var v645 float64
	_ = v645
	var v646 float64
	_ = v646
	var v647 float64
	_ = v647
	var v648 float64
	_ = v648
	var v653 float64
	_ = v653
	var v654 float64
	_ = v654
	var v655 float64
	_ = v655
	var v680 float64
	_ = v680
	var v692 float64
	_ = v692
	var v714 float64
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 != 0 {
		v721 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v721
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(0) {
		v721 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = int32(1)
	v21 = int32(-1)
	v22 = int32(0)
	if base.B2i32(v21 < v18)&base.B2i32(v22 < v14) == v22 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	if v14 <= int32(0) {
		v721 = v2
		goto L1
	} else {
		goto L130
	}
L5:
	;
	if v193 < int32(0) {
		goto L4
	} else {
		goto L52
	}
L6:
	;
	v193 = v183
	goto L5
L7:
	;
	if v21 <= v53 {
		v88 = v21
		v90 = v22
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v53 = v18
	v57 = v22
	goto L7
L9:
	;
	goto L10
L10:
	;
	v34 = v18
	v38 = v22
	goto L11
L11:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v38<<(uint(int32(1))%32)))))
	if v44 != 0 {
		v183 = int32(1)
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v53 = v48
	v57 = v46
	goto L7
L13:
	;
	v45 = int32(1)
	v46 = v38 + v45
	v48 = v34 - v45
	if v48 <= v21 {
		v53 = v48
		v57 = v46
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if v46 < v14 {
		v34 = v48
		v38 = v46
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if v53 != v88 {
		v129 = v57
		v130 = v90
		goto L24
	} else {
		goto L25
	}
L17:
	;
	goto L18
L18:
	;
	v69 = v21
	v71 = v22
	goto L19
L19:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71<<(uint(int32(1))%32))+uint32(_consts[801]))))
	if v76 != 0 {
		v183 = int32(-1)
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v88 = v80
	v90 = v78
	goto L16
L21:
	;
	v77 = int32(1)
	v78 = v71 + v77
	v80 = v69 - v77
	if v80 <= v53 {
		v88 = v80
		v90 = v78
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v78 < v20 {
		v69 = v80
		v71 = v78
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	if v14 < v129 {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v99 = v57
	v100 = v90
	goto L26
L26:
	;
	if v14 <= v99 {
		v129 = v99
		v130 = v100
		goto L24
	} else {
		goto L28
	}
L27:
	;
	if base.I32_extend16_s(v115) < base.I32_extend16_s(v113) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	if v20 <= v100 {
		v129 = v99
		v130 = v100
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v104 = int32(1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v99<<(uint(v104)%32)))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100<<(uint(v104)%32))+uint32(_consts[801]))))
	if v113 == v115 {
		v99 = v99 + v104
		v100 = v100 + v104
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v122 = int32(1)
	goto L33
L32:
	;
	v122 = int32(-1)
	goto L33
L33:
	;
	v193 = v122
	goto L5
L34:
	;
	v133 = v129
	goto L36
L35:
	;
	v133 = v14
	goto L36
L36:
	;
	v140 = v129
	goto L37
L37:
	;
	if v133 == v140 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v183 = v166
	goto L6
L39:
	;
	if v20 < v130 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v166 = int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v140<<(uint(v166)%32)))))
	if v172 == int32(0) {
		v140 = v140 + v166
		goto L37
	} else {
		goto L51
	}
L42:
	;
	v145 = v130
	goto L44
L43:
	;
	v145 = v20
	goto L44
L44:
	;
	v153 = v130
	goto L45
L45:
	;
	if v145 == v153 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v183 = int32(-1)
	goto L6
L47:
	;
	v193 = int32(0)
	goto L5
L48:
	;
	goto L49
L49:
	;
	v157 = int32(1)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32))+uint32(_consts[801]))))
	if v162 == int32(0) {
		v153 = v153 + v157
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	goto L38
L52:
	;
	v197 = int32(2)
	v198 = int32(0)
	if base.B2i32(v198 < v18)&base.B2i32(v198 < v14) == v198 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	if int32(0) < v370 {
		goto L4
	} else {
		goto L100
	}
L54:
	;
	v370 = v360
	goto L53
L55:
	;
	if v198 <= v230 {
		v265 = v198
		v267 = v198
		goto L64
	} else {
		goto L65
	}
L56:
	;
	v230 = v18
	v234 = v198
	goto L55
L57:
	;
	goto L58
L58:
	;
	v211 = v18
	v215 = v198
	goto L59
L59:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v215<<(uint(int32(1))%32)))))
	if v221 != 0 {
		v360 = int32(1)
		goto L54
	} else {
		goto L61
	}
L60:
	;
	v230 = v225
	v234 = v223
	goto L55
L61:
	;
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v211 - v222
	if v225 <= v198 {
		v230 = v225
		v234 = v223
		goto L55
	} else {
		goto L62
	}
L62:
	;
	if v223 < v14 {
		v211 = v225
		v215 = v223
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	if v230 != v265 {
		v306 = v234
		v307 = v267
		goto L72
	} else {
		goto L73
	}
L65:
	;
	goto L66
L66:
	;
	v246 = v198
	v248 = v198
	goto L67
L67:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248<<(uint(int32(1))%32))+uint32(_consts[802]))))
	if v253 != 0 {
		v360 = int32(-1)
		goto L54
	} else {
		goto L69
	}
L68:
	;
	v265 = v257
	v267 = v255
	goto L64
L69:
	;
	v254 = int32(1)
	v255 = v248 + v254
	v257 = v246 - v254
	if v257 <= v230 {
		v265 = v257
		v267 = v255
		goto L64
	} else {
		goto L70
	}
L70:
	;
	if v255 < v197 {
		v246 = v257
		v248 = v255
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	if v14 < v306 {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	v276 = v234
	v277 = v267
	goto L74
L74:
	;
	if v14 <= v276 {
		v306 = v276
		v307 = v277
		goto L72
	} else {
		goto L76
	}
L75:
	;
	if base.I32_extend16_s(v292) < base.I32_extend16_s(v290) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v197 <= v277 {
		v306 = v276
		v307 = v277
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v281 = int32(1)
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v276<<(uint(v281)%32)))))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277<<(uint(v281)%32))+uint32(_consts[802]))))
	if v290 == v292 {
		v276 = v276 + v281
		v277 = v277 + v281
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	v299 = int32(1)
	goto L81
L80:
	;
	v299 = int32(-1)
	goto L81
L81:
	;
	v370 = v299
	goto L53
L82:
	;
	v310 = v306
	goto L84
L83:
	;
	v310 = v14
	goto L84
L84:
	;
	v317 = v306
	goto L85
L85:
	;
	if v310 == v317 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v360 = v343
	goto L54
L87:
	;
	if v197 < v307 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v343 = int32(1)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v317<<(uint(v343)%32)))))
	if v349 == int32(0) {
		v317 = v317 + v343
		goto L85
	} else {
		goto L99
	}
L90:
	;
	v322 = v307
	goto L92
L91:
	;
	v322 = v197
	goto L92
L92:
	;
	v330 = v307
	goto L93
L93:
	;
	if v322 == v330 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v360 = int32(-1)
	goto L54
L95:
	;
	v370 = int32(0)
	goto L53
L96:
	;
	goto L97
L97:
	;
	v334 = int32(1)
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v330<<(uint(v334)%32))+uint32(_consts[802]))))
	if v339 == int32(0) {
		v330 = v330 + v334
		goto L93
	} else {
		goto L98
	}
L98:
	;
	goto L94
L99:
	;
	goto L86
L100:
	;
	v373 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v373
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v373
	F_sub_var(m, l0, int32(1774696), v11+int32(8))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	return int32(0)
L102:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if int32(0) < v386 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v392))))
	v394 = base.F64_convert_i32_s(v393)
	v405 = base.I64_reinterpret_f64(v394)
	if v405 <= int64(4503599627370495) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	v545 = v2
	goto L105
L105:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	if v547 != 0 {
		goto L126
	} else {
		goto L127
	}
L106:
	;
	v545 = v389<<(uint(int32(2))%32) + v543
	goto L105
L107:
	;
	if base.F64_lt(base.F64_abs(v537), float64(2.147483648e+09)) != 0 {
		goto L123
	} else {
		goto L124
	}
L108:
	;
	v537 = v515
	goto L107
L109:
	;
	v441 = v439 + int32(614242)
	v445 = base.F64_convert_i32_s(int32(base.Ui32(v441)>>(uint(int32(20))%32)) + v438)
	v447 = base.F64_mul(v445, float64(0.30102999566361177))
	v460 = base.F64_add(base.F64_reinterpret_i64(v436&int64(4294967295)|base.I64_extend_i32_u(v441&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v463 = base.F64_mul(v460, base.F64_mul(v460, float64(0.5)))
	v468 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v460, v463)) & int64(-4294967296))
	v469 = float64(0.4342944818781689)
	v470 = base.F64_mul(v468, v469)
	v471 = base.F64_add(v447, v470)
	v476 = base.F64_div(v460, base.F64_add(v460, float64(2)))
	v477 = base.F64_mul(v476, v476)
	v478 = base.F64_mul(v477, v477)
	v503 = base.F64_add(base.F64_mul(v476, base.F64_add(v463, base.F64_add(base.F64_mul(v478, base.F64_add(base.F64_mul(v478, base.F64_add(base.F64_mul(v478, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v477, base.F64_add(base.F64_mul(v478, base.F64_add(base.F64_mul(v478, base.F64_add(base.F64_mul(v478, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v460, v468), v463))
	v515 = base.F64_add(v471, base.F64_add(base.F64_add(v470, base.F64_sub(v447, v471)), base.F64_add(base.F64_mul(v503, v469), base.F64_add(base.F64_mul(v445, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v503, v468), float64(2.5082946711645275e-11))))))
	goto L108
L110:
	;
	v432 = base.I64_reinterpret_f64(base.F64_mul(v394, float64(1.8014398509481984e+16)))
	v436 = v432
	v438 = int32(-1077)
	v439 = base.I32_wrap_i64(int64(base.Ui64(v432) >> (uint(int64(32)) % 64)))
	goto L109
L111:
	;
	if base.F64_eq(v394, float64(0)) != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v405) {
		v515 = v394
		goto L108
	} else {
		goto L118
	}
L114:
	;
	v537 = base.F64_div(float64(-1), base.F64_mul(v394, v394))
	goto L107
L115:
	;
	goto L116
L116:
	;
	if int64(0) <= v405 {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v537 = base.F64_div(base.F64_sub(v394, v394), float64(0))
	goto L107
L118:
	;
	v420 = int32(-1023)
	v422 = int64(base.Ui64(v405) >> (uint(int64(32)) % 64))
	if v422 != int64(1072693248) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v436 = v405
	v438 = v420
	v439 = base.I32_wrap_i64(v422)
	goto L109
L120:
	;
	goto L121
L121:
	;
	if base.I32_wrap_i64(v405) != 0 {
		v436 = v405
		v438 = v420
		v439 = int32(1072693248)
		goto L109
	} else {
		goto L122
	}
L122:
	;
	v537 = float64(0)
	goto L107
L123:
	;
	v541 = base.I32_trunc_f64_s(v537)
	v543 = v541
	goto L106
L124:
	;
	goto L125
L125:
	;
	v543 = int32(-2147483648)
	goto L106
L126:
	;
	F_pfree(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L101
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v721 = v545
	goto L1
L129:
	;
	goto L128
L130:
	;
	v553 = v18 << (uint(int32(2)) % 32)
	v554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17))))
	if v14 != int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v557 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+2)))
	v563 = v557 + v554*int32(10000)
	v564 = v553 - int32(4)
	goto L133
L132:
	;
	v563 = v554
	v564 = v553
	goto L133
L133:
	;
	v569 = F_log(m, base.F64_convert_i32_s(v563))
	mBase = m.M
	v571 = base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v564), float64(2.302585092994046)), v569))
	v582 = base.I64_reinterpret_f64(v571)
	if v582 <= int64(4503599627370495) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	if base.F64_lt(base.F64_abs(v714), float64(2.147483648e+09)) != 0 {
		goto L150
	} else {
		goto L151
	}
L135:
	;
	v714 = v692
	goto L134
L136:
	;
	v618 = v616 + int32(614242)
	v622 = base.F64_convert_i32_s(int32(base.Ui32(v618)>>(uint(int32(20))%32)) + v615)
	v624 = base.F64_mul(v622, float64(0.30102999566361177))
	v637 = base.F64_add(base.F64_reinterpret_i64(v613&int64(4294967295)|base.I64_extend_i32_u(v618&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v640 = base.F64_mul(v637, base.F64_mul(v637, float64(0.5)))
	v645 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v637, v640)) & int64(-4294967296))
	v646 = float64(0.4342944818781689)
	v647 = base.F64_mul(v645, v646)
	v648 = base.F64_add(v624, v647)
	v653 = base.F64_div(v637, base.F64_add(v637, float64(2)))
	v654 = base.F64_mul(v653, v653)
	v655 = base.F64_mul(v654, v654)
	v680 = base.F64_add(base.F64_mul(v653, base.F64_add(v640, base.F64_add(base.F64_mul(v655, base.F64_add(base.F64_mul(v655, base.F64_add(base.F64_mul(v655, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v654, base.F64_add(base.F64_mul(v655, base.F64_add(base.F64_mul(v655, base.F64_add(base.F64_mul(v655, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v637, v645), v640))
	v692 = base.F64_add(v648, base.F64_add(base.F64_add(v647, base.F64_sub(v624, v648)), base.F64_add(base.F64_mul(v680, v646), base.F64_add(base.F64_mul(v622, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v680, v645), float64(2.5082946711645275e-11))))))
	goto L135
L137:
	;
	v609 = base.I64_reinterpret_f64(base.F64_mul(v571, float64(1.8014398509481984e+16)))
	v613 = v609
	v615 = int32(-1077)
	v616 = base.I32_wrap_i64(int64(base.Ui64(v609) >> (uint(int64(32)) % 64)))
	goto L136
L138:
	;
	if base.F64_eq(v571, float64(0)) != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v582) {
		v692 = v571
		goto L135
	} else {
		goto L145
	}
L141:
	;
	v714 = base.F64_div(float64(-1), base.F64_mul(v571, v571))
	goto L134
L142:
	;
	goto L143
L143:
	;
	if int64(0) <= v582 {
		goto L137
	} else {
		goto L144
	}
L144:
	;
	v714 = base.F64_div(base.F64_sub(v571, v571), float64(0))
	goto L134
L145:
	;
	v597 = int32(-1023)
	v599 = int64(base.Ui64(v582) >> (uint(int64(32)) % 64))
	if v599 != int64(1072693248) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v613 = v582
	v615 = v597
	v616 = base.I32_wrap_i64(v599)
	goto L136
L147:
	;
	goto L148
L148:
	;
	if base.I32_wrap_i64(v582) != 0 {
		v613 = v582
		v615 = v597
		v616 = int32(1072693248)
		goto L136
	} else {
		goto L149
	}
L149:
	;
	v714 = float64(0)
	goto L134
L150:
	;
	v718 = base.I32_trunc_f64_s(v714)
	v721 = v718
	goto L1
L151:
	;
	goto L152
L152:
	;
	v721 = int32(-2147483648)
	goto L1
}
func F_examine_opclause_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v12 == int32(27) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v16 = v15
	} else {
		v16 = v11
	}
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v17 == int32(27) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v22 = v20
		v23 = v21
	} else {
		v22 = v10
		v23 = v17
	}
	if v23 == int32(7) {
		v29 = v22
		v30 = v16
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
		} else {
		}
		if l2 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
		} else {
		}
		v33 = int32(1)
		if l3 == int32(0) {
			v41 = v33
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(base.B2i32(v23 == int32(7)))
			v41 = v33
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		if v26 != int32(7) {
			v41 = int32(0)
		} else {
			v29 = v16
			v30 = v22
			if l1 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
			} else {
			}
			if l2 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
			} else {
			}
			v33 = int32(1)
			if l3 == int32(0) {
				v41 = v33
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(base.B2i32(v23 == int32(7)))
				v41 = v33
			}
		}
	}
	return v41
}
func F_executeComparison(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v324 int64
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int64
	_ = v364
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v486 int64
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v526 int64
	_ = v526
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v711 int64
	_ = v711
	var v712 int64
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v768 int64
	_ = v768
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int64
	_ = v792
	var v793 int64
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int64
	_ = v827
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v20 != v21 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(224)
	return v893
L2:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v30 = int32(2)
	switch v20 {
	case 0:
		v848 = v20
		goto L8
	case 1:
		goto L20
	case 2:
		goto L21
	case 3:
		goto L17
	default:
		goto L18
	case 16, 17, 18:
		v893 = v30
		goto L1
	case 32:
		goto L19
	}
L5:
	;
	v893 = base.B2i32(v19 == int32(9))
	goto L1
L6:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v893 = int32(2)
	goto L1
L8:
	;
	switch v19 - int32(8) {
	case 0:
		goto L317
	case 1:
		goto L323
	case 2:
		goto L322
	case 3:
		goto L321
	case 4:
		goto L320
	case 5:
		goto L319
	default:
		goto L318
	}
L9:
	;
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
	if v296 == int32(-2147483648) {
		goto L305
	} else {
		goto L306
	}
L10:
	;
	v825 = F_DirectFunctionCall2Coll(m, v821, int32(0), v823, v822)
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L22
	} else {
		goto L303
	}
L11:
	;
	if v294 <= int32(1183) {
		goto L271
	} else {
		goto L272
	}
L12:
	;
	if v293&int32(1) != 0 {
		goto L259
	} else {
		goto L260
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L22
	} else {
		goto L254
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L22
	} else {
		goto L249
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L22
	} else {
		goto L244
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L22
	} else {
		goto L239
	}
L17:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v602 != 0 {
		goto L233
	} else {
		goto L234
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L22
	} else {
		goto L230
	}
L19:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v297 <= int32(1183) {
		goto L128
	} else {
		goto L129
	}
L20:
	;
	if v19 == int32(8) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v35 = F_DirectFunctionCall2Coll(m, int32(1346), int32(0), v33, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v848 = v35
	goto L8
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v41 != v42 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v116 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	goto L49
L27:
	;
	v893 = int32(0)
	goto L1
L28:
	;
	goto L29
L29:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v41) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v893 = base.B2i32(v108 == int32(0))
	goto L1
L31:
	;
	v108 = int32(0)
	goto L30
L32:
	;
	v82 = v77
	v83 = v78
	v84 = v79
	goto L42
L33:
	;
	if (v45|v46)&int32(3) != 0 {
		v77 = v45
		v78 = v46
		v79 = v41
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v70 = v45
	v71 = v46
	v72 = v41
	goto L35
L35:
	;
	if v72 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L36:
	;
	v54 = v45
	v55 = v46
	v56 = v41
	goto L37
L37:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v59 != v60 {
		v77 = v54
		v78 = v55
		v79 = v56
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v70 = v65
	v71 = v63
	v72 = v67
	goto L35
L39:
	;
	v62 = int32(4)
	v63 = v55 + v62
	v65 = v54 + v62
	v67 = v56 - v62
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v54 = v65
		v55 = v63
		v56 = v67
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v77 = v70
	v78 = v71
	v79 = v72
	goto L32
L42:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == v88 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v108 = v87 - v88
	goto L30
L44:
	;
	v90 = int32(1)
	v95 = v84 - v90
	if v95 != 0 {
		v82 = v82 + v90
		v83 = v83 + v90
		v84 = v95
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	goto L31
L48:
	;
	v226 = base.B2i32(v113 < v111)
	if v113 < v111 {
		goto L95
	} else {
		goto L96
	}
L49:
	;
	if v117 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	goto L51
L51:
	;
	if v122 == int32(6) {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v126 = F_pg_server_to_any(m, v114, v113, int32(6))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L22
	} else {
		goto L53
	}
L53:
	;
	v129 = F_pg_server_to_any(m, v112, v111, int32(6))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L22
	} else {
		goto L54
	}
L54:
	;
	v131 = base.B2i32(v114 == v126)
	if v131 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v134 = F_strlen(m, v126)
	mBase = m.M
	v135 = v134
	goto L57
L56:
	;
	v135 = v113
	goto L57
L57:
	;
	v136 = base.B2i32(v112 == v129)
	if v136 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v139 = F_strlen(m, v129)
	mBase = m.M
	v140 = v139
	goto L60
L59:
	;
	v140 = v111
	goto L60
L60:
	;
	v141 = base.B2i32(v135 < v140)
	if v135 < v140 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v142 = v135
	goto L63
L62:
	;
	v142 = v140
	goto L63
L63:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v142) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	if v204 != 0 {
		goto L82
	} else {
		goto L83
	}
L65:
	;
	v204 = int32(0)
	goto L64
L66:
	;
	v178 = v173
	v179 = v174
	v180 = v175
	goto L76
L67:
	;
	if (v126|v129)&int32(3) != 0 {
		v173 = v126
		v174 = v129
		v175 = v142
		goto L66
	} else {
		goto L70
	}
L68:
	;
	v166 = v126
	v167 = v129
	v168 = v142
	goto L69
L69:
	;
	if v168 == int32(0) {
		goto L65
	} else {
		goto L75
	}
L70:
	;
	v150 = v126
	v151 = v129
	v152 = v142
	goto L71
L71:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v155 != v156 {
		v173 = v150
		v174 = v151
		v175 = v152
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v166 = v161
	v167 = v159
	v168 = v163
	goto L69
L73:
	;
	v158 = int32(4)
	v159 = v151 + v158
	v161 = v150 + v158
	v163 = v152 - v158
	if base.Ui32(int32(3)) < base.Ui32(v163) {
		v150 = v161
		v151 = v159
		v152 = v163
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v173 = v166
	v174 = v167
	v175 = v168
	goto L66
L76:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v183 == v184 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v204 = v183 - v184
	goto L64
L78:
	;
	v186 = int32(1)
	v191 = v180 - v186
	if v191 != 0 {
		v178 = v178 + v186
		v179 = v179 + v186
		v180 = v191
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	goto L65
L82:
	;
	v207 = v204
	goto L84
L83:
	;
	v207 = base.B2i32(v140 < v135) - v141
	goto L84
L84:
	;
	if v136&base.B2i32(v114 == v126) != 0 {
		v848 = v207
		goto L8
	} else {
		goto L85
	}
L85:
	;
	if v131 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_pfree(m, v126)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L22
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v136 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_pfree(m, v129)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L22
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v207 != 0 {
		v848 = v207
		goto L8
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	goto L48
L95:
	;
	v227 = v113
	goto L97
L96:
	;
	v227 = v111
	goto L97
L97:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v227) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v289 != 0 {
		goto L116
	} else {
		goto L117
	}
L99:
	;
	v289 = int32(0)
	goto L98
L100:
	;
	v263 = v258
	v264 = v259
	v265 = v260
	goto L110
L101:
	;
	if (v114|v112)&int32(3) != 0 {
		v258 = v114
		v259 = v112
		v260 = v227
		goto L100
	} else {
		goto L104
	}
L102:
	;
	v251 = v114
	v252 = v112
	v253 = v227
	goto L103
L103:
	;
	if v253 == int32(0) {
		goto L99
	} else {
		goto L109
	}
L104:
	;
	v235 = v114
	v236 = v112
	v237 = v227
	goto L105
L105:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	if v240 != v241 {
		v258 = v235
		v259 = v236
		v260 = v237
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v251 = v246
	v252 = v244
	v253 = v248
	goto L103
L107:
	;
	v243 = int32(4)
	v244 = v236 + v243
	v246 = v235 + v243
	v248 = v237 - v243
	if base.Ui32(int32(3)) < base.Ui32(v248) {
		v235 = v246
		v236 = v244
		v237 = v248
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v258 = v251
	v259 = v252
	v260 = v253
	goto L100
L110:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v268 == v269 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v289 = v268 - v269
	goto L98
L112:
	;
	v271 = int32(1)
	v276 = v265 - v271
	if v276 != 0 {
		v263 = v263 + v271
		v264 = v264 + v271
		v265 = v276
		goto L110
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	goto L111
L115:
	;
	goto L99
L116:
	;
	v292 = v289
	goto L118
L117:
	;
	v292 = base.B2i32(v111 < v113) - v226
	goto L118
L118:
	;
	v848 = v292
	goto L8
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L22
	} else {
		goto L227
	}
L120:
	;
	if v297 == int32(1114) {
		goto L11
	} else {
		goto L226
	}
L121:
	;
	if v294 <= int32(1183) {
		goto L193
	} else {
		goto L194
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L22
	} else {
		goto L187
	}
L123:
	;
	if v294 == int32(1114) {
		v893 = v30
		goto L1
	} else {
		goto L186
	}
L124:
	;
	if v293&int32(1) == int32(0) {
		goto L14
	} else {
		goto L184
	}
L125:
	;
	if v294 == int32(1184) {
		v893 = v30
		goto L1
	} else {
		goto L182
	}
L126:
	;
	if v294 <= int32(1183) {
		goto L171
	} else {
		goto L172
	}
L127:
	;
	if v294 <= int32(1183) {
		goto L136
	} else {
		goto L137
	}
L128:
	;
	switch v297 - int32(1082) {
	case 0:
		goto L127
	case 1:
		goto L126
	default:
		goto L120
	}
L129:
	;
	goto L130
L130:
	;
	if v297 == int32(1184) {
		goto L121
	} else {
		goto L131
	}
L131:
	;
	if v297 != int32(1266) {
		goto L119
	} else {
		goto L132
	}
L132:
	;
	v306 = int32(1446)
	if int32(1183) < v294 {
		goto L125
	} else {
		goto L133
	}
L133:
	;
	switch v294 - int32(1082) {
	case 0:
		v893 = v30
		goto L1
	case 1:
		goto L124
	default:
		goto L123
	}
L134:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L22
	} else {
		goto L166
	}
L135:
	;
	if v294 == int32(1114) {
		goto L9
	} else {
		goto L165
	}
L136:
	;
	switch v294 - int32(1082) {
	case 0:
		v821 = int32(1447)
		v822 = v295
		v823 = v296
		goto L10
	case 1:
		v893 = v30
		goto L1
	default:
		goto L135
	}
L137:
	;
	goto L138
L138:
	;
	if v294 != int32(1184) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v294 != int32(1266) {
		goto L134
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v293&int32(1) == int32(0) {
		goto L16
	} else {
		goto L143
	}
L142:
	;
	v893 = v30
	goto L1
L143:
	;
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
	v327 = m.G0
	v329 = v327 - int32(48)
	m.G0 = v329
	if v296 == int32(-2147483648) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v848 = v385
	goto L8
L145:
	;
	m.G0 = v329 + int32(48)
	goto L144
L146:
	;
	v334 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v324)
	mBase = m.M
	v385 = v334
	goto L145
L147:
	;
	goto L148
L148:
	;
	if v296 == int32(2147483647) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v383 = F_timestamp_cmp_internal(m, v382, v324)
	mBase = m.M
	v385 = v383
	goto L145
L150:
	;
	v382 = int64(9223372036854775807)
	goto L149
L151:
	;
	goto L152
L152:
	;
	if v296 <= int32(106751982) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v324 == int64(-9223372036854775807-1) {
		goto L162
	} else {
		goto L163
	}
L154:
	;
	F_j2date(m, v296+int32(2451545), v329+int32(24), v329+int32(20), v329+int32(16))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v329)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v329)+12)) = int32(0)
	v356 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v357 = F_DetermineTimeZoneOffset(m, v329+int32(4), v356)
	mBase = m.M
	v364 = base.I64_extend_i32_s(v357)*int64(1000000) + base.I64_extend_i32_s(v296)*int64(86400000000)
	if base.Ui64(v364+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v382 = v364
		goto L149
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v324 == int64(9223372036854775807) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	if v364 < int64(-211813488000000000) {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v376 = int32(-1)
	goto L161
L160:
	;
	v376 = int32(1)
	goto L161
L161:
	;
	v385 = v376
	goto L145
L162:
	;
	v381 = int32(1)
	goto L164
L163:
	;
	v381 = int32(-1)
	goto L164
L164:
	;
	v385 = v381
	goto L145
L165:
	;
	goto L134
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v294
	F_errmsg_internal(m, int32(64310), v17+int32(48))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L22
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(525321), int32(3756), int32(393033))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L22
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L22
	} else {
		goto L179
	}
L170:
	;
	if v294 == int32(1114) {
		v893 = v30
		goto L1
	} else {
		goto L178
	}
L171:
	;
	switch v294 - int32(1082) {
	case 0:
		v893 = v30
		goto L1
	case 1:
		v821 = int32(1448)
		v822 = v295
		v823 = v296
		goto L10
	default:
		goto L170
	}
L172:
	;
	goto L173
L173:
	;
	if v294 == int32(1184) {
		v893 = v30
		goto L1
	} else {
		goto L174
	}
L174:
	;
	if v294 != int32(1266) {
		goto L169
	} else {
		goto L175
	}
L175:
	;
	if v293&int32(1) == int32(0) {
		goto L15
	} else {
		goto L176
	}
L176:
	;
	v423 = F_DirectFunctionCall1Coll(m, int32(1435), int32(0), v296)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L22
	} else {
		goto L177
	}
L177:
	;
	v821 = int32(1446)
	v822 = v295
	v823 = v423
	goto L10
L178:
	;
	goto L169
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v294
	F_errmsg_internal(m, int32(64310), v17+int32(80))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L22
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(525321), int32(3782), int32(393033))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L22
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	if v294 != int32(1266) {
		goto L122
	} else {
		goto L183
	}
L183:
	;
	v821 = v306
	v822 = v295
	v823 = v296
	goto L10
L184:
	;
	v453 = F_DirectFunctionCall1Coll(m, int32(1435), int32(0), v295)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L22
	} else {
		goto L185
	}
L185:
	;
	v821 = v306
	v822 = v453
	v823 = v296
	goto L10
L186:
	;
	goto L122
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v294
	F_errmsg_internal(m, int32(64310), v17+int32(112))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L22
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(525321), int32(3808), int32(393033))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L22
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L22
	} else {
		goto L223
	}
L191:
	;
	if v294 == int32(1114) {
		goto L12
	} else {
		goto L222
	}
L192:
	;
	if v293&int32(1) == int32(0) {
		goto L13
	} else {
		goto L200
	}
L193:
	;
	switch v294 - int32(1082) {
	case 0:
		goto L192
	case 1:
		v893 = v30
		goto L1
	default:
		goto L191
	}
L194:
	;
	goto L195
L195:
	;
	if v294 == int32(1184) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v821 = int32(1449)
	v822 = v295
	v823 = v296
	goto L10
L197:
	;
	goto L198
L198:
	;
	if v294 != int32(1266) {
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v893 = v30
	goto L1
L200:
	;
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	v489 = m.G0
	v491 = v489 - int32(48)
	m.G0 = v491
	if v295 == int32(-2147483648) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v848 = int32(0) - v547
	goto L8
L202:
	;
	m.G0 = v491 + int32(48)
	goto L201
L203:
	;
	v496 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v486)
	mBase = m.M
	v547 = v496
	goto L202
L204:
	;
	goto L205
L205:
	;
	if v295 == int32(2147483647) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v545 = F_timestamp_cmp_internal(m, v544, v486)
	mBase = m.M
	v547 = v545
	goto L202
L207:
	;
	v544 = int64(9223372036854775807)
	goto L206
L208:
	;
	goto L209
L209:
	;
	if v295 <= int32(106751982) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v486 == int64(-9223372036854775807-1) {
		goto L219
	} else {
		goto L220
	}
L211:
	;
	F_j2date(m, v295+int32(2451545), v491+int32(24), v491+int32(20), v491+int32(16))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v491)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v491)+12)) = int32(0)
	v518 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v519 = F_DetermineTimeZoneOffset(m, v491+int32(4), v518)
	mBase = m.M
	v526 = base.I64_extend_i32_s(v519)*int64(1000000) + base.I64_extend_i32_s(v295)*int64(86400000000)
	if base.Ui64(v526+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v544 = v526
		goto L206
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if v486 == int64(9223372036854775807) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	if v526 < int64(-211813488000000000) {
		goto L210
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v538 = int32(-1)
	goto L218
L217:
	;
	v538 = int32(1)
	goto L218
L218:
	;
	v547 = v538
	goto L202
L219:
	;
	v543 = int32(1)
	goto L221
L220:
	;
	v543 = int32(-1)
	goto L221
L221:
	;
	v547 = v543
	goto L202
L222:
	;
	goto L190
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v294
	F_errmsg_internal(m, int32(64310), v17+int32(176))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L22
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(525321), int32(3866), int32(393033))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L22
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	goto L119
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v297
	F_errmsg_internal(m, int32(64310), v17+int32(32))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L22
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(525321), int32(3871), int32(393033))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L22
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v590
	F_errmsg_internal(m, int32(498692), v17)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L22
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(525321), int32(3403), int32(160509))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L22
	} else {
		goto L232
	}
L232:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L233:
	;
	v603 = int32(1)
	goto L235
L234:
	;
	v603 = int32(-1)
	goto L235
L235:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v602 != v605 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v607 = v603
	goto L238
L237:
	;
	v607 = int32(0)
	goto L238
L238:
	;
	v848 = v607
	goto L8
L239:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L22
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(7584)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = int32(375311)
	F_errmsg(m, int32(423149), v17-int32(-64))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L22
	} else {
		goto L241
	}
L241:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L22
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L22
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L22
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(7838)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(394484)
	F_errmsg(m, int32(423149), v17+int32(96))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L22
	} else {
		goto L246
	}
L246:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L22
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L22
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L22
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = int32(7838)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = int32(394484)
	F_errmsg(m, int32(423149), v17+int32(128))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L22
	} else {
		goto L251
	}
L251:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L22
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L22
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L22
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(7584)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = int32(375311)
	F_errmsg(m, int32(423149), v17+int32(192))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L22
	} else {
		goto L256
	}
L256:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L22
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L22
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	v711 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	v713 = F_timestamp_cmp_timestamptz_internal(m, v711, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L22
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L22
	} else {
		goto L263
	}
L262:
	;
	v848 = int32(0) - v713
	goto L8
L263:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L22
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = int32(7584)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = int32(249329)
	F_errmsg(m, int32(423149), v17+int32(208))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L22
	} else {
		goto L265
	}
L265:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L22
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L22
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	if v293&int32(1) != 0 {
		goto L294
	} else {
		goto L295
	}
L269:
	;
	v768 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	if v295 == int32(-2147483648) {
		goto L282
	} else {
		goto L283
	}
L270:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L22
	} else {
		goto L278
	}
L271:
	;
	switch v294 - int32(1082) {
	case 0:
		goto L269
	case 1:
		v893 = v30
		goto L1
	default:
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	if v294 == int32(1184) {
		goto L268
	} else {
		goto L276
	}
L274:
	;
	if v294 != int32(1114) {
		goto L270
	} else {
		goto L275
	}
L275:
	;
	v821 = int32(1449)
	v822 = v295
	v823 = v296
	goto L10
L276:
	;
	if v294 == int32(1266) {
		v893 = v30
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L270
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v294
	F_errmsg_internal(m, int32(64310), v17+int32(144))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L22
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(525321), int32(3837), int32(393033))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L22
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	v848 = int32(0) - v788
	goto L8
L282:
	;
	v772 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v768)
	mBase = m.M
	v788 = v772
	goto L281
L283:
	;
	goto L284
L284:
	;
	if v295 == int32(2147483647) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v776 = F_timestamp_cmp_internal(m, int64(9223372036854775807), v768)
	mBase = m.M
	v788 = v776
	goto L281
L286:
	;
	goto L287
L287:
	;
	if v295 <= int32(106751982) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v782 = F_timestamp_cmp_internal(m, base.I64_extend_i32_s(v295)*int64(86400000000), v768)
	mBase = m.M
	v788 = v782
	goto L281
L289:
	;
	goto L290
L290:
	;
	if v768 == int64(9223372036854775807) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v787 = int32(-1)
	goto L293
L292:
	;
	v787 = int32(1)
	goto L293
L293:
	;
	v788 = v787
	goto L281
L294:
	;
	v792 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	v793 = *(*int64)(unsafe.Add(mBase, uint32(v295)))
	v794 = F_timestamp_cmp_timestamptz_internal(m, v792, v793)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L22
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L22
	} else {
		goto L298
	}
L297:
	;
	v848 = v794
	goto L8
L298:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L22
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = int32(7584)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = int32(249329)
	F_errmsg(m, int32(423149), v17+int32(160))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L22
	} else {
		goto L300
	}
L300:
	;
	F_errhint(m, int32(607057), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L22
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(525321), int32(3672), int32(84231))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L22
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	v848 = v825
	goto L8
L304:
	;
	v848 = v847
	goto L8
L305:
	;
	v831 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v827)
	mBase = m.M
	v847 = v831
	goto L304
L306:
	;
	goto L307
L307:
	;
	if v296 == int32(2147483647) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v835 = F_timestamp_cmp_internal(m, int64(9223372036854775807), v827)
	mBase = m.M
	v847 = v835
	goto L304
L309:
	;
	goto L310
L310:
	;
	if v296 <= int32(106751982) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v841 = F_timestamp_cmp_internal(m, base.I64_extend_i32_s(v296)*int64(86400000000), v827)
	mBase = m.M
	v847 = v841
	goto L304
L312:
	;
	goto L313
L313:
	;
	if v827 == int64(9223372036854775807) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v846 = int32(-1)
	goto L316
L315:
	;
	v846 = int32(1)
	goto L316
L316:
	;
	v847 = v846
	goto L304
L317:
	;
	v893 = base.B2i32(v848 == int32(0))
	goto L1
L318:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L22
	} else {
		goto L324
	}
L319:
	;
	v893 = base.B2i32(int32(0) <= v848)
	goto L1
L320:
	;
	v893 = base.B2i32(v848 <= int32(0))
	goto L1
L321:
	;
	v893 = base.B2i32(int32(0) < v848)
	goto L1
L322:
	;
	v893 = int32(base.Ui32(v848) >> (uint(int32(31)) % 32))
	goto L1
L323:
	;
	v893 = base.B2i32(v848 != int32(0))
	goto L1
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v19
	F_errmsg_internal(m, int32(506048), v17+int32(16))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L22
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(525321), int32(3427), int32(160509))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L22
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expandNSItemVars(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	v6 = int32(0)
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v214
L5:
	;
	v31 = v6
	v32 = v6
	goto L10
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if int32(0) < v20 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v214 = v6
	goto L4
L9:
	;
	goto L8
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v41 = v38 + v31<<(uint(int32(5))%32)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+30)))
	if v42 != 0 {
		v195 = v32
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v214 = v195
	goto L4
L12:
	;
	v202 = v31 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v202 < v203 {
		v31 = v202
		v32 = v195
		goto L10
	} else {
		goto L38
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v31<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 == int32(0) {
		v195 = v32
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v57 = F_makeVar(m, v52, v53, v54, v55, v56, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+36)) = v63
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+44)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v57)+40)) = uint16(v65)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	if v69 == int32(0) {
		v147 = l0
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v68 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v73 = v69 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v69) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v84 = l0
	v85 = int32(0)
	goto L22
L20:
	;
	v110 = l0
	goto L21
L21:
	;
	v120 = int32(0)
	if v73 == v120 {
		v147 = v110
		goto L17
	} else {
		goto L25
	}
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v103 = v85 + int32(8)
	if v103 != v69&int32(-8) {
		v84 = v101
		v85 = v103
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v110 = v101
	goto L21
L24:
	;
	goto L23
L25:
	;
	v128 = v110
	v129 = v120
	goto L26
L26:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v140 = v129 + int32(1)
	if v140 != v73 {
		v128 = v138
		v129 = v140
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v147 = v138
	goto L17
L28:
	;
	goto L27
L29:
	;
	v178 = F_lappend(m, v32, v57)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L35
	}
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	if v159 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v162 < v68 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164+v68<<(uint(int32(2))%32)-int32(4))))
	if v170 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v174 = F_bms_union(m, v173, v170)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v174
	goto L29
L35:
	;
	if l4 == int32(0) {
		v195 = v178
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v183 = F_lappend(m, v182, v47)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v183
	v195 = v178
	goto L12
L38:
	;
	goto L11
}
func F_expand_generated_columns_in_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == v4 {
		v101 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v101
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)))
	if v15 != int32(1) {
		v101 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_palloc0(m, int32(136))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(101)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v29 = F_makeAlias(m, v25+int32(4), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v37 == v31 {
		v101 = l0
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+18)))
	if v40 != int32(1) {
		v101 = l0
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if int32(0) < v43 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v52 = int32(0)
	v53 = v43
	v56 = v4
	goto L12
L10:
	;
	v93 = v4
	goto L11
L11:
	;
	v96 = int32(0)
	v99 = F_ReplaceVarsFromTargetList(m, l0, l2, v19, v93, v96, int32(1), l2, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L22
	}
L12:
	;
	v62 = v52 + int32(1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52*int32(100)+(v36+int32(110)+v53<<(uint(int32(4))%32))))))
	if v67 == int32(118) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v93 = v84
	goto L11
L14:
	;
	v70 = F_build_generation_expression(m, l1, v62)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v83 = v53
	v84 = v56
	goto L16
L16:
	;
	if v62 < v83 {
		v52 = v62
		v53 = v83
		v56 = v84
		goto L12
	} else {
		goto L21
	}
L17:
	;
	F_ChangeVarNodes(m, v70, int32(1), l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v76 = int32(0)
	v78 = F_makeTargetEntry(m, v70, base.I32_extend16_s(v62), v76, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v80 = F_lappend(m, v56, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v83 = v82
	v84 = v80
	goto L16
L21:
	;
	goto L13
L22:
	;
	v101 = v99
	goto L1
}
func F_expand_virtual_generated_columns(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v18 == v2 {
		v202 = v17
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(48)
	return v202
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		v202 = v17
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = v2
	v29 = v17
	goto L4
L4:
	;
	v37 = v25 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v25<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if v43 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v202 = v188
	goto L1
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v48 = F_table_open(m, v46, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v188 = v29
	goto L8
L8:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v37 < v195 {
		v25 = v37
		v29 = v188
		goto L4
	} else {
		goto L41
	}
L9:
	;
	F_sequence_close(m, v48, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L40
	}
L10:
	;
	return int32(0)
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v53 == int32(0) {
		v173 = v29
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+18)))
	if v56 != int32(1) {
		v173 = v29
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v60 <= v59 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = l0
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v139
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v138
	if v126 != 0 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v126 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v68 = v59
	v70 = int32(0)
	v71 = v60
	goto L18
L18:
	;
	v84 = v52 + int32(20) + v71<<(uint(int32(4))%32) + v68*int32(100)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+90)))
	if v85 == int32(118) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v126 = v120
	goto L14
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v117 < v121 {
		v68 = v117
		v70 = v120
		v71 = v121
		goto L18
	} else {
		goto L31
	}
L21:
	;
	v89 = v68 + int32(1)
	v90 = F_build_generation_expression(m, v48, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v103 = v68 + int32(1)
	v104 = base.I32_extend16_s(v103)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+68))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v84)+76))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+96))
	v109 = F_makeVar(m, v37, v104, v105, v106, v107, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L28
	}
L24:
	;
	F_ChangeVarNodes(m, v90, int32(1), v37)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v96 = int32(0)
	v98 = F_makeTargetEntry(m, v90, base.I32_extend16_s(v89), v96, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v100 = F_lappend(m, v70, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v117 = v89
	v120 = v100
	goto L20
L28:
	;
	v111 = int32(0)
	v113 = F_makeTargetEntry(m, v109, v104, v111, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v115 = F_lappend(m, v70, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v117 = v103
	v120 = v115
	goto L20
L31:
	;
	goto L19
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v153 = v147<<(uint(int32(2))%32) + int32(4)
	goto L34
L33:
	;
	v153 = int32(4)
	goto L34
L34:
	;
	v154 = F_palloc0(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	if v157 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(1)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v166 = F_replace_rte_variables(m, v29, v160, int32(0), int32(851), v15+int32(8), v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v173 = v166
	goto L9
L40:
	;
	v188 = v173
	goto L8
L41:
	;
	goto L5
}
