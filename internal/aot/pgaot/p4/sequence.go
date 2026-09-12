package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SequenceChangePersistence(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	F_LockRelationOid(m, l0, int32(8))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_init_sequence(m, l0, v6+int32(28), v6+int32(24))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)))
			if v19 != int32(112) {
				v32 = F_read_seq_tuple(m, v17, v6+int32(20), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_RelationSetNewRelfilenumber(m, v17, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_fill_seq_with_data(m, v17, v6)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
							F_UnlockReleaseBuffer(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_sequence_close(m, v17, int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									m.G0 = v6 + int32(32)
									return
								}
							}
						}
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, _consts[27]))
				if v23 <= int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
					if v26 != 0 {
						v32 = F_read_seq_tuple(m, v17, v6+int32(20), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_RelationSetNewRelfilenumber(m, v17, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_fill_seq_with_data(m, v17, v6)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
									F_UnlockReleaseBuffer(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_sequence_close(m, v17, int32(0))
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											m.G0 = v6 + int32(32)
											return
										}
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
						if v27 != 0 {
							v32 = F_read_seq_tuple(m, v17, v6+int32(20), v6)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_RelationSetNewRelfilenumber(m, v17, l1)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									F_fill_seq_with_data(m, v17, v6)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
										F_UnlockReleaseBuffer(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_sequence_close(m, v17, int32(0))
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return
											} else {
												m.G0 = v6 + int32(32)
												return
											}
										}
									}
								}
							}
						} else {
							v28 = F_GetTopTransactionId(m)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								v32 = F_read_seq_tuple(m, v17, v6+int32(20), v6)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_RelationSetNewRelfilenumber(m, v17, l1)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return
									} else {
										F_fill_seq_with_data(m, v17, v6)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
											F_UnlockReleaseBuffer(m, v38)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return
											} else {
												F_sequence_close(m, v17, int32(0))
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
													return
												} else {
													m.G0 = v6 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v28 = F_GetTopTransactionId(m)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v32 = F_read_seq_tuple(m, v17, v6+int32(20), v6)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_RelationSetNewRelfilenumber(m, v17, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								F_fill_seq_with_data(m, v17, v6)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
									F_UnlockReleaseBuffer(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_sequence_close(m, v17, int32(0))
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											m.G0 = v6 + int32(32)
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
func F_has_sequence_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[168]))
			v21 = F_convert_any_priv_string(m, v16, int32(1693520))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_textToQualifiedNameList(m, v11)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = F_makeRangeVarFromNameList(m, v23)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = int32(0)
						v31 = F_RangeVarGetRelidExtended(m, v25, v27, v27, v27, v27)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_get_rel_relkind(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								if v33 != int32(83) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											v44 = F_text_to_cstring(m, v11)
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v44
												F_errmsg(m, int32(437432), v8)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523802), int32(2153), int32(399966))
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
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
								} else {
									v55 = F_pg_class_aclcheck(m, v31, v19, v21)
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return base.B2i32(v55 == int32(0))
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
func F_sequenceIsOwned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v15 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v11, int32(1), int32(3), int32(184), int32(1259))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v11+int32(48), int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_systable_beginscan(m, v15, int32(2673), int32(1), int32(0), int32(2), v11)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v38 = F_systable_getnext(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v38 == int32(0) {
		v73 = v5
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v44 = v38
	goto L9
L9:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v55 != int32(1259) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v73 = v5
	goto L5
L11:
	;
	v65 = F_systable_getnext(m, v36)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+24)))
	if v58 != l1&int32(255) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v62
	v73 = int32(1)
	goto L5
L14:
	;
	if v65 != 0 {
		v44 = v65
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	F_sequence_close(m, v15, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v11 + int32(96)
	return v73
}
