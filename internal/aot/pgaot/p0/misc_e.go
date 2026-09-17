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
	if v69 != int32(_a_F_ER_get_flat_size_0) {
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
						F_errmsg(m, int32(_a_F_ExecARDeleteTriggers_0), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecARDeleteTriggers_1), int32(2817), int32(_a_F_ExecARDeleteTriggers_2))
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
	F_errmsg(m, int32(_a_F_ExecARUpdateTriggers_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ExecARUpdateTriggers_1), int32(3164), int32(_a_F_ExecARUpdateTriggers_2))
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
	var v42 int32
	_ = v42
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
	v42 = v3
	goto L9
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v46 = v43 + v42*int32(60)
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
	v74 = v42 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v74 < v75 {
		v42 = v74
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
	v70 = F_ExecCallTriggerFunc(m, v9+int32(4), v42, v63, v64, v69)
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
	F_errmsg(m, int32(_a_F_ExecBSDeleteTriggers_0), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ExecBSDeleteTriggers_1), int32(2677), int32(_a_F_ExecBSDeleteTriggers_2))
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
		v10 = int32(_a_F_ExecGetRootToChildMap_0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v18
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
					*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v11
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v35)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					return v42
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_ExecGetRootToChildMap[0])) = v11
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L14
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+17)))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+18)))
	if v29 != int32(1) {
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
		v41 = int32(0)
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v42 = int32(_a_F_ExecInitGenerated_0)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0])) = v45
	v49 = F_palloc0(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L16
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+13)))
	if v37 != 0 {
		v41 = int32(0)
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v38 = F_ExecGetUpdatedCols(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v41 = v38
	goto L8
L16:
	;
	if int32(0) < v25 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l2 == int32(2) {
		goto L55
	} else {
		goto L56
	}
L18:
	;
	v57 = int32(0)
	v59 = v4
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_pfree(m, v49)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L53
	}
L21:
	;
	v70 = v57 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v71<<(uint(int32(4))%32)+v57*int32(100))+110)))
	if v78 == int32(0) {
		v160 = v59
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v160 != 0 {
		v187 = v160
		v188 = v49
		goto L17
	} else {
		goto L52
	}
L23:
	;
	if v70 != v25 {
		v57 = v70
		v59 = v160
		goto L21
	} else {
		goto L51
	}
L24:
	;
	v81 = F_build_column_default(m, v20, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v41 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
	F_pull_varattnos(m, v81, int32(1), v18+int32(12))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v78 == int32(115) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v93 = int32(0)
	if base.B2i32(v41 == v93)|base.B2i32(v92 == v93) != 0 {
		v138 = v93
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v138 == int32(0) {
		v160 = v59
		goto L23
	} else {
		goto L44
	}
L32:
	;
	goto L31
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v103 < v104 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v106 = v103
	goto L36
L35:
	;
	v106 = v104
	goto L36
L36:
	;
	if v106 <= int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v109 = int32(1)
	goto L39
L38:
	;
	v109 = v106
	goto L39
L39:
	;
	v110 = int32(8)
	v115 = int32(0)
	goto L40
L40:
	;
	v122 = v115 << (uint(int32(2)) % 32)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v92+v110+v122)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v41+v110+v122)))
	v127 = v124 & v126
	v129 = base.B2i32(v127 != int32(0))
	if v127 != 0 {
		v138 = v129
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v138 = v129
	goto L32
L42:
	;
	v131 = v115 + int32(1)
	if v131 != v109 {
		v115 = v131
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L29
L45:
	;
	v146 = F_ExecPrepareExpr(m, v81, l1)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	v151 = v59
	goto L47
L47:
	;
	if l2 != int32(2) {
		v160 = v151
		goto L23
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49+v57<<(uint(int32(2))%32)))) = v146
	v151 = v59 + int32(1)
	goto L47
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v157 = F_bms_add_member(m, v154, v57+int32(8))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v157
	v160 = v151
	goto L23
L51:
	;
	goto L22
L52:
	;
	goto L20
L53:
	;
	v180 = int32(0)
	v187 = v180
	v188 = v180
	goto L17
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInitGenerated[0])) = v43
	goto L2
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v188
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v201)
	goto L54
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v188
	goto L54
L58:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v229 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecInitGenerated_1), v18)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_ExecInitGenerated_2), int32(479), int32(_a_F_ExecInitGenerated_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	v6 = F_MakeTupleTableSlot(m, l1, int32(_a_F_ExecInitNullTupleSlot_0))
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
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v24 = v22 << (uint(int32(2)) % 32)
				if v18&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v24)) == int32(0) {
					if v24 == int32(0) {
					} else {
						v34 = v24 + v18
						v36 = v18 + int32(4)
						if base.Ui32(v36) < base.Ui32(v34) {
							v38 = v34
						} else {
							v38 = v36
						}
						v44 = (v18^int32(-1)+v38)&int32(-4) + int32(4)
						if v44 == int32(0) {
						} else {
							base.MemoryFill(m, v18, int32(0), v44)
						}
					}
				} else {
					v44 = v24
					if v44 == int32(0) {
					} else {
						base.MemoryFill(m, v18, int32(0), v44)
					}
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				if v53 != 0 {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
					base.MemoryFill(m, v54, int32(1), v53)
				} else {
				}
				v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
				v59 = v57 & int32(_a_F_ExecInitNullTupleSlot_1)
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)) = uint16(v59)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
				*(*uint16)(unsafe.Add(mBase, uint32(v6)+6)) = uint16(v62)
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
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMaterial[0]))
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
	v56 = int32(0)
	if base.B2i32(v53 == v56)|base.B2i32(v14 == int32(1)) == v56 {
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
	v54 = v21
	v55 = v2
	goto L6
L11:
	;
	goto L12
L12:
	;
	v23 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMaterial[1]))
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
		v54 = v23
		v55 = v2
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
	v54 = int32(0)
	v55 = v42
	goto L6
L21:
	;
	return v115
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	m.T0[v112].(func(*base.Module, int32))(m, v110)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L53
	}
L23:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v84 != 0 {
		v110 = v83
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v79 = F_tuplestore_gettupleslot(m, v55, base.B2i32(v14 == int32(1)), int32(0), v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v63 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v53 != 0 {
		v83 = v74
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v66 = int32(0)
	v68 = F_tuplestore_advance(m, v55, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v75 = v73
	goto L24
L31:
	;
	if v68 == int32(0) {
		v115 = v66
		goto L21
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v75 = v74
	goto L24
L34:
	;
	if v79 != 0 {
		v115 = v75
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v14 != int32(1) {
		v110 = v75
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v83 = v75
	goto L23
L37:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	if v86 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_ExecReScan(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = m.T0[v89].(func(*base.Module, int32) int32)(m, v85)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	if v54 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v90 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v92&int32(2) == int32(0) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v97 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)) = uint8(v97)
	return int32(0)
L47:
	;
	goto L46
L48:
	;
	F_tuplestore_puttupleslot(m, v55, v90)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+32))
	m.T0[v106].(func(*base.Module, int32, int32))(m, v83, v90)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	return v83
L53:
	;
	v115 = v110
	goto L21
}
func F_ExecMaterializesOutput(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = l0 - int32(348)
	return base.B2i32(base.Ui32(v3) < base.Ui32(int32(15))) & int32(base.Ui32(int32(_a_F_ExecMaterializesOutput_0))>>(uint(v3)%32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v25+v45))) = v50
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
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v11 = int32(0)
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
	if v19 <= v11 {
		v25 = int32(0)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = v21 + v11<<(uint(int32(2))%32)
	goto L3
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33+v11<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+108)))
	F_ExecBatchInsert(m, v46, v47, v48, v49, v50, l0, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L14
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	F_list_free(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if base.B2i32(v25 == int32(0))|base.B2i32(v30 <= v11) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v33 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	return
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	F_list_free(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+188)) = int64(0)
	return
L14:
	;
	v11 = v11 + int32(1)
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
						v82 = int32(_a_F_ExecProcessReturning_0)
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
							v82 = int32(_a_F_ExecProcessReturning_0)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
								v82 = int32(_a_F_ExecProcessReturning_0)
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
								v82 = int32(_a_F_ExecProcessReturning_0)
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
									v82 = int32(_a_F_ExecProcessReturning_0)
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
										v82 = int32(_a_F_ExecProcessReturning_0)
										v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
										v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
										v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
											v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
							v82 = int32(_a_F_ExecProcessReturning_0)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
								v82 = int32(_a_F_ExecProcessReturning_0)
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
									v82 = int32(_a_F_ExecProcessReturning_0)
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
				F_errmsg_internal(m, int32(_a_F_ExecProcessReturning_2), v12)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ExecProcessReturning_3), int32(316), int32(_a_F_ExecProcessReturning_4))
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
							v82 = int32(_a_F_ExecProcessReturning_0)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
								v82 = int32(_a_F_ExecProcessReturning_0)
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
									v82 = int32(_a_F_ExecProcessReturning_0)
									v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
									v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
										v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
						v82 = int32(_a_F_ExecProcessReturning_0)
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
							v82 = int32(_a_F_ExecProcessReturning_0)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
								v82 = int32(_a_F_ExecProcessReturning_0)
								v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
								v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
									v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
					v82 = int32(_a_F_ExecProcessReturning_0)
					v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
					v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
						v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
						v82 = int32(_a_F_ExecProcessReturning_0)
						v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
							v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
							v82 = int32(_a_F_ExecProcessReturning_0)
							v83 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0]))
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v85
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, v15+int32(4), v76, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecProcessReturning[0])) = v83
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)))
								v97 = v95 & int32(_a_F_ExecProcessReturning_1)
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v200 int32
	_ = v200
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSecLabelStmt[0]))
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
	v513 = m.ExcPending
	if v513 != 0 {
		goto L29
	} else {
		goto L150
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L29
	} else {
		goto L146
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L29
	} else {
		goto L142
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L29
	} else {
		goto L138
	}
L5:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v87 - int32(1) {
	case 0, 5, 8, 11, 13, 17, 18, 20, 21, 22, 28, 29, 32, 33, 35, 36, 37, 40, 41, 48, 50:
		goto L27
	default:
		goto L28
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
	v81 = v24
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
	v39 = v3
	goto L16
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34+v39<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if base.B2i32(v50 == int32(0))|base.B2i32(v50 != v53) != 0 {
		v71 = v50
		v72 = v53
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L2
L18:
	;
	if v71-v72 == int32(0) {
		v81 = v46
		goto L5
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v56 = v15
	v57 = v47
	goto L21
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v61 == int32(0) {
		v71 = v61
		v72 = v60
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v71 = v61
	v72 = v60
	goto L19
L23:
	;
	v64 = int32(1)
	if v61 == v60 {
		v56 = v56 + v64
		v57 = v57 + v64
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v77 = v39 + int32(1)
	if v33 != v77 {
		v39 = v77
		goto L16
	} else {
		goto L26
	}
L26:
	;
	goto L17
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_get_object_address(m, l0, v87, v106, v11+int32(44), int32(4), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L29
	} else {
		goto L34
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_0), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(160), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSecLabelStmt[1]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v119
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_check_object_ownership(m, v114, v116, v11+int32(16), v115, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v126 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+119)))
	v133 = v131 - int32(99)
	if base.B2i32(base.Ui32(int32(19)) < base.Ui32(v133))|base.B2i32(int32(1)<<(uint(v133)%32)&int32(_a_F_ExecSecLabelStmt_3) == int32(0)) != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	m.T0[v145].(func(*base.Module, int32, int32))(m, l0, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L29
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v150 = int32(0)
	v151 = m.G0
	v153 = v151 - int32(240)
	m.G0 = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = int32(1)
	if v155 <= int32(3591) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	F_relation_close(m, v434, int32(3))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L29
	} else {
		goto L133
	}
L42:
	;
	if v229 != 0 {
		goto L67
	} else {
		goto L68
	}
L43:
	;
	goto L42
L44:
	;
	v229 = int32(0)
	goto L43
L45:
	;
	if base.B2i32(base.Ui32(v155-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v155-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v229 = v158
		goto L43
	} else {
		goto L66
	}
L46:
	;
	if v155 <= int32(2670) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	if v155 <= int32(_a_F_ExecSecLabelStmt_4) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	switch v155 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v229 = v158
		goto L43
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L44
	default:
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v170 = v155 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v170))|base.B2i32(int32(1)<<(uint(v170)%32)&int32(226492515) == int32(0)) != 0 {
		goto L45
	} else {
		goto L54
	}
L52:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v155-int32(2396)) {
		goto L44
	} else {
		goto L53
	}
L53:
	;
	v229 = v158
	goto L43
L54:
	;
	v229 = v158
	goto L43
L55:
	;
	if base.Ui32(v155-int32(3592)) < base.Ui32(int32(2)) {
		v229 = v158
		goto L43
	} else {
		goto L64
	}
L56:
	;
	v183 = v155 - int32(_a_F_ExecSecLabelStmt_5)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v183))|base.B2i32(int32(1)<<(uint(v183)%32)&int32(963) == int32(0)) != 0 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	switch v155 - int32(_a_F_ExecSecLabelStmt_6) {
	case 0, 1, 2, 3, 4, 59, 60:
		v229 = v158
		goto L43
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L44
	default:
		goto L60
	}
L59:
	;
	v229 = v158
	goto L43
L60:
	;
	if base.Ui32(v155-int32(_a_F_ExecSecLabelStmt_7)) < base.Ui32(int32(3)) {
		v229 = v158
		goto L43
	} else {
		goto L61
	}
L61:
	;
	v200 = v155 - int32(_a_F_ExecSecLabelStmt_8)
	if base.Ui32(int32(15)) < base.Ui32(v200) {
		goto L44
	} else {
		goto L62
	}
L62:
	;
	if int32(1)<<(uint(v200)%32)&int32(_a_F_ExecSecLabelStmt_9) != 0 {
		v229 = v158
		goto L43
	} else {
		goto L63
	}
L63:
	;
	goto L44
L64:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v155-int32(4060)) {
		goto L44
	} else {
		goto L65
	}
L65:
	;
	v229 = v158
	goto L43
L66:
	;
	goto L44
L67:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v230
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v236
	v238 = F_cstring_to_text(m, v148)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L29
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+12)) = uint8(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v325
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+4)) = uint8(v325)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v335
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v337
	v339 = F_cstring_to_text(m, v148)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L29
	} else {
		goto L101
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v238
	if v149 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v241 = F_cstring_to_text(m, v149)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L29
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v245 = v153 + int32(48)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v245, int32(1), int32(3), int32(184), v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L29
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+28)) = v241
	goto L73
L75:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v153+int32(96), int32(2), int32(3), int32(184), v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L29
	} else {
		goto L76
	}
L76:
	;
	v262 = int32(3)
	v265 = F_cstring_to_text(m, v148)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L29
	} else {
		goto L77
	}
L77:
	;
	F_ScanKeyInit(m, v153+int32(144), v262, v262, int32(67), v265)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	v271 = F_table_open(m, int32(3592), int32(3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L29
	} else {
		goto L80
	}
L79:
	;
	v434 = v271
	goto L41
L80:
	;
	v277 = F_systable_beginscan(m, v271, int32(3593), int32(1), int32(0), int32(3), v245)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L29
	} else {
		goto L81
	}
L81:
	;
	v279 = F_systable_getnext(m, v277)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L29
	} else {
		goto L82
	}
L82:
	;
	if v279 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v149 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v302 = v150
	goto L85
L85:
	;
	F_systable_endscan(m, v277)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L29
	} else {
		goto L93
	}
L86:
	;
	F_simple_heap_delete(m, v271, v279+int32(4))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L29
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v289 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+3)) = uint8(v289)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v271)+52))
	v298 = F_heap_modify_tuple(m, v279, v293, v153+int32(16), v153+int32(8), v153)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L29
	} else {
		goto L91
	}
L89:
	;
	F_systable_endscan(m, v277)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L29
	} else {
		goto L90
	}
L90:
	;
	goto L79
L91:
	;
	F_CatalogTupleUpdate(m, v271, v279+int32(4), v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L29
	} else {
		goto L92
	}
L92:
	;
	v302 = v298
	goto L85
L93:
	;
	v305 = int32(0)
	if base.B2i32(v149 == v305)|v302 == v305 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v271)+52))
	v315 = F_heap_form_tuple(m, v310, v153+int32(16), v153+int32(8))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L29
	} else {
		goto L97
	}
L95:
	;
	v319 = v302
	goto L96
L96:
	;
	if v319 == int32(0) {
		goto L79
	} else {
		goto L99
	}
L97:
	;
	F_CatalogTupleInsert(m, v271, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L29
	} else {
		goto L98
	}
L98:
	;
	v319 = v315
	goto L96
L99:
	;
	F_pfree(m, v319)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L29
	} else {
		goto L100
	}
L100:
	;
	goto L79
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+28)) = v339
	if v149 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v342 = F_cstring_to_text(m, v149)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L29
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v346 = v153 + int32(48)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v346, int32(1), int32(3), int32(184), v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L29
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+32)) = v342
	goto L104
L106:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v153+int32(96), int32(2), int32(3), int32(184), v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L29
	} else {
		goto L107
	}
L107:
	;
	v363 = int32(3)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ScanKeyInit(m, v153+int32(144), v363, v363, int32(65), v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L29
	} else {
		goto L108
	}
L108:
	;
	v374 = F_cstring_to_text(m, v148)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L29
	} else {
		goto L109
	}
L109:
	;
	F_ScanKeyInit(m, v153+int32(192), int32(4), int32(3), int32(67), v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L29
	} else {
		goto L110
	}
L110:
	;
	v380 = F_table_open(m, int32(3596), int32(3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L29
	} else {
		goto L112
	}
L111:
	;
	v434 = v380
	goto L41
L112:
	;
	v386 = F_systable_beginscan(m, v380, int32(3597), int32(1), int32(0), int32(4), v346)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L29
	} else {
		goto L113
	}
L113:
	;
	v388 = F_systable_getnext(m, v386)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L29
	} else {
		goto L114
	}
L114:
	;
	if v388 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v149 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v411 = v150
	goto L117
L117:
	;
	F_systable_endscan(m, v386)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L29
	} else {
		goto L125
	}
L118:
	;
	F_simple_heap_delete(m, v380, v388+int32(4))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L29
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v398 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+4)) = uint8(v398)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v380)+52))
	v407 = F_heap_modify_tuple(m, v388, v402, v153+int32(16), v153+int32(8), v153)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L29
	} else {
		goto L123
	}
L121:
	;
	F_systable_endscan(m, v386)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L29
	} else {
		goto L122
	}
L122:
	;
	goto L111
L123:
	;
	F_CatalogTupleUpdate(m, v380, v388+int32(4), v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L29
	} else {
		goto L124
	}
L124:
	;
	v411 = v407
	goto L117
L125:
	;
	v414 = int32(0)
	if base.B2i32(v149 == v414)|v411 == v414 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v380)+52))
	v424 = F_heap_form_tuple(m, v419, v153+int32(16), v153+int32(8))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L29
	} else {
		goto L129
	}
L127:
	;
	v428 = v411
	goto L128
L128:
	;
	if v428 == int32(0) {
		goto L111
	} else {
		goto L131
	}
L129:
	;
	F_CatalogTupleInsert(m, v380, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L29
	} else {
		goto L130
	}
L130:
	;
	v428 = v424
	goto L128
L131:
	;
	F_pfree(m, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L29
	} else {
		goto L132
	}
L132:
	;
	goto L111
L133:
	;
	m.G0 = v153 + int32(240)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v444 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_relation_close(m, v444, int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L29
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	m.G0 = v11 + int32(48)
	return
L137:
	;
	goto L136
L138:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L29
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_10), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L29
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(131), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L29
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L29
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_11), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L29
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(135), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L29
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L29
	} else {
		goto L147
	}
L147:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v498
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_12), v11+int32(32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L29
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(154), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L29
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L29
	} else {
		goto L151
	}
L151:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v518 + int32(4)
	F_errmsg(m, int32(_a_F_ExecSecLabelStmt_13), v11)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L29
	} else {
		goto L152
	}
L152:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+48))
	v527 = int32(*(*int8)(unsafe.Add(mBase, uint32(v526)+119)))
	F_errdetail_relkind_not_supported(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L29
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_ExecSecLabelStmt_1), int32(195), int32(_a_F_ExecSecLabelStmt_2))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L29
	} else {
		goto L154
	}
L154:
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
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[0]))
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
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[1]))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
			if v18 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l3))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v18)) == int32(0) {
					v30 = base.B2i32(base.Ui32(v18) < base.Ui32(l3))
				} else {
					v30 = int32(base.Ui32(v18-l3) >> (uint(int32(31)) % 32))
				}
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[1]))
				if v30 == int32(0) {
					v41 = v32
				} else {
					v35 = v32
					*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
					v41 = v35
				}
			} else {
				v35 = v16
				*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v17 + base.I64_extend_i32_s(l3-base.I32_wrap_i64(v17))
				v41 = v35
			}
			v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)+56))
			*(*int64)(unsafe.Add(mBase, uint32(v41)+56)) = v42 + int64(1)
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_ExpireTreeKnownAssignedTransactionIds[0]))
			F_LWLockRelease(m, v47+int32(512))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				return
			}
		}
	}
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v54 int64
	_ = v54
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	v3 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = base.I64_reinterpret_f64(l1)
	v15 = v13 & int64(4503599627370495)
	v19 = int64(base.Ui64(v13)>>(uint(int64(52))%64)) & int64(2047)
	if v19 != v3 {
		if v19 != int64(2047) {
			v74 = v19 + int64(15360)
			v75 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v77 = v15 << (uint(int64(60)) % 64)
		} else {
			v74 = int64(32767)
			v75 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
			v77 = v15 << (uint(int64(60)) % 64)
		}
	} else {
		if v15 == int64(0) {
			v37 = int64(0)
			v74 = v37
			v75 = v3
			v77 = v37
		} else {
			v39 = int64(0)
			v41 = base.I32_wrap_i64(base.I64_clz(v15))
			v43 = v41 + int32(49)
			if v43&int32(64) != 0 {
				v62 = int64(0)
				v63 = v15 << (uint(base.I64_extend_i32_u(v41+int32(-15))) % 64)
			} else {
				if v43 == int32(0) {
					v62 = v15
					v63 = v39
				} else {
					v54 = base.I64_extend_i32_u(v43)
					v62 = v15 << (uint(v54) % 64)
					v63 = v39<<(uint(v54)%64) | int64(base.Ui64(v15)>>(uint(base.I64_extend_i32_u(int32(64)-v43))%64))
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v62
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v63
			v67 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
			v73 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
			v74 = base.I64_extend_i32_u(int32(_a_F___extenddftf2_0) - v41)
			v75 = v67 ^ int64(281474976710656)
			v77 = v73
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v77
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v13&int64(-9223372036854775807-1) | v74<<(uint(int64(48))%64) | v75
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
				F_errmsg(m, int32(_a_F_each_array_start_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_each_array_start_1), int32(2176), int32(_a_F_each_array_start_2))
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
				F_errmsg(m, int32(_a_F_each_scalar_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_each_scalar_1), int32(2190), int32(_a_F_each_scalar_2))
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
	var v53 int32
	_ = v53
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
		v53 = v21
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v59 = v53
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
		v53 = v45
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v53 = v45
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
		v15 = int32(_a_F_elements_array_element_end_0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v18
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
							*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
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
							*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
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
								*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
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
								*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
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
						*(*int32)(unsafe.Add(mBase, _c_F_elements_array_element_end[0])) = v16
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13870(m, l0, int32(_a_F_elements_object_start_0), int32(2427), int32(_a_F_elements_object_start_1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
				F_errmsg(m, int32(_a_F_elements_scalar_0), v6)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_elements_scalar_1), int32(2442), int32(_a_F_elements_scalar_2))
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
	var v21 float64
	_ = v21
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 float32
	_ = v66
	var v68 float32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 float64
	_ = v111
	var v114 int32
	_ = v114
	var v121 float64
	_ = v121
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v130 float64
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 float32
	_ = v142
	var v144 float32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v176 int32
	_ = v176
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 float64
	_ = v187
	var v190 int32
	_ = v190
	var v197 float64
	_ = v197
	var v200 float64
	_ = v200
	var v201 int32
	_ = v201
	var v206 float64
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 float32
	_ = v333
	var v334 float32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v380 float64
	_ = v380
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 float32
	_ = v477
	var v478 int32
	_ = v478
	var v480 float32
	_ = v480
	var v491 int32
	_ = v491
	var v506 float64
	_ = v506
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 float64
	_ = v525
	var v533 float64
	_ = v533
	var v534 float64
	_ = v534
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v569 float64
	_ = v569
	var v570 float64
	_ = v570
	var v587 float32
	_ = v587
	var v588 float64
	_ = v588
	var v591 int32
	_ = v591
	var v592 float64
	_ = v592
	var v594 int32
	_ = v594
	var v598 float32
	_ = v598
	var v599 float64
	_ = v599
	var v602 int32
	_ = v602
	var v603 float64
	_ = v603
	var v605 float64
	_ = v605
	var v607 float64
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v635 float64
	_ = v635
	var v636 float64
	_ = v636
	var v653 float32
	_ = v653
	var v654 float64
	_ = v654
	var v657 int32
	_ = v657
	var v658 float64
	_ = v658
	var v660 float64
	_ = v660
	var v681 float64
	_ = v681
	var v682 float64
	_ = v682
	var v696 float64
	_ = v696
	var v704 float64
	_ = v704
	var v707 float64
	_ = v707
	var v736 float64
	_ = v736
	var v738 float64
	_ = v738
	var v739 float64
	_ = v739
	var v740 float64
	_ = v740
	var v747 float64
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v754 float64
	_ = v754
	var v761 float64
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v785 float64
	_ = v785
	var v786 float64
	_ = v786
	var v803 float32
	_ = v803
	var v804 float64
	_ = v804
	var v807 int32
	_ = v807
	var v808 float64
	_ = v808
	var v810 int32
	_ = v810
	var v814 float32
	_ = v814
	var v815 float64
	_ = v815
	var v818 int32
	_ = v818
	var v819 float64
	_ = v819
	var v821 float64
	_ = v821
	var v823 float64
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v851 float64
	_ = v851
	var v852 float64
	_ = v852
	var v869 float32
	_ = v869
	var v870 float64
	_ = v870
	var v873 int32
	_ = v873
	var v874 float64
	_ = v874
	var v876 float64
	_ = v876
	var v897 float64
	_ = v897
	var v898 float64
	_ = v898
	var v912 float64
	_ = v912
	var v920 float64
	_ = v920
	var v923 float64
	_ = v923
	var v948 float64
	_ = v948
	var v953 float64
	_ = v953
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 float64
	_ = v965
	var v968 float64
	_ = v968
	var v969 float64
	_ = v969
	var v974 float64
	_ = v974
	var v982 float64
	_ = v982
	var v990 float64
	_ = v990
	var v991 int32
	_ = v991
	var v992 float64
	_ = v992
	var v998 float64
	_ = v998
	var v1005 float64
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 float64
	_ = v1007
	var v1013 float64
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1022 float64
	_ = v1022
	var v1024 float64
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1063 float32
	_ = v1063
	var v1066 float64
	_ = v1066
	var v1069 float32
	_ = v1069
	var v1072 float64
	_ = v1072
	var v1076 float64
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1113 float64
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1196 float64
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1214 float64
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 float64
	_ = v1217
	var v1218 float64
	_ = v1218
	var v1219 float64
	_ = v1219
	var v1221 float64
	_ = v1221
	var v1224 float64
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 float64
	_ = v1243
	var v1251 float64
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	v2 = int32(0)
	v21 = float64(0)
	v36 = m.G0
	v38 = v36 - int32(256)
	m.G0 = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v46 = v38 + int32(160)
	v48 = v38 + int32(128)
	F_get_join_variables(m, v42, v43, v44, v46, v48, v38+int32(47))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v56 = v38 + int32(127)
	v57 = int32(0)
	v58 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v62 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v132 = v38 + int32(126)
	v133 = int32(0)
	v134 = float64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v133)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v138 != 0 {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+28)))
	if v100 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = *(*float32)(unsafe.Add(mBase, uint32(v65)+8))
	v68 = *(*float32)(unsafe.Add(mBase, uint32(v65)+16))
	v95 = base.F64_promote_f32(v68)
	v96 = base.F64_promote_f32(v66)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v70 == int32(16) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v95 = float64(2)
	v96 = v58
	goto L4
L9:
	;
	goto L10
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v74 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v81 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+76))
	if v77 != int32(5) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v95 = float64(-1)
	v96 = v58
	goto L4
L14:
	;
	v95 = float64(0)
	v96 = v58
	goto L4
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v84 != int32(6) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+8)))
	switch v88 - int32(_a_F_eqjoinsel_0) {
	case 0:
		goto L17
	default:
		goto L14
	case 5:
		v95 = float64(-1)
		v96 = v58
		goto L4
	}
L17:
	;
	v95 = float64(1)
	v96 = v58
	goto L4
L18:
	;
	v101 = base.F64_neg(base.F64_sub(float64(1), v96))
	goto L20
L19:
	;
	v101 = v95
	goto L20
L20:
	;
	if base.F64_gt(v101, float64(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = F_clamp_row_est(m, v101)
	mBase = m.M
	v130 = v104
	goto L3
L22:
	;
	goto L23
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v105 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v108)
	v130 = float64(200)
	goto L3
L25:
	;
	goto L26
L26:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v105)+120))
	if base.F64_le(v111, float64(0)) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v114 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v114)
	v130 = float64(200)
	goto L3
L28:
	;
	goto L29
L29:
	;
	if base.F64_lt(v101, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v121 = F_clamp_row_est(m, base.F64_mul(v111, base.F64_neg(v101)))
	mBase = m.M
	v130 = v121
	goto L3
L31:
	;
	goto L32
L32:
	;
	if base.F64_lt(v111, float64(200)) != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v124 = F_clamp_row_est(m, v111)
	mBase = m.M
	v130 = v124
	goto L3
L34:
	;
	goto L35
L35:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v125)
	v130 = float64(200)
	goto L3
L36:
	;
	v207 = F_get_opcode(m, v41)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L69
	}
L37:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+28)))
	if v176 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+22)))
	v141 = v139 + v140
	v142 = *(*float32)(unsafe.Add(mBase, uint32(v141)+8))
	v144 = *(*float32)(unsafe.Add(mBase, uint32(v141)+16))
	v171 = base.F64_promote_f32(v144)
	v172 = base.F64_promote_f32(v142)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v146 == int32(16) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v171 = float64(2)
	v172 = v134
	goto L37
L42:
	;
	goto L43
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v150 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v157 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	if v153 != int32(5) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v171 = float64(-1)
	v172 = v134
	goto L37
L47:
	;
	v171 = float64(0)
	v172 = v134
	goto L37
L48:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v160 != int32(6) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+8)))
	switch v164 - int32(_a_F_eqjoinsel_0) {
	case 0:
		goto L50
	default:
		goto L47
	case 5:
		v171 = float64(-1)
		v172 = v134
		goto L37
	}
L50:
	;
	v171 = float64(1)
	v172 = v134
	goto L37
L51:
	;
	v177 = base.F64_neg(base.F64_sub(float64(1), v172))
	goto L53
L52:
	;
	v177 = v171
	goto L53
L53:
	;
	if base.F64_gt(v177, float64(0)) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v180 = F_clamp_row_est(m, v177)
	mBase = m.M
	v206 = v180
	goto L36
L55:
	;
	goto L56
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v181 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v184)
	v206 = float64(200)
	goto L36
L58:
	;
	goto L59
L59:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v181)+120))
	if base.F64_le(v187, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v190)
	v206 = float64(200)
	goto L36
L61:
	;
	goto L62
L62:
	;
	if base.F64_lt(v177, float64(0)) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v197 = F_clamp_row_est(m, base.F64_mul(v187, base.F64_neg(v177)))
	mBase = m.M
	v206 = v197
	goto L36
L64:
	;
	goto L65
L65:
	;
	if base.F64_lt(v187, float64(200)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v200 = F_clamp_row_est(m, v187)
	mBase = m.M
	v206 = v200
	goto L36
L67:
	;
	goto L68
L68:
	;
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v201)
	v206 = float64(200)
	goto L36
L69:
	;
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+120)) = v209
	v211 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+112)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+104)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+96)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+88)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+48)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+56)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+64)) = v211
	*(*int64)(unsafe.Add(mBase, uint32(v38)+72)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v209
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v38)+168))
	if v230 == v209 {
		v253 = v209
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v38)+168))
	if v254 == int32(0) {
		v298 = v2
		v299 = v2
		goto L81
	} else {
		goto L82
	}
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	if v233 == int32(0) {
		v253 = v209
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v239 = int32(0)
	v241 = F_get_attstatsslot(m, v38+int32(88), v230, int32(1), v239, v239)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v241 == int32(0) {
		v253 = v209
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	v249 = int32(0)
	v251 = F_get_attstatsslot(m, v38+int32(48), v247, int32(1), v249, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v253 = v251
	goto L70
L76:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	switch v1114 {
	case 0, 1, 2:
		v1224 = v1113
		goto L244
	default:
		goto L242
	case 4, 5:
		goto L245
	}
L77:
	;
	if v1058 != 0 {
		goto L233
	} else {
		goto L234
	}
L78:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+126)))
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+127)))
	v1053 = int32(0)
	v1056 = v1047
	v1057 = v1048
	v1058 = v1049
	v1059 = v1051
	v1061 = v1052
	goto L77
L79:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300)+16))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+22)))
	v309 = v307 + v308
	if v253 == int32(0) {
		v1047 = v298
		v1048 = v309
		v1049 = v299
		goto L78
	} else {
		goto L103
	}
L80:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	if v301 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L81:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	if v300 != 0 {
		goto L79
	} else {
		goto L99
	}
L82:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+22)))
	v259 = v257 + v258
	if v253 == int32(0) {
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+189)))
	if v264 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v280 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L94
	}
L85:
	;
	v272 = v254
	goto L87
L86:
	;
	if v207 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v276 = F_get_attstatsslot(m, v38+int32(88), v272, int32(1), int32(0), int32(3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L93
	}
L88:
	;
	v298 = v2
	v299 = v259
	goto L81
L89:
	;
	goto L90
L90:
	;
	v267 = F_get_func_leakproof(m, v207)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v267 == int32(0) {
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v38)+168))
	v272 = v271
	goto L87
L93:
	;
	v298 = v276
	v299 = v259
	goto L81
L94:
	;
	if v280 == int32(0) {
		v298 = v2
		v299 = v259
		goto L81
	} else {
		goto L95
	}
L95:
	;
	v284 = F_get_func_name(m, v207)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v284
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_1), v38+int32(32))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_3), int32(_a_F_eqjoinsel_4))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v298 = v2
	v299 = v259
	goto L81
L99:
	;
	v1047 = v298
	v1048 = v2
	v1049 = v299
	goto L78
L100:
	;
	v1047 = v2
	v1048 = v2
	v1049 = v259
	goto L78
L101:
	;
	goto L102
L102:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+16))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+22)))
	v1047 = v2
	v1048 = v304 + v305
	v1049 = v259
	goto L78
L103:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+157)))
	if v314 != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v1027 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L1
	} else {
		goto L228
	}
L105:
	;
	v322 = v300
	goto L107
L106:
	;
	if v207 == int32(0) {
		v1047 = v298
		v1048 = v309
		v1049 = v299
		goto L78
	} else {
		goto L108
	}
L107:
	;
	v326 = F_get_attstatsslot(m, v38+int32(48), v322, int32(1), int32(0), int32(3))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L111
	}
L108:
	;
	v317 = F_get_func_leakproof(m, v207)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	if v317 == int32(0) {
		goto L104
	} else {
		goto L110
	}
L110:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	v322 = v321
	goto L107
L111:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+126)))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+127)))
	if v326&v298 != int32(1) {
		v1053 = v326
		v1056 = v298
		v1057 = v309
		v1058 = v299
		v1059 = v328
		v1061 = v329
		goto L77
	} else {
		goto L112
	}
L112:
	;
	v333 = *(*float32)(unsafe.Add(mBase, uint32(v299)+8))
	v334 = *(*float32)(unsafe.Add(mBase, uint32(v309)+8))
	v336 = v38 + int32(192)
	F_fmgr_info(m, v207, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+252)) = uint8(v339)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+244)) = uint8(v339)
	v343 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+238)) = uint16(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+236)) = uint8(v339)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+232)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v38)+224)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+220)) = v336
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v352 = F_palloc0(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v355 = F_palloc0(m, v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	if v357 <= int32(0) {
		v736 = v21
		v738 = v21
		v739 = v21
		v740 = v21
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v747 = float64(0)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v748 <= int32(0) {
		v948 = v21
		v953 = v747
		goto L171
	} else {
		goto L172
	}
L117:
	;
	v365 = v2
	v366 = v2
	v380 = v21
	goto L118
L118:
	;
	v396 = v366 << (uint(int32(2)) % 32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v396+v397)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+240)) = v399
	v401 = int32(0)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	if v402 <= v401 {
		v491 = v365
		v506 = v380
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v525 = float64(0)
	if base.F64_lt(v506, v525) != 0 {
		v533 = v525
		goto L135
	} else {
		goto L136
	}
L120:
	;
	v522 = v366 + int32(1)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	if v522 < v523 {
		v365 = v491
		v366 = v522
		v380 = v506
		goto L118
	} else {
		goto L134
	}
L121:
	;
	v405 = v401
	v409 = v402
	goto L122
L122:
	;
	v440 = v405 + v355
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	if v441 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v470 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v470)
	*(*uint8)(unsafe.Add(mBase, uint32(v366+v352))) = uint8(v470)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	v477 = *(*float32)(unsafe.Add(mBase, uint32(v475+v396)))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	v480 = *(*float32)(unsafe.Add(mBase, uint32(v478+v445)))
	v491 = v365 + v470
	v506 = base.F64_add(v380, base.F64_promote_f32(base.F32_mul(v477, v480)))
	goto L120
L124:
	;
	goto L123
L125:
	;
	v445 = v405 << (uint(int32(2)) % 32)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v445+v446)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+248)) = v448
	v450 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+236)) = uint8(v450)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v38)+220))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v456 = m.T0[v455].(func(*base.Module, int32) int32)(m, v38+int32(220))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	v466 = v409
	goto L127
L127:
	;
	v468 = v405 + int32(1)
	if v468 < v466 {
		v405 = v468
		v409 = v466
		goto L122
	} else {
		goto L133
	}
L128:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+236)))
	if v456 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v460 = v458
	goto L131
L130:
	;
	v460 = int32(1)
	goto L131
L131:
	;
	if v460 == int32(0) {
		goto L124
	} else {
		goto L132
	}
L132:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v466 = v463
	goto L127
L133:
	;
	v491 = v365
	v506 = v380
	goto L120
L134:
	;
	goto L119
L135:
	;
	v534 = base.F64_convert_i32_s(v491)
	if v523 <= int32(0) {
		v736 = v533
		v738 = v534
		v739 = v21
		v740 = v21
		goto L116
	} else {
		goto L138
	}
L136:
	;
	if base.F64_gt(v506, float64(1)) == int32(0) {
		v533 = v506
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v533 = float64(1)
	goto L135
L138:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	if v523 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v696 = float64(0)
	if base.F64_lt(v681, v696) != 0 {
		v704 = v696
		goto L166
	} else {
		goto L167
	}
L140:
	;
	v653 = *(*float32)(unsafe.Add(mBase, uint32(v537+v615<<(uint(int32(2))%32))))
	v654 = base.F64_promote_f32(v653)
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615+v352))))
	if v657 != 0 {
		goto L160
	} else {
		goto L161
	}
L141:
	;
	v615 = int32(0)
	v635 = float64(0)
	v636 = v21
	goto L140
L142:
	;
	goto L143
L143:
	;
	v546 = int32(0)
	v549 = v546
	v555 = v546
	v569 = float64(0)
	v570 = v21
	goto L144
L144:
	;
	v587 = *(*float32)(unsafe.Add(mBase, uint32(v537+v549<<(uint(int32(2))%32))))
	v588 = base.F64_promote_f32(v587)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549+v352))))
	if v591 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v523&int32(1) == int32(0) {
		v681 = v603
		v682 = v607
		goto L139
	} else {
		goto L159
	}
L146:
	;
	v592 = base.F64_add(v569, v588)
	goto L148
L147:
	;
	v592 = v569
	goto L148
L148:
	;
	v594 = v549 | int32(1)
	v598 = *(*float32)(unsafe.Add(mBase, uint32(v537+v594<<(uint(int32(2))%32))))
	v599 = base.F64_promote_f32(v598)
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594+v352))))
	if v602 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v603 = base.F64_add(v592, v599)
	goto L151
L150:
	;
	v603 = v592
	goto L151
L151:
	;
	if v591 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v605 = v570
	goto L154
L153:
	;
	v605 = base.F64_add(v570, v588)
	goto L154
L154:
	;
	if v602 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v607 = v605
	goto L157
L156:
	;
	v607 = base.F64_add(v605, v599)
	goto L157
L157:
	;
	v608 = int32(2)
	v609 = v549 + v608
	v611 = v555 + v608
	if v611 != v523&int32(2147483646) {
		v549 = v609
		v555 = v611
		v569 = v603
		v570 = v607
		goto L144
	} else {
		goto L158
	}
L158:
	;
	goto L145
L159:
	;
	v615 = v609
	v635 = v603
	v636 = v607
	goto L140
L160:
	;
	v658 = base.F64_add(v635, v654)
	goto L162
L161:
	;
	v658 = v635
	goto L162
L162:
	;
	if v657 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v660 = v636
	goto L165
L164:
	;
	v660 = base.F64_add(v636, v654)
	goto L165
L165:
	;
	v681 = v658
	v682 = v660
	goto L139
L166:
	;
	if base.F64_lt(v682, float64(0)) != 0 {
		v736 = v533
		v738 = v534
		v739 = v21
		v740 = v704
		goto L116
	} else {
		goto L169
	}
L167:
	;
	if base.F64_gt(v681, float64(1)) == int32(0) {
		v704 = v681
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v704 = float64(1)
	goto L166
L169:
	;
	v707 = float64(1)
	if base.F64_gt(v682, v707) != 0 {
		v736 = v533
		v738 = v534
		v739 = v707
		v740 = v704
		goto L116
	} else {
		goto L170
	}
L170:
	;
	v736 = v533
	v738 = v534
	v739 = v682
	v740 = v704
	goto L116
L171:
	;
	F_pfree(m, v352)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L205
	}
L172:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v748 == int32(1) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v912 = float64(0)
	if base.F64_lt(v897, v912) != 0 {
		v920 = v912
		goto L200
	} else {
		goto L201
	}
L174:
	;
	v869 = *(*float32)(unsafe.Add(mBase, uint32(v751+v831<<(uint(int32(2))%32))))
	v870 = base.F64_promote_f32(v869)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831+v355))))
	if v873 != 0 {
		goto L194
	} else {
		goto L195
	}
L175:
	;
	v754 = float64(0)
	v831 = int32(0)
	v851 = v754
	v852 = v754
	goto L174
L176:
	;
	goto L177
L177:
	;
	v761 = float64(0)
	v762 = int32(0)
	v765 = v762
	v771 = v762
	v785 = v761
	v786 = v761
	goto L178
L178:
	;
	v803 = *(*float32)(unsafe.Add(mBase, uint32(v751+v765<<(uint(int32(2))%32))))
	v804 = base.F64_promote_f32(v803)
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765+v355))))
	if v807 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v748&int32(1) == int32(0) {
		v897 = v823
		v898 = v819
		goto L173
	} else {
		goto L193
	}
L180:
	;
	v808 = v786
	goto L182
L181:
	;
	v808 = base.F64_add(v786, v804)
	goto L182
L182:
	;
	v810 = v765 | int32(1)
	v814 = *(*float32)(unsafe.Add(mBase, uint32(v751+v810<<(uint(int32(2))%32))))
	v815 = base.F64_promote_f32(v814)
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810+v355))))
	if v818 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v819 = v808
	goto L185
L184:
	;
	v819 = base.F64_add(v808, v815)
	goto L185
L185:
	;
	if v807 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v821 = base.F64_add(v785, v804)
	goto L188
L187:
	;
	v821 = v785
	goto L188
L188:
	;
	if v818 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v823 = base.F64_add(v821, v815)
	goto L191
L190:
	;
	v823 = v821
	goto L191
L191:
	;
	v824 = int32(2)
	v825 = v765 + v824
	v827 = v771 + v824
	if v827 != v748&int32(2147483646) {
		v765 = v825
		v771 = v827
		v785 = v823
		v786 = v819
		goto L178
	} else {
		goto L192
	}
L192:
	;
	goto L179
L193:
	;
	v831 = v825
	v851 = v823
	v852 = v819
	goto L174
L194:
	;
	v874 = v852
	goto L196
L195:
	;
	v874 = base.F64_add(v852, v870)
	goto L196
L196:
	;
	if v873 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v876 = base.F64_add(v851, v870)
	goto L199
L198:
	;
	v876 = v851
	goto L199
L199:
	;
	v897 = v876
	v898 = v874
	goto L173
L200:
	;
	if base.F64_lt(v898, float64(0)) != 0 {
		v948 = v920
		v953 = v747
		goto L171
	} else {
		goto L203
	}
L201:
	;
	if base.F64_gt(v897, float64(1)) == int32(0) {
		v920 = v897
		goto L200
	} else {
		goto L202
	}
L202:
	;
	v920 = float64(1)
	goto L200
L203:
	;
	v923 = float64(1)
	if base.F64_gt(v898, v923) != 0 {
		v948 = v920
		v953 = v923
		goto L171
	} else {
		goto L204
	}
L204:
	;
	v948 = v920
	v953 = v898
	goto L171
L205:
	;
	F_pfree(m, v355)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v965 = float64(1)
	v968 = base.F64_sub(base.F64_sub(base.F64_sub(v965, base.F64_promote_f32(v334)), v948), v953)
	v969 = float64(0)
	v974 = base.F64_sub(base.F64_sub(base.F64_sub(v965, base.F64_promote_f32(v333)), v740), v739)
	if base.F64_lt(v974, v969) != 0 {
		v982 = v969
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if base.F64_lt(v968, float64(0)) != 0 {
		v990 = v969
		goto L210
	} else {
		goto L211
	}
L208:
	;
	if base.F64_gt(v974, float64(1)) == int32(0) {
		v982 = v974
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v982 = float64(1)
	goto L207
L210:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v992 = base.F64_convert_i32_s(v991)
	if base.F64_lt(v992, v206) != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	if base.F64_gt(v968, float64(1)) == int32(0) {
		v990 = v968
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v990 = float64(1)
	goto L210
L213:
	;
	v998 = base.F64_add(v736, base.F64_div(base.F64_mul(v739, v990), base.F64_sub(v206, v992)))
	goto L215
L214:
	;
	v998 = v736
	goto L215
L215:
	;
	if base.F64_gt(v206, v738) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v1005 = base.F64_add(base.F64_div(base.F64_mul(v982, base.F64_add(v953, v990)), base.F64_sub(v206, v738)), v998)
	goto L218
L217:
	;
	v1005 = v998
	goto L218
L218:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v1007 = base.F64_convert_i32_s(v1006)
	if base.F64_lt(v1007, v130) != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1013 = base.F64_add(v736, base.F64_div(base.F64_mul(v953, v982), base.F64_sub(v130, v1007)))
	goto L221
L220:
	;
	v1013 = v736
	goto L221
L221:
	;
	v1014 = int32(1)
	if base.F64_gt(v130, v738) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1022 = base.F64_add(base.F64_div(base.F64_mul(base.F64_add(v739, v982), v990), base.F64_sub(v130, v738)), v1013)
	goto L224
L223:
	;
	v1022 = v1013
	goto L224
L224:
	;
	if base.F64_lt(v1005, v1022) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1024 = v1005
	goto L227
L226:
	;
	v1024 = v1022
	goto L227
L227:
	;
	v1078 = v1014
	v1082 = v1014
	v1086 = v309
	v1091 = v299
	v1094 = v328
	v1095 = v329
	v1113 = v1024
	goto L76
L228:
	;
	if v1027 == int32(0) {
		v1047 = v298
		v1048 = v309
		v1049 = v299
		goto L78
	} else {
		goto L229
	}
L229:
	;
	v1031 = F_get_func_name(m, v207)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v1031
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_1), v38+int32(16))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_3), int32(_a_F_eqjoinsel_4))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v1047 = v298
	v1048 = v309
	v1049 = v299
	goto L78
L233:
	;
	v1063 = *(*float32)(unsafe.Add(mBase, uint32(v1058)+8))
	v1066 = base.F64_promote_f32(v1063)
	goto L235
L234:
	;
	v1066 = float64(0)
	goto L235
L235:
	;
	if v1057 != 0 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1069 = *(*float32)(unsafe.Add(mBase, uint32(v1057)+8))
	v1072 = base.F64_promote_f32(v1069)
	goto L238
L237:
	;
	v1072 = float64(0)
	goto L238
L238:
	;
	if base.F64_gt(v130, v206) != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1076 = v130
	goto L241
L240:
	;
	v1076 = v206
	goto L241
L241:
	;
	v1078 = v1053
	v1082 = v1056
	v1086 = v1057
	v1091 = v1058
	v1094 = v1059
	v1095 = v1061
	v1113 = base.F64_div(base.F64_mul(base.F64_sub(float64(1), v1066), base.F64_sub(float64(1), v1072)), v1076)
	goto L76
L242:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L300
	}
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L297
	}
L244:
	;
	F_free_attstatsslot(m, v38+int32(88))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L283
	}
L245:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v1115 == int32(0) {
		goto L243
	} else {
		goto L246
	}
L246:
	;
	v1120 = int32(0)
	if v1115 == v1120 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	if v1180 == int32(0) {
		goto L243
	} else {
		goto L268
	}
L248:
	;
	if v1174 != 0 {
		goto L263
	} else {
		goto L264
	}
L249:
	;
	v1174 = int32(0)
	goto L248
L250:
	;
	goto L251
L251:
	;
	v1128 = int32(1)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+4))
	if v1129 <= v1128 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1132 = v1128
	goto L254
L253:
	;
	v1132 = v1129
	goto L254
L254:
	;
	v1137 = int32(0)
	v1139 = int32(-1)
	goto L256
L255:
	;
	v1174 = v1166
	goto L248
L256:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1115+int32(8)+v1137<<(uint(int32(2))%32))))
	if v1147 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+int32(220)))) = v1158
	v1166 = int32(1)
	goto L255
L258:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(base.I32_popcnt(v1147)))|base.B2i32(int32(0) <= v1139) != 0 {
		v1166 = v1120
		goto L255
	} else {
		goto L261
	}
L259:
	;
	v1158 = v1139
	goto L260
L260:
	;
	v1160 = v1137 + int32(1)
	if v1160 != v1132 {
		v1137 = v1160
		v1139 = v1158
		goto L256
	} else {
		goto L262
	}
L261:
	;
	v1158 = base.I32_ctz(v1147) | v1137<<(uint(int32(5))%32)
	goto L260
L262:
	;
	goto L257
L263:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v38)+220))
	v1176 = F_find_base_rel(m, v42, v1175)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L1
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1178 = F_find_join_rel(m, v42, v1115)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L267
	}
L266:
	;
	v1180 = v1176
	goto L247
L267:
	;
	v1180 = v1178
	goto L247
L268:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+47)))
	if v1183 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1218 = *(*float64)(unsafe.Add(mBase, uint32(v1180)+16))
	v1219 = base.F64_mul(v1113, v1218)
	if base.F64_gt(v1219, v1217) != 0 {
		goto L280
	} else {
		goto L281
	}
L270:
	;
	v1188 = int32(1)
	v1196 = F_eqjoinsel_semi(m, v207, v40, v38+int32(128), v130, v206, v1095&v1188, v1094&v1188, v38+int32(88), v38+int32(48), v1091, v1082, v1078, v1180)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1198 = F_get_commutator(m, v41)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L274
	}
L273:
	;
	v1217 = v1196
	goto L269
L274:
	;
	if v1198 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1200 = F_get_opcode(m, v1198)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	v1203 = int32(0)
	goto L277
L277:
	;
	v1206 = int32(1)
	v1214 = F_eqjoinsel_semi(m, v1203, v40, v38+int32(160), v206, v130, v1094&v1206, v1095&v1206, v38+int32(48), v38+int32(88), v1086, v1078, v1082, v1180)
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L279
	}
L278:
	;
	v1203 = v1200
	goto L277
L279:
	;
	v1217 = v1214
	goto L269
L280:
	;
	v1221 = v1217
	goto L282
L281:
	;
	v1221 = v1219
	goto L282
L282:
	;
	v1224 = v1221
	goto L244
L283:
	;
	F_free_attstatsslot(m, v38+int32(48))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v38)+168))
	if v1235 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v38)+172))
	m.T0[v1236].(func(*base.Module, int32))(m, v1235)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v38)+136))
	if v1239 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L287
L289:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v38)+140))
	m.T0[v1240].(func(*base.Module, int32))(m, v1239)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1243 = float64(0)
	if base.F64_lt(v1224, v1243) != 0 {
		v1251 = v1243
		goto L293
	} else {
		goto L294
	}
L292:
	;
	goto L291
L293:
	;
	v1252 = F_Float8GetDatum(m, v1251)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L296
	}
L294:
	;
	if base.F64_gt(v1224, float64(1)) == int32(0) {
		v1251 = v1224
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1251 = float64(1)
	goto L293
L296:
	;
	m.G0 = v38 + int32(256)
	return v1252
L297:
	;
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_5), int32(0))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(_a_F_eqjoinsel_6), int32(_a_F_eqjoinsel_7))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1276
	F_errmsg_internal(m, int32(_a_F_eqjoinsel_8), v38)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_eqjoinsel_2), int32(2422), int32(_a_F_eqjoinsel_9))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errcode_for_file_access(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_errcode_for_file_access[0]))
	if int32(0) <= v5 {
		v9 = v5 * int32(100)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_file_access[1])))
		switch v13 - int32(2) {
		case 0, 61, 67:
			v25 = int32(16797828)
		default:
			v25 = int32(2600)
		case 18:
			v25 = int32(33686021)
		case 27:
			v25 = int32(_a_F_errcode_for_file_access_0)
		case 29, 52, 53:
			v25 = int32(151027844)
		case 31, 39:
			v25 = int32(197)
		case 35:
			v25 = int32(50463237)
		case 42:
			v25 = int32(16908805)
		case 46:
			v25 = int32(_a_F_errcode_for_file_access_1)
		case 49:
			v25 = int32(_a_F_errcode_for_file_access_2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_errcode_for_file_access[2]))) = v25
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_errcode_for_file_access[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errcode_for_file_access_3), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errcode_for_file_access_4), int32(882), int32(_a_F_errcode_for_file_access_5))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
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
	var v6 int32
	_ = v6
	Fn13873(m, l0, l1, int32(_a_F_errhint_internal_0), int32(1346))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_errmsg(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13874(m, l0, l1, int32(_a_F_errmsg_0), int32(1077))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
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
			F_errmsg(m, int32(_a_F_errorConflictingDefElem_0), int32(0))
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
					F_errfinish(m, int32(_a_F_errorConflictingDefElem_1), int32(376), int32(_a_F_errorConflictingDefElem_2))
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
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = l0 - int32(10)
	if base.Ui32(v3) <= base.Ui32(int32(13)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_error_severity[0])))
		v10 = v8
	} else {
		v10 = int32(_a_F_error_severity_0)
	}
	return v10
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v88 int64
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v13 = l0 + l1
	v14 = l0
	v16 = l2
	v21 = v8
	goto L5
L3:
	;
	v88 = v8
	goto L4
L4:
	;
	m.G0 = v11 + int32(16)
	return v88
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v22 != int32(92) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v88 = v77
	goto L4
L7:
	;
	v77 = v21 + int64(1)
	if base.Ui32(v75) < base.Ui32(v13) {
		v14 = v75
		v16 = v16 + int32(1)
		v21 = v77
		goto L5
	} else {
		goto L18
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v22)
	v75 = v14 + int32(1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v29 = v14 + int32(3)
	if base.Ui32(v13) <= base.Ui32(v29) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = v14 + int32(1)
	if base.Ui32(v13) <= base.Ui32(v63) {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v31&int32(252) != int32(48) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)))
	if v36&int32(248) != int32(48) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v41&int32(248) != int32(48) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v55 = v41 + (v36<<(uint(int32(3))%32)&int32(56) | v31<<(uint(int32(6))%32)) - int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v55)
	v75 = v14 + int32(4)
	goto L7
L16:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v65 != int32(92) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v68 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v68)
	v75 = v14 + int32(2)
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
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_esc_decode_0)
	F_errmsg(m, int32(_a_F_esc_decode_1), v11)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_esc_decode_2), int32(513), int32(_a_F_esc_decode_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	v3 = int64(0)
	if l1 != 0 {
		v6 = l0
		v8 = v3
		for {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v13 == int32(92) {
				v16 = int64(2)
			} else {
				v16 = int64(1)
			}
			if base.I32_extend8_s(v13) <= int32(0) {
				v20 = int64(4)
			} else {
				v20 = v16
			}
			v21 = v8 + v20
			v23 = v6 + int32(1)
			if base.Ui32(v23) < base.Ui32(l0+l1) {
				v6 = v23
				v8 = v21
				continue
			} else {
				break
			}
			break
		}
		v27 = v21
	} else {
		v27 = v3
	}
	return v27
}
func F_estimate_ln_dweight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v375 int64
	_ = v375
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 float64
	_ = v396
	var v407 int64
	_ = v407
	var v422 int32
	_ = v422
	var v424 int64
	_ = v424
	var v434 int64
	_ = v434
	var v438 int64
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 float64
	_ = v447
	var v449 float64
	_ = v449
	var v462 float64
	_ = v462
	var v465 float64
	_ = v465
	var v470 float64
	_ = v470
	var v471 float64
	_ = v471
	var v472 float64
	_ = v472
	var v473 float64
	_ = v473
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v480 float64
	_ = v480
	var v505 float64
	_ = v505
	var v517 float64
	_ = v517
	var v539 float64
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v565 float64
	_ = v565
	var v567 float64
	_ = v567
	var v578 int64
	_ = v578
	var v593 int32
	_ = v593
	var v595 int64
	_ = v595
	var v605 int64
	_ = v605
	var v609 int64
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 float64
	_ = v618
	var v620 float64
	_ = v620
	var v633 float64
	_ = v633
	var v636 float64
	_ = v636
	var v641 float64
	_ = v641
	var v642 float64
	_ = v642
	var v643 float64
	_ = v643
	var v644 float64
	_ = v644
	var v649 float64
	_ = v649
	var v650 float64
	_ = v650
	var v651 float64
	_ = v651
	var v676 float64
	_ = v676
	var v688 float64
	_ = v688
	var v710 float64
	_ = v710
	var v713 int32
	_ = v713
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 != 0 {
		v713 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v713
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		v713 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = int32(1)
	v19 = int32(-1)
	v20 = int32(0)
	if base.B2i32(v19 < v16)&base.B2i32(v20 < v12) == v20 {
		v51 = v16
		v55 = v20
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if v12 <= int32(0) {
		v713 = v2
		goto L1
	} else {
		goto L118
	}
L5:
	;
	if v193 < int32(0) {
		goto L4
	} else {
		goto L48
	}
L6:
	;
	v193 = v183
	goto L5
L7:
	;
	if int32(0)|base.B2i32(v19 <= v51) != 0 {
		v87 = v19
		v89 = v20
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v32 = v16
	v36 = v20
	goto L9
L9:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v36<<(uint(int32(1))%32)))))
	if v42 != 0 {
		v183 = int32(1)
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v51 = v46
	v55 = v44
	goto L7
L11:
	;
	v43 = int32(1)
	v44 = v36 + v43
	v46 = v32 - v43
	if v46 <= v19 {
		v51 = v46
		v55 = v44
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v44 < v12 {
		v32 = v46
		v36 = v44
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if v51 != v87 {
		v129 = v55
		v130 = v89
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v68 = v19
	v70 = v20
	goto L16
L16:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70<<(uint(int32(1))%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v75 != 0 {
		v183 = int32(-1)
		goto L6
	} else {
		goto L18
	}
L17:
	;
	v87 = v79
	v89 = v77
	goto L14
L18:
	;
	v76 = int32(1)
	v77 = v70 + v76
	v79 = v68 - v76
	if v79 <= v51 {
		v87 = v79
		v89 = v77
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v77 < v18 {
		v68 = v79
		v70 = v77
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	if v12 < v129 {
		goto L30
	} else {
		goto L31
	}
L22:
	;
	v98 = v55
	v99 = v89
	goto L23
L23:
	;
	if base.B2i32(v12 <= v98)|base.B2i32(v18 <= v99) != 0 {
		v129 = v98
		v130 = v99
		goto L21
	} else {
		goto L25
	}
L24:
	;
	if base.I32_extend16_s(v115) < base.I32_extend16_s(v113) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v104 = int32(1)
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v98<<(uint(v104)%32)))))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99<<(uint(v104)%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v113 == v115 {
		v98 = v98 + v104
		v99 = v99 + v104
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v122 = int32(1)
	goto L29
L28:
	;
	v122 = int32(-1)
	goto L29
L29:
	;
	v193 = v122
	goto L5
L30:
	;
	v133 = v129
	goto L32
L31:
	;
	v133 = v12
	goto L32
L32:
	;
	v140 = v129
	goto L33
L33:
	;
	if v133 == v140 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v183 = v166
	goto L6
L35:
	;
	if v18 < v130 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v166 = int32(1)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v140<<(uint(v166)%32)))))
	if v172 == int32(0) {
		v140 = v140 + v166
		goto L33
	} else {
		goto L47
	}
L38:
	;
	v145 = v130
	goto L40
L39:
	;
	v145 = v18
	goto L40
L40:
	;
	v153 = v130
	goto L41
L41:
	;
	if v145 == v153 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v183 = int32(-1)
	goto L6
L43:
	;
	v193 = int32(0)
	goto L5
L44:
	;
	goto L45
L45:
	;
	v157 = int32(1)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32))+uint32(_c_F_estimate_ln_dweight[0]))))
	if v162 == int32(0) {
		v153 = v153 + v157
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	goto L34
L48:
	;
	v197 = int32(2)
	v198 = int32(0)
	if base.B2i32(v198 < v16)&base.B2i32(v198 < v12) == v198 {
		v230 = v16
		v234 = v198
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if int32(0) < v372 {
		goto L4
	} else {
		goto L92
	}
L50:
	;
	v372 = v362
	goto L49
L51:
	;
	if int32(0)|base.B2i32(v198 <= v230) != 0 {
		v266 = v198
		v268 = v198
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v211 = v16
	v215 = v198
	goto L53
L53:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v215<<(uint(int32(1))%32)))))
	if v221 != 0 {
		v362 = int32(1)
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v230 = v225
	v234 = v223
	goto L51
L55:
	;
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v211 - v222
	if v225 <= v198 {
		v230 = v225
		v234 = v223
		goto L51
	} else {
		goto L56
	}
L56:
	;
	if v223 < v12 {
		v211 = v225
		v215 = v223
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	if v230 != v266 {
		v308 = v234
		v309 = v268
		goto L65
	} else {
		goto L66
	}
L59:
	;
	v247 = v198
	v249 = v198
	goto L60
L60:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249<<(uint(int32(1))%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v254 != 0 {
		v362 = int32(-1)
		goto L50
	} else {
		goto L62
	}
L61:
	;
	v266 = v258
	v268 = v256
	goto L58
L62:
	;
	v255 = int32(1)
	v256 = v249 + v255
	v258 = v247 - v255
	if v258 <= v230 {
		v266 = v258
		v268 = v256
		goto L58
	} else {
		goto L63
	}
L63:
	;
	if v256 < v197 {
		v247 = v258
		v249 = v256
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	if v12 < v308 {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v277 = v234
	v278 = v268
	goto L67
L67:
	;
	if base.B2i32(v12 <= v277)|base.B2i32(v197 <= v278) != 0 {
		v308 = v277
		v309 = v278
		goto L65
	} else {
		goto L69
	}
L68:
	;
	if base.I32_extend16_s(v294) < base.I32_extend16_s(v292) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v283 = int32(1)
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v277<<(uint(v283)%32)))))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278<<(uint(v283)%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v292 == v294 {
		v277 = v277 + v283
		v278 = v278 + v283
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v301 = int32(1)
	goto L73
L72:
	;
	v301 = int32(-1)
	goto L73
L73:
	;
	v372 = v301
	goto L49
L74:
	;
	v312 = v308
	goto L76
L75:
	;
	v312 = v12
	goto L76
L76:
	;
	v319 = v308
	goto L77
L77:
	;
	if v312 == v319 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v362 = v345
	goto L50
L79:
	;
	if v197 < v309 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v345 = int32(1)
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v319<<(uint(v345)%32)))))
	if v351 == int32(0) {
		v319 = v319 + v345
		goto L77
	} else {
		goto L91
	}
L82:
	;
	v324 = v309
	goto L84
L83:
	;
	v324 = v197
	goto L84
L84:
	;
	v332 = v309
	goto L85
L85:
	;
	if v324 == v332 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v362 = int32(-1)
	goto L50
L87:
	;
	v372 = int32(0)
	goto L49
L88:
	;
	goto L89
L89:
	;
	v336 = int32(1)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v332<<(uint(v336)%32))+uint32(_c_F_estimate_ln_dweight[1]))))
	if v341 == int32(0) {
		v332 = v332 + v336
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	goto L78
L92:
	;
	v375 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v375
	F_sub_var(m, l0, int32(_a_F_estimate_ln_dweight_0), v9+int32(8))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	return int32(0)
L94:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if int32(0) < v388 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v395 = int32(*(*int16)(unsafe.Add(mBase, uint32(v394))))
	v396 = base.F64_convert_i32_s(v395)
	v407 = base.I64_reinterpret_f64(v396)
	if v407 <= int64(4503599627370495) {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	v542 = v2
	goto L97
L97:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if v543 != 0 {
		goto L114
	} else {
		goto L115
	}
L98:
	;
	v542 = v391<<(uint(int32(2))%32) + base.I32_trunc_sat_f64_s(v539)
	goto L97
L99:
	;
	v539 = v517
	goto L98
L100:
	;
	v443 = v441 + int32(_a_F_estimate_ln_dweight_1)
	v447 = base.F64_convert_i32_s(int32(base.Ui32(v443)>>(uint(int32(20))%32)) + v440)
	v449 = base.F64_mul(v447, float64(0.30102999566361177))
	v462 = base.F64_add(base.F64_reinterpret_i64(v438&int64(4294967295)|base.I64_extend_i32_u(v443&int32(_a_F_estimate_ln_dweight_2)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v465 = base.F64_mul(v462, base.F64_mul(v462, float64(0.5)))
	v470 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v462, v465)) & int64(-4294967296))
	v471 = float64(0.4342944818781689)
	v472 = base.F64_mul(v470, v471)
	v473 = base.F64_add(v449, v472)
	v478 = base.F64_div(v462, base.F64_add(v462, float64(2)))
	v479 = base.F64_mul(v478, v478)
	v480 = base.F64_mul(v479, v479)
	v505 = base.F64_add(base.F64_mul(v478, base.F64_add(v465, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v479, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, base.F64_add(base.F64_mul(v480, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v462, v470), v465))
	v517 = base.F64_add(v473, base.F64_add(base.F64_add(v472, base.F64_sub(v449, v473)), base.F64_add(base.F64_mul(v505, v471), base.F64_add(base.F64_mul(v447, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v505, v470), float64(2.5082946711645275e-11))))))
	goto L99
L101:
	;
	v434 = base.I64_reinterpret_f64(base.F64_mul(v396, float64(1.8014398509481984e+16)))
	v438 = v434
	v440 = int32(-1077)
	v441 = base.I32_wrap_i64(int64(base.Ui64(v434) >> (uint(int64(32)) % 64)))
	goto L100
L102:
	;
	if base.F64_eq(v396, float64(0)) != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v407) {
		v517 = v396
		goto L99
	} else {
		goto L109
	}
L105:
	;
	v539 = base.F64_div(float64(-1), base.F64_mul(v396, v396))
	goto L98
L106:
	;
	goto L107
L107:
	;
	if int64(0) <= v407 {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v539 = base.F64_div(base.F64_sub(v396, v396), float64(0))
	goto L98
L109:
	;
	v422 = int32(-1023)
	v424 = int64(base.Ui64(v407) >> (uint(int64(32)) % 64))
	if v424 != int64(1072693248) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v438 = v407
	v440 = v422
	v441 = base.I32_wrap_i64(v424)
	goto L100
L111:
	;
	goto L112
L112:
	;
	if base.I32_wrap_i64(v407) != 0 {
		v438 = v407
		v440 = v422
		v441 = int32(1072693248)
		goto L100
	} else {
		goto L113
	}
L113:
	;
	v539 = float64(0)
	goto L98
L114:
	;
	F_pfree(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L93
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v713 = v542
	goto L1
L117:
	;
	goto L116
L118:
	;
	v549 = v16 << (uint(int32(2)) % 32)
	v550 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15))))
	if v12 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v553 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+2)))
	v559 = v553 + v550*int32(_a_F_estimate_ln_dweight_3)
	v560 = v549 - int32(4)
	goto L121
L120:
	;
	v559 = v550
	v560 = v549
	goto L121
L121:
	;
	v565 = F_log(m, base.F64_convert_i32_s(v559))
	mBase = m.M
	v567 = base.F64_abs(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v560), float64(2.302585092994046)), v565))
	v578 = base.I64_reinterpret_f64(v567)
	if v578 <= int64(4503599627370495) {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v713 = base.I32_trunc_sat_f64_s(v710)
	goto L1
L123:
	;
	v710 = v688
	goto L122
L124:
	;
	v614 = v612 + int32(_a_F_estimate_ln_dweight_1)
	v618 = base.F64_convert_i32_s(int32(base.Ui32(v614)>>(uint(int32(20))%32)) + v611)
	v620 = base.F64_mul(v618, float64(0.30102999566361177))
	v633 = base.F64_add(base.F64_reinterpret_i64(v609&int64(4294967295)|base.I64_extend_i32_u(v614&int32(_a_F_estimate_ln_dweight_2)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v636 = base.F64_mul(v633, base.F64_mul(v633, float64(0.5)))
	v641 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v633, v636)) & int64(-4294967296))
	v642 = float64(0.4342944818781689)
	v643 = base.F64_mul(v641, v642)
	v644 = base.F64_add(v620, v643)
	v649 = base.F64_div(v633, base.F64_add(v633, float64(2)))
	v650 = base.F64_mul(v649, v649)
	v651 = base.F64_mul(v650, v650)
	v676 = base.F64_add(base.F64_mul(v649, base.F64_add(v636, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v650, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, base.F64_add(base.F64_mul(v651, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v633, v641), v636))
	v688 = base.F64_add(v644, base.F64_add(base.F64_add(v643, base.F64_sub(v620, v644)), base.F64_add(base.F64_mul(v676, v642), base.F64_add(base.F64_mul(v618, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v676, v641), float64(2.5082946711645275e-11))))))
	goto L123
L125:
	;
	v605 = base.I64_reinterpret_f64(base.F64_mul(v567, float64(1.8014398509481984e+16)))
	v609 = v605
	v611 = int32(-1077)
	v612 = base.I32_wrap_i64(int64(base.Ui64(v605) >> (uint(int64(32)) % 64)))
	goto L124
L126:
	;
	if base.F64_eq(v567, float64(0)) != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v578) {
		v688 = v567
		goto L123
	} else {
		goto L133
	}
L129:
	;
	v710 = base.F64_div(float64(-1), base.F64_mul(v567, v567))
	goto L122
L130:
	;
	goto L131
L131:
	;
	if int64(0) <= v578 {
		goto L125
	} else {
		goto L132
	}
L132:
	;
	v710 = base.F64_div(base.F64_sub(v567, v567), float64(0))
	goto L122
L133:
	;
	v593 = int32(-1023)
	v595 = int64(base.Ui64(v578) >> (uint(int64(32)) % 64))
	if v595 != int64(1072693248) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v609 = v578
	v611 = v593
	v612 = base.I32_wrap_i64(v595)
	goto L124
L135:
	;
	goto L136
L136:
	;
	if base.I32_wrap_i64(v578) != 0 {
		v609 = v578
		v611 = v593
		v612 = int32(1072693248)
		goto L124
	} else {
		goto L137
	}
L137:
	;
	v710 = float64(0)
	goto L122
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
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int64
	_ = v367
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v489 int64
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v529 int64
	_ = v529
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v678 int64
	_ = v678
	var v679 int64
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v726 int64
	_ = v726
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int64
	_ = v750
	var v751 int64
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int64
	_ = v776
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v20 != v21 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errhint(m, int32(_a_F_executeComparison_0), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L23
	} else {
		goto L316
	}
L2:
	;
	m.G0 = v17 + int32(224)
	return v842
L3:
	;
	v23 = int32(0)
	if base.B2i32(v20 == v23)|base.B2i32(v21 == v23) == v23 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v33 = int32(2)
	switch v20 {
	case 0:
		v797 = v20
		goto L9
	case 1:
		goto L21
	case 2:
		goto L22
	case 3:
		goto L18
	default:
		goto L19
	case 16, 17, 18:
		v842 = v33
		goto L2
	case 32:
		goto L20
	}
L6:
	;
	v842 = int32(2)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v842 = base.B2i32(v19 == int32(9))
	goto L2
L9:
	;
	switch v19 - int32(8) {
	case 0:
		goto L306
	case 1:
		goto L312
	case 2:
		goto L311
	case 3:
		goto L310
	case 4:
		goto L309
	case 5:
		goto L308
	default:
		goto L307
	}
L10:
	;
	v776 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	if v299 == int32(-2147483648) {
		goto L294
	} else {
		goto L295
	}
L11:
	;
	v774 = F_DirectFunctionCall2Coll(m, v770, int32(0), v772, v771)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L23
	} else {
		goto L292
	}
L12:
	;
	if v297 <= int32(1183) {
		goto L262
	} else {
		goto L263
	}
L13:
	;
	if v296&int32(1) != 0 {
		goto L252
	} else {
		goto L253
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L23
	} else {
		goto L249
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L23
	} else {
		goto L246
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L23
	} else {
		goto L243
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L23
	} else {
		goto L240
	}
L18:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v605 != 0 {
		goto L234
	} else {
		goto L235
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L23
	} else {
		goto L231
	}
L20:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+35)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v300 <= int32(1183) {
		goto L129
	} else {
		goto L130
	}
L21:
	;
	if v19 == int32(8) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v38 = F_DirectFunctionCall2Coll(m, int32(1327), int32(0), v36, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v797 = v38
	goto L9
L25:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v44 != v45 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[0]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	goto L50
L28:
	;
	v842 = int32(0)
	goto L2
L29:
	;
	goto L30
L30:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v44) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v842 = base.B2i32(v111 == int32(0))
	goto L2
L32:
	;
	v111 = int32(0)
	goto L31
L33:
	;
	v85 = v80
	v86 = v81
	v87 = v82
	goto L43
L34:
	;
	if (v48|v49)&int32(3) != 0 {
		v80 = v48
		v81 = v49
		v82 = v44
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v73 = v48
	v74 = v49
	v75 = v44
	goto L36
L36:
	;
	if v75 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v57 = v48
	v58 = v49
	v59 = v44
	goto L38
L38:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v62 != v63 {
		v80 = v57
		v81 = v58
		v82 = v59
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v73 = v68
	v74 = v66
	v75 = v70
	goto L36
L40:
	;
	v65 = int32(4)
	v66 = v58 + v65
	v68 = v57 + v65
	v70 = v59 - v65
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v57 = v68
		v58 = v66
		v59 = v70
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v80 = v73
	v81 = v74
	v82 = v75
	goto L33
L43:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v90 == v91 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v111 = v90 - v91
	goto L31
L45:
	;
	v93 = int32(1)
	v98 = v87 - v93
	if v98 != 0 {
		v85 = v85 + v93
		v86 = v86 + v93
		v87 = v98
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	v229 = base.B2i32(v116 < v114)
	if v116 < v114 {
		goto L96
	} else {
		goto L97
	}
L50:
	;
	if v120 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[0]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	goto L52
L52:
	;
	if v125 == int32(6) {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v129 = F_pg_server_to_any(m, v117, v116, int32(6))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v132 = F_pg_server_to_any(m, v115, v114, int32(6))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v134 = base.B2i32(v129 == v117)
	if v134 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v137 = F_strlen(m, v129)
	mBase = m.M
	v138 = v137
	goto L58
L57:
	;
	v138 = v116
	goto L58
L58:
	;
	v139 = base.B2i32(v132 == v115)
	if v139 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v142 = F_strlen(m, v132)
	mBase = m.M
	v143 = v142
	goto L61
L60:
	;
	v143 = v114
	goto L61
L61:
	;
	v144 = base.B2i32(v138 < v143)
	if v138 < v143 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v145 = v138
	goto L64
L63:
	;
	v145 = v143
	goto L64
L64:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v145) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v207 != 0 {
		goto L83
	} else {
		goto L84
	}
L66:
	;
	v207 = int32(0)
	goto L65
L67:
	;
	v181 = v176
	v182 = v177
	v183 = v178
	goto L77
L68:
	;
	if (v129|v132)&int32(3) != 0 {
		v176 = v129
		v177 = v132
		v178 = v145
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v169 = v129
	v170 = v132
	v171 = v145
	goto L70
L70:
	;
	if v171 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v153 = v129
	v154 = v132
	v155 = v145
	goto L72
L72:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v158 != v159 {
		v176 = v153
		v177 = v154
		v178 = v155
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v169 = v164
	v170 = v162
	v171 = v166
	goto L70
L74:
	;
	v161 = int32(4)
	v162 = v154 + v161
	v164 = v153 + v161
	v166 = v155 - v161
	if base.Ui32(int32(3)) < base.Ui32(v166) {
		v153 = v164
		v154 = v162
		v155 = v166
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v176 = v169
	v177 = v170
	v178 = v171
	goto L67
L77:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 == v187 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v207 = v186 - v187
	goto L65
L79:
	;
	v189 = int32(1)
	v194 = v183 - v189
	if v194 != 0 {
		v181 = v181 + v189
		v182 = v182 + v189
		v183 = v194
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	v210 = v207
	goto L85
L84:
	;
	v210 = base.B2i32(v143 < v138) - v144
	goto L85
L85:
	;
	if v139&base.B2i32(v129 == v117) != 0 {
		v797 = v210
		goto L9
	} else {
		goto L86
	}
L86:
	;
	if v134 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_pfree(m, v129)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L23
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	if v139 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v132)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L23
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v210 != 0 {
		v797 = v210
		goto L9
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	goto L49
L96:
	;
	v230 = v116
	goto L98
L97:
	;
	v230 = v114
	goto L98
L98:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v230) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if v292 != 0 {
		goto L117
	} else {
		goto L118
	}
L100:
	;
	v292 = int32(0)
	goto L99
L101:
	;
	v266 = v261
	v267 = v262
	v268 = v263
	goto L111
L102:
	;
	if (v117|v115)&int32(3) != 0 {
		v261 = v117
		v262 = v115
		v263 = v230
		goto L101
	} else {
		goto L105
	}
L103:
	;
	v254 = v117
	v255 = v115
	v256 = v230
	goto L104
L104:
	;
	if v256 == int32(0) {
		goto L100
	} else {
		goto L110
	}
L105:
	;
	v238 = v117
	v239 = v115
	v240 = v230
	goto L106
L106:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v243 != v244 {
		v261 = v238
		v262 = v239
		v263 = v240
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v254 = v249
	v255 = v247
	v256 = v251
	goto L104
L108:
	;
	v246 = int32(4)
	v247 = v239 + v246
	v249 = v238 + v246
	v251 = v240 - v246
	if base.Ui32(int32(3)) < base.Ui32(v251) {
		v238 = v249
		v239 = v247
		v240 = v251
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v261 = v254
	v262 = v255
	v263 = v256
	goto L101
L111:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v271 == v272 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v292 = v271 - v272
	goto L99
L113:
	;
	v274 = int32(1)
	v279 = v268 - v274
	if v279 != 0 {
		v266 = v266 + v274
		v267 = v267 + v274
		v268 = v279
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	goto L100
L117:
	;
	v295 = v292
	goto L119
L118:
	;
	v295 = base.B2i32(v114 < v116) - v229
	goto L119
L119:
	;
	v797 = v295
	goto L9
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L23
	} else {
		goto L228
	}
L121:
	;
	if v300 == int32(1114) {
		goto L12
	} else {
		goto L227
	}
L122:
	;
	if v297 <= int32(1183) {
		goto L194
	} else {
		goto L195
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L23
	} else {
		goto L188
	}
L124:
	;
	if v297 == int32(1114) {
		v842 = v33
		goto L2
	} else {
		goto L187
	}
L125:
	;
	if v296&int32(1) == int32(0) {
		goto L15
	} else {
		goto L185
	}
L126:
	;
	if v297 == int32(1184) {
		v842 = v33
		goto L2
	} else {
		goto L183
	}
L127:
	;
	if v297 <= int32(1183) {
		goto L172
	} else {
		goto L173
	}
L128:
	;
	if v297 <= int32(1183) {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	switch v300 - int32(1082) {
	case 0:
		goto L128
	case 1:
		goto L127
	default:
		goto L121
	}
L130:
	;
	goto L131
L131:
	;
	if v300 == int32(1184) {
		goto L122
	} else {
		goto L132
	}
L132:
	;
	if v300 != int32(1266) {
		goto L120
	} else {
		goto L133
	}
L133:
	;
	v309 = int32(1427)
	if int32(1183) < v297 {
		goto L126
	} else {
		goto L134
	}
L134:
	;
	switch v297 - int32(1082) {
	case 0:
		v842 = v33
		goto L2
	case 1:
		goto L125
	default:
		goto L124
	}
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L23
	} else {
		goto L167
	}
L136:
	;
	if v297 == int32(1114) {
		goto L10
	} else {
		goto L166
	}
L137:
	;
	switch v297 - int32(1082) {
	case 0:
		v770 = int32(1428)
		v771 = v298
		v772 = v299
		goto L11
	case 1:
		v842 = v33
		goto L2
	default:
		goto L136
	}
L138:
	;
	goto L139
L139:
	;
	if v297 != int32(1184) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v297 != int32(1266) {
		goto L135
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	if v296&int32(1) == int32(0) {
		goto L17
	} else {
		goto L144
	}
L143:
	;
	v842 = v33
	goto L2
L144:
	;
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	v330 = m.G0
	v332 = v330 - int32(48)
	m.G0 = v332
	if v299 == int32(-2147483648) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v797 = v388
	goto L9
L146:
	;
	m.G0 = v332 + int32(48)
	goto L145
L147:
	;
	v337 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v327)
	mBase = m.M
	v388 = v337
	goto L146
L148:
	;
	goto L149
L149:
	;
	if v299 == int32(2147483647) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v386 = F_timestamp_cmp_internal(m, v385, v327)
	mBase = m.M
	v388 = v386
	goto L146
L151:
	;
	v385 = int64(9223372036854775807)
	goto L150
L152:
	;
	goto L153
L153:
	;
	if v299 <= int32(106751982) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v327 == int64(-9223372036854775807-1) {
		goto L163
	} else {
		goto L164
	}
L155:
	;
	F_j2date(m, v299+int32(_a_F_executeComparison_1), v332+int32(24), v332+int32(20), v332+int32(16))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v332)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+12)) = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[1]))
	v360 = F_DetermineTimeZoneOffset(m, v332+int32(4), v359)
	mBase = m.M
	v367 = base.I64_extend_i32_s(v360)*int64(1000000) + base.I64_extend_i32_s(v299)*int64(86400000000)
	if base.Ui64(v367+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v385 = v367
		goto L150
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	if v327 == int64(9223372036854775807) {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	if v367 < int64(-211813488000000000) {
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v379 = int32(-1)
	goto L162
L161:
	;
	v379 = int32(1)
	goto L162
L162:
	;
	v388 = v379
	goto L146
L163:
	;
	v384 = int32(1)
	goto L165
L164:
	;
	v384 = int32(-1)
	goto L165
L165:
	;
	v388 = v384
	goto L146
L166:
	;
	goto L135
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v297
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(48))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L23
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3756), int32(_a_F_executeComparison_4))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L23
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L23
	} else {
		goto L180
	}
L171:
	;
	if v297 == int32(1114) {
		v842 = v33
		goto L2
	} else {
		goto L179
	}
L172:
	;
	switch v297 - int32(1082) {
	case 0:
		v842 = v33
		goto L2
	case 1:
		v770 = int32(1429)
		v771 = v298
		v772 = v299
		goto L11
	default:
		goto L171
	}
L173:
	;
	goto L174
L174:
	;
	if v297 == int32(1184) {
		v842 = v33
		goto L2
	} else {
		goto L175
	}
L175:
	;
	if v297 != int32(1266) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	if v296&int32(1) == int32(0) {
		goto L16
	} else {
		goto L177
	}
L177:
	;
	v426 = F_DirectFunctionCall1Coll(m, int32(1416), int32(0), v299)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L23
	} else {
		goto L178
	}
L178:
	;
	v770 = int32(1427)
	v771 = v298
	v772 = v426
	goto L11
L179:
	;
	goto L170
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v297
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(80))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L23
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3782), int32(_a_F_executeComparison_4))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L23
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	if v297 != int32(1266) {
		goto L123
	} else {
		goto L184
	}
L184:
	;
	v770 = v309
	v771 = v298
	v772 = v299
	goto L11
L185:
	;
	v456 = F_DirectFunctionCall1Coll(m, int32(1416), int32(0), v298)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L23
	} else {
		goto L186
	}
L186:
	;
	v770 = v309
	v771 = v456
	v772 = v299
	goto L11
L187:
	;
	goto L123
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v297
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(112))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L23
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3808), int32(_a_F_executeComparison_4))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L23
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L23
	} else {
		goto L224
	}
L192:
	;
	if v297 == int32(1114) {
		goto L13
	} else {
		goto L223
	}
L193:
	;
	if v296&int32(1) == int32(0) {
		goto L14
	} else {
		goto L201
	}
L194:
	;
	switch v297 - int32(1082) {
	case 0:
		goto L193
	case 1:
		v842 = v33
		goto L2
	default:
		goto L192
	}
L195:
	;
	goto L196
L196:
	;
	if v297 == int32(1184) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v770 = int32(1430)
	v771 = v298
	v772 = v299
	goto L11
L198:
	;
	goto L199
L199:
	;
	if v297 != int32(1266) {
		goto L191
	} else {
		goto L200
	}
L200:
	;
	v842 = v33
	goto L2
L201:
	;
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	v492 = m.G0
	v494 = v492 - int32(48)
	m.G0 = v494
	if v298 == int32(-2147483648) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v797 = int32(0) - v550
	goto L9
L203:
	;
	m.G0 = v494 + int32(48)
	goto L202
L204:
	;
	v499 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v489)
	mBase = m.M
	v550 = v499
	goto L203
L205:
	;
	goto L206
L206:
	;
	if v298 == int32(2147483647) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v548 = F_timestamp_cmp_internal(m, v547, v489)
	mBase = m.M
	v550 = v548
	goto L203
L208:
	;
	v547 = int64(9223372036854775807)
	goto L207
L209:
	;
	goto L210
L210:
	;
	if v298 <= int32(106751982) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if v489 == int64(-9223372036854775807-1) {
		goto L220
	} else {
		goto L221
	}
L212:
	;
	F_j2date(m, v298+int32(_a_F_executeComparison_1), v494+int32(24), v494+int32(20), v494+int32(16))
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v494)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v494)+12)) = int32(0)
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_executeComparison[1]))
	v522 = F_DetermineTimeZoneOffset(m, v494+int32(4), v521)
	mBase = m.M
	v529 = base.I64_extend_i32_s(v522)*int64(1000000) + base.I64_extend_i32_s(v298)*int64(86400000000)
	if base.Ui64(v529+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v547 = v529
		goto L207
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	if v489 == int64(9223372036854775807) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	if v529 < int64(-211813488000000000) {
		goto L211
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v541 = int32(-1)
	goto L219
L218:
	;
	v541 = int32(1)
	goto L219
L219:
	;
	v550 = v541
	goto L203
L220:
	;
	v546 = int32(1)
	goto L222
L221:
	;
	v546 = int32(-1)
	goto L222
L222:
	;
	v550 = v546
	goto L203
L223:
	;
	goto L191
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v297
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(176))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L23
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3866), int32(_a_F_executeComparison_4))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L23
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	goto L120
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v300
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(32))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L23
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3871), int32(_a_F_executeComparison_4))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L23
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v593
	F_errmsg_internal(m, int32(_a_F_executeComparison_5), v17)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L23
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3403), int32(_a_F_executeComparison_6))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L23
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	v606 = int32(1)
	goto L236
L235:
	;
	v606 = int32(-1)
	goto L236
L236:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v605 != v608 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v610 = v606
	goto L239
L238:
	;
	v610 = int32(0)
	goto L239
L239:
	;
	v797 = v610
	goto L9
L240:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L23
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(_a_F_executeComparison_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = int32(_a_F_executeComparison_8)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17-int32(-64))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L23
	} else {
		goto L242
	}
L242:
	;
	goto L1
L243:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L23
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(_a_F_executeComparison_11)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17+int32(96))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L23
	} else {
		goto L245
	}
L245:
	;
	goto L1
L246:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L23
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+132)) = int32(_a_F_executeComparison_10)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = int32(_a_F_executeComparison_11)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17+int32(128))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L23
	} else {
		goto L248
	}
L248:
	;
	goto L1
L249:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L23
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = int32(_a_F_executeComparison_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = int32(_a_F_executeComparison_8)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17+int32(192))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L23
	} else {
		goto L251
	}
L251:
	;
	goto L1
L252:
	;
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	v680 = F_timestamp_cmp_timestamptz_internal(m, v678, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L23
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L23
	} else {
		goto L256
	}
L255:
	;
	v797 = int32(0) - v680
	goto L9
L256:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L23
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = int32(_a_F_executeComparison_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = int32(_a_F_executeComparison_12)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17+int32(208))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L23
	} else {
		goto L258
	}
L258:
	;
	goto L1
L259:
	;
	if v296&int32(1) != 0 {
		goto L285
	} else {
		goto L286
	}
L260:
	;
	v726 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	if v298 == int32(-2147483648) {
		goto L273
	} else {
		goto L274
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L23
	} else {
		goto L269
	}
L262:
	;
	switch v297 - int32(1082) {
	case 0:
		goto L260
	case 1:
		v842 = v33
		goto L2
	default:
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	if v297 == int32(1184) {
		goto L259
	} else {
		goto L267
	}
L265:
	;
	if v297 != int32(1114) {
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v770 = int32(1430)
	v771 = v298
	v772 = v299
	goto L11
L267:
	;
	if v297 == int32(1266) {
		v842 = v33
		goto L2
	} else {
		goto L268
	}
L268:
	;
	goto L261
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v297
	F_errmsg_internal(m, int32(_a_F_executeComparison_2), v17+int32(144))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L23
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3837), int32(_a_F_executeComparison_4))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L23
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	v797 = int32(0) - v746
	goto L9
L273:
	;
	v730 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v726)
	mBase = m.M
	v746 = v730
	goto L272
L274:
	;
	goto L275
L275:
	;
	if v298 == int32(2147483647) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v734 = F_timestamp_cmp_internal(m, int64(9223372036854775807), v726)
	mBase = m.M
	v746 = v734
	goto L272
L277:
	;
	goto L278
L278:
	;
	if v298 <= int32(106751982) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v740 = F_timestamp_cmp_internal(m, base.I64_extend_i32_s(v298)*int64(86400000000), v726)
	mBase = m.M
	v746 = v740
	goto L272
L280:
	;
	goto L281
L281:
	;
	if v726 == int64(9223372036854775807) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v745 = int32(-1)
	goto L284
L283:
	;
	v745 = int32(1)
	goto L284
L284:
	;
	v746 = v745
	goto L272
L285:
	;
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v299)))
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	v752 = F_timestamp_cmp_timestamptz_internal(m, v750, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L23
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L23
	} else {
		goto L289
	}
L288:
	;
	v797 = v752
	goto L9
L289:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L23
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+164)) = int32(_a_F_executeComparison_7)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = int32(_a_F_executeComparison_12)
	F_errmsg(m, int32(_a_F_executeComparison_9), v17+int32(160))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L23
	} else {
		goto L291
	}
L291:
	;
	goto L1
L292:
	;
	v797 = v774
	goto L9
L293:
	;
	v797 = v796
	goto L9
L294:
	;
	v780 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v776)
	mBase = m.M
	v796 = v780
	goto L293
L295:
	;
	goto L296
L296:
	;
	if v299 == int32(2147483647) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v784 = F_timestamp_cmp_internal(m, int64(9223372036854775807), v776)
	mBase = m.M
	v796 = v784
	goto L293
L298:
	;
	goto L299
L299:
	;
	if v299 <= int32(106751982) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v790 = F_timestamp_cmp_internal(m, base.I64_extend_i32_s(v299)*int64(86400000000), v776)
	mBase = m.M
	v796 = v790
	goto L293
L301:
	;
	goto L302
L302:
	;
	if v776 == int64(9223372036854775807) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v795 = int32(-1)
	goto L305
L304:
	;
	v795 = int32(1)
	goto L305
L305:
	;
	v796 = v795
	goto L293
L306:
	;
	v842 = base.B2i32(v797 == int32(0))
	goto L2
L307:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L23
	} else {
		goto L313
	}
L308:
	;
	v842 = base.B2i32(int32(0) <= v797)
	goto L2
L309:
	;
	v842 = base.B2i32(v797 <= int32(0))
	goto L2
L310:
	;
	v842 = base.B2i32(int32(0) < v797)
	goto L2
L311:
	;
	v842 = int32(base.Ui32(v797) >> (uint(int32(31)) % 32))
	goto L2
L312:
	;
	v842 = base.B2i32(v797 != int32(0))
	goto L2
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v19
	F_errmsg_internal(m, int32(_a_F_executeComparison_13), v17+int32(16))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L23
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3427), int32(_a_F_executeComparison_6))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L23
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errfinish(m, int32(_a_F_executeComparison_3), int32(3672), int32(_a_F_executeComparison_14))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L23
	} else {
		goto L317
	}
L317:
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
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
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
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if int32(0) < v24 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = v6
	v39 = v6
	goto L10
L8:
	;
	v221 = v6
	goto L9
L9:
	;
	return v221
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v45 = v42 + v36<<(uint(int32(5))%32)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+30)))
	if v46 != 0 {
		v202 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v221 = v202
	goto L9
L12:
	;
	v206 = v36 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v206 < v207 {
		v36 = v206
		v39 = v202
		goto L10
	} else {
		goto L38
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v36<<(uint(int32(2))%32))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 == int32(0) {
		v202 = v39
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+4)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	v61 = F_makeVar(m, v56, v57, v58, v59, v60, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v67
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = l3
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+40)) = uint16(v69)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
	if v73 == int32(0) {
		v151 = l0
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v72 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v77 = v73 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v73) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v88 = l0
	v89 = int32(0)
	goto L22
L20:
	;
	v116 = l0
	goto L21
L21:
	;
	v132 = v116
	v133 = int32(0)
	goto L26
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v107 = v89 + int32(8)
	if v107 != v73&int32(-8) {
		v88 = v105
		v89 = v107
		goto L22
	} else {
		goto L24
	}
L23:
	;
	if v77 == int32(0) {
		v151 = v105
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v116 = v105
	goto L21
L26:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v144 = v133 + int32(1)
	if v144 != v77 {
		v132 = v142
		v133 = v144
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v151 = v142
	goto L17
L28:
	;
	goto L27
L29:
	;
	v182 = F_lappend(m, v39, v61)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L35
	}
L30:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v151)+20))
	if v163 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 < v72 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168+v72<<(uint(int32(2))%32)-int32(4))))
	if v174 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
	v178 = F_bms_union(m, v177, v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v178
	goto L29
L35:
	;
	if l4 == int32(0) {
		v202 = v182
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v187 = F_lappend(m, v186, v51)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v187
	v202 = v182
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == v4 {
		v96 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v96
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+18)))
	if v14 != int32(1) {
		v96 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_palloc0(m, int32(136))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(101)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v28 = F_makeAlias(m, v24+int32(4), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v28
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v36 == v30 {
		v96 = l0
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+18)))
	if v39 != int32(1) {
		v96 = l0
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if int32(0) < v42 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = int32(0)
	v50 = v42
	v53 = v4
	goto L12
L10:
	;
	v89 = v4
	goto L11
L11:
	;
	v91 = int32(0)
	v94 = F_ReplaceVarsFromTargetList(m, l0, l2, v18, v89, v91, int32(1), l2, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L22
	}
L12:
	;
	v58 = v49 + int32(1)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49*int32(100)+(v35+v50<<(uint(int32(4))%32)))+110)))
	if v63 == int32(118) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v89 = v80
	goto L11
L14:
	;
	v66 = F_build_generation_expression(m, l1, v58)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v79 = v50
	v80 = v53
	goto L16
L16:
	;
	if v58 < v79 {
		v49 = v58
		v50 = v79
		v53 = v80
		goto L12
	} else {
		goto L21
	}
L17:
	;
	F_ChangeVarNodes(m, v66, int32(1), l2)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(0)
	v74 = F_makeTargetEntry(m, v66, base.I32_extend16_s(v58), v72, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v76 = F_lappend(m, v53, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v79 = v78
	v80 = v76
	goto L16
L21:
	;
	goto L13
L22:
	;
	v96 = v94
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
	var v26 int32
	_ = v26
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	v26 = v2
	v29 = v17
	goto L4
L4:
	;
	v37 = v26 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v26<<(uint(int32(2))%32))))
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
		v26 = v37
		v29 = v188
		goto L4
	} else {
		goto L40
	}
L9:
	;
	F_relation_close(m, v48, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L39
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v59 < v61 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = v59
	v67 = v59
	v68 = v61
	goto L17
L15:
	;
	v126 = v59
	goto L16
L16:
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
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v81 = v52 + v68<<(uint(int32(4))%32) + v66*int32(100)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+110)))
	if v82 == int32(118) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v126 = v120
	goto L16
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v116 < v121 {
		v66 = v116
		v67 = v120
		v68 = v121
		goto L17
	} else {
		goto L30
	}
L20:
	;
	v86 = v66 + int32(1)
	v87 = F_build_generation_expression(m, v48, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v100 = v66 + int32(1)
	v101 = base.I32_extend16_s(v100)
	v103 = v81 + int32(20)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+68))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+96))
	v108 = F_makeVar(m, v37, v101, v104, v105, v106, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L27
	}
L23:
	;
	F_ChangeVarNodes(m, v87, int32(1), v37)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v93 = int32(0)
	v95 = F_makeTargetEntry(m, v87, base.I32_extend16_s(v86), v93, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v97 = F_lappend(m, v67, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v116 = v86
	v120 = v97
	goto L19
L27:
	;
	v110 = int32(0)
	v112 = F_makeTargetEntry(m, v108, v101, v110, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v114 = F_lappend(m, v67, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v116 = v100
	v120 = v114
	goto L19
L30:
	;
	goto L18
L31:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v153 = v147<<(uint(int32(2))%32) + int32(4)
	goto L33
L32:
	;
	v153 = int32(4)
	goto L33
L33:
	;
	v154 = F_palloc0(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v29)+108))
	if v157 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(1)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v166 = F_replace_rte_variables(m, v29, v160, int32(0), int32(851), v15+int32(8), v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v173 = v166
	goto L9
L39:
	;
	v188 = v173
	goto L8
L40:
	;
	goto L5
}
