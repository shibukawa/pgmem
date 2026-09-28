package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonTableDestroyOpaque(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v3 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableDestroyOpaque_0))
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v3))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v5
		return
	}
}
func F_JsonTableGetValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	v9 = F_GetJsonTableExecContext(m, l0, int32(_a_F_JsonTableGetValue_0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
		v15 = l1 << (uint(int32(2)) % 32)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15)))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+112)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v21)
			return int64(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26+v15)))
			if v28 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)) = uint8(v32)
				v34 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v31
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
				v37 = m.T0[v36].(func(*base.Module, int32, int32, int32) int64)(m, v28, v29, l4)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)) = uint8(v30)
					*(*int64)(unsafe.Add(mBase, uint32(v29)+40)) = v34
					return v37
				}
			} else {
				v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+120)))
				v43 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v43)
				return v42
			}
		}
	}
}
func F_appendJSONKeyValueFmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0]))
	v17 = F_palloc(m, int32(128))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0])) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
	v23 = F_pvsnprintf(m, v17, int32(128), l3, l4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = v17
	v33 = v23
	goto L7
L5:
	;
	v51 = v17
	goto L6
L6:
	;
	if v51 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	F_pfree(m, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v51 = v38
	goto L6
L9:
	;
	v38 = F_palloc(m, v33)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_appendJSONKeyValueFmt[0])) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l4
	v43 = F_pvsnprintf(m, v38, v33, l3, l4)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(v33) <= base.Ui32(v43) {
		v32 = v38
		v33 = v43
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	F_pfree(m, v51)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_escape_json(m, l0, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_appendStringInfoChar(m, l0, int32(58))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_escape_json(m, l0, v51)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_appendStringInfoString(m, l0, v51)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L13
L22:
	;
	goto L13
L23:
	;
	m.G0 = v12 + int32(16)
	return
}
func F_json_agg_transfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_json_agg_transfn_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_json_categorize_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v28 int32
	_ = v28
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_getBaseType(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
		if v10 <= int32(1081) {
			switch v10 - int32(16) {
			case 0:
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1243)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1)
				m.G0 = v8 + int32(16)
				return
			case 1, 2, 3, 6:
				v74 = F_get_element_type(m, v10)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					if v74 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
						m.G0 = v8 + int32(16)
						return
					} else {
						switch v10 - int32(2277) {
						case 0, 10:
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
							m.G0 = v8 + int32(16)
							return
						case 1, 2, 3, 4, 5, 6, 7, 8, 9:
							v84 = F_type_is_rowtype(m, v10)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								if v84 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
									m.G0 = v8 + int32(16)
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
									if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
										v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											if v98 != int32(1) {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v112 = m.ExcPending
												if v112 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
												if v102 == int32(0) {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									} else {
										F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return
										} else {
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						default:
							if v10 != int32(_a_F_json_categorize_type_1) {
								v84 = F_type_is_rowtype(m, v10)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									if v84 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
										m.G0 = v8 + int32(16)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
										if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
											v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return
											} else {
												if v98 != int32(1) {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v112 = m.ExcPending
													if v112 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
													if v102 == int32(0) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
														*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										} else {
											F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												m.G0 = v8 + int32(16)
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			case 4, 5, 7:
				F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
					m.G0 = v8 + int32(16)
					return
				}
			default:
				if base.Ui32(v10-int32(700)) < base.Ui32(int32(2)) {
					F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					if v10 != int32(114) {
						v74 = F_get_element_type(m, v10)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							if v74 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							} else {
								switch v10 - int32(2277) {
								case 0, 10:
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								case 1, 2, 3, 4, 5, 6, 7, 8, 9:
									v84 = F_type_is_rowtype(m, v10)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										if v84 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
											m.G0 = v8 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
											if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
												v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													if v98 != int32(1) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
														if v102 == int32(0) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											} else {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								default:
									if v10 != int32(_a_F_json_categorize_type_1) {
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						F_getTypeOutputInfo(m, int32(114), l3, v8+int32(15))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(6)
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			if v10 <= int32(1183) {
				if v10 == int32(1082) {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1085)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
					m.G0 = v8 + int32(16)
					return
				} else {
					if v10 != int32(1114) {
						v74 = F_get_element_type(m, v10)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							if v74 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
								m.G0 = v8 + int32(16)
								return
							} else {
								switch v10 - int32(2277) {
								case 0, 10:
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								case 1, 2, 3, 4, 5, 6, 7, 8, 9:
									v84 = F_type_is_rowtype(m, v10)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										if v84 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
											m.G0 = v8 + int32(16)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
											if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
												v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return
												} else {
													if v98 != int32(1) {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
														if v102 == int32(0) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											} else {
												F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return
												} else {
													m.G0 = v8 + int32(16)
													return
												}
											}
										}
									}
								default:
									if v10 != int32(_a_F_json_categorize_type_1) {
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1313)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(4)
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				if v10 == int32(1184) {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1151)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(5)
					m.G0 = v8 + int32(16)
					return
				} else {
					if v10 == int32(1700) {
						F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(2)
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						if v10 != int32(3802) {
							v74 = F_get_element_type(m, v10)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								if v74 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
									m.G0 = v8 + int32(16)
									return
								} else {
									switch v10 - int32(2277) {
									case 0, 10:
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
										m.G0 = v8 + int32(16)
										return
									case 1, 2, 3, 4, 5, 6, 7, 8, 9:
										v84 = F_type_is_rowtype(m, v10)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											if v84 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
												m.G0 = v8 + int32(16)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
												if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
													v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
													mBase = m.M
													v99 = m.ExcPending
													if v99 != 0 {
														return
													} else {
														if v98 != int32(1) {
															F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v8 + int32(16)
																return
															}
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
															if v102 == int32(0) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																m.G0 = v8 + int32(16)
																return
															}
														}
													}
												} else {
													F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														m.G0 = v8 + int32(16)
														return
													}
												}
											}
										}
									default:
										if v10 != int32(_a_F_json_categorize_type_1) {
											v84 = F_type_is_rowtype(m, v10)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												if v84 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(2291)
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(9)
													m.G0 = v8 + int32(16)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(11)
													if base.Ui32(int32(_a_F_json_categorize_type_0)) <= base.Ui32(v10) {
														v98 = F_find_coercion_pathway(m, int32(114), v10, int32(3), v8+int32(8))
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return
														} else {
															if v98 != int32(1) {
																F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																mBase = m.M
																v112 = m.ExcPending
																if v112 != 0 {
																	return
																} else {
																	m.G0 = v8 + int32(16)
																	return
																}
															} else {
																v102 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
																if v102 == int32(0) {
																	F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
																	mBase = m.M
																	v112 = m.ExcPending
																	if v112 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(16)
																		return
																	}
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v102
																	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(10)
																	m.G0 = v8 + int32(16)
																	return
																}
															}
														}
													} else {
														F_getTypeOutputInfo(m, v10, l3, v8+int32(15))
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															m.G0 = v8 + int32(16)
															return
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(751)
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(8)
											m.G0 = v8 + int32(16)
											return
										}
									}
								}
							}
						} else {
							F_getTypeOutputInfo(m, int32(3802), l3, v8+int32(15))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if l1 != 0 {
									v54 = int32(7)
								} else {
									v54 = int32(6)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v54
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_json_manifest_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v10 - int32(3) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = l1
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = F_strtox_2(m, l1, v6+int32(-4), int32(10), int64(-9223372036854775807-1))
		mBase = m.M
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
		if v21 != 0 {
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_json_manifest_scalar_0)
			m.T0[v86].(func(*base.Module, int32, int32, int32))(m, v85, int32(_a_F_json_manifest_scalar_1), v6+int32(-32))
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v22 = base.I32_wrap_i64(v19)
			if base.Ui32(v22-int32(3)) <= base.Ui32(int32(-3)) {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_json_manifest_scalar_2)
				m.T0[v95].(func(*base.Module, int32, int32, int32))(m, v94, int32(_a_F_json_manifest_scalar_1), v6+int32(-48))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				m.T0[v27].(func(*base.Module, int32, int32))(m, v14, v22)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
					m.G0 = v8 - int32(-64)
					return int32(0)
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = l1
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v40 = F_strtox_2(m, l1, v6+int32(-4), int32(10), int64(-1))
		mBase = m.M
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
		if v42 != 0 {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a_F_json_manifest_scalar_3)
			m.T0[v104].(func(*base.Module, int32, int32, int32))(m, v103, int32(_a_F_json_manifest_scalar_1), v6+int32(-16))
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
			m.T0[v43].(func(*base.Module, int32, int64))(m, v35, v40)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
				m.G0 = v8 - int32(-64)
				return int32(0)
			}
		}
	default:
		v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_json_manifest_scalar_4)
		m.T0[v68].(func(*base.Module, int32, int32, int32))(m, v67, int32(_a_F_json_manifest_scalar_1), v8)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	case 5:
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		switch v48 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 2:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 3:
			F_pfree(m, l1)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
				m.G0 = v8 - int32(-64)
				return int32(0)
			}
		case 4:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		case 5:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		default:
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(7)
			m.G0 = v8 - int32(-64)
			return int32(0)
		}
	case 9:
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if base.Ui32(v58) <= base.Ui32(int32(2)) {
			*(*int32)(unsafe.Add(mBase, uint32(l0+v58<<(uint(int32(2))%32))+40)) = l1
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(11)
		m.G0 = v8 - int32(-64)
		return int32(0)
	case 10:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
		m.G0 = v8 - int32(-64)
		return int32(0)
	}
}
func F_json_object_agg_unique_transfn(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_json_object_agg_transfn_worker(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_json_parse_manifest_incremental_chunk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
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
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v397 int32
	_ = v397
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v740 int64
	_ = v740
	var v741 int64
	_ = v741
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v747 int64
	_ = v747
	var v750 int64
	_ = v750
	var v752 int64
	_ = v752
	var v753 int64
	_ = v753
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	v4 = l3
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l0 == int32(_a_F_json_parse_manifest_incremental_chunk_0) {
		v429 = int32(16)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v429 == v4^int32(1) {
		goto L122
	} else {
		goto L123
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v29 == int32(_a_F_json_parse_manifest_incremental_chunk_1) {
		v429 = int32(16)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v33 != int32(1) {
		v429 = int32(2)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = l0 + int32(68)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)) = uint8(v4)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
	v47 = F_json_lex(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v429 = v397
	goto L1
L6:
	;
	return
L7:
	;
	if v47 != 0 {
		v397 = v47
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v50 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v397 = int32(0)
	goto L5
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v54 = int32(_a_F_json_parse_manifest_incremental_chunk_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v53))) = uint16(v54)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v58 = v56 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v58
	if v58 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v62 = v50
	goto L12
L12:
	;
	v67 = v62
	v71 = v49
	goto L14
L13:
	;
	v62 = v58
	goto L12
L14:
	;
	v82 = v67 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v85 = v84 + v82
	v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85))))
	if v86 == v71 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L9
L16:
	;
	if v369 != 0 {
		v67 = v369
		v71 = v371
		goto L14
	} else {
		goto L121
	}
L17:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v369 = v368
	v371 = v365
	goto L16
L18:
	;
	if base.Ui32(int32(11)) < base.Ui32(v71) {
		v365 = v71
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v86&int32(32) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v90 = F_json_lex(m, l0)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	if v90 != 0 {
		v397 = v90
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v365 = v92
	goto L17
L24:
	;
	v117 = v86 & int32(255)
	if v86&int32(64) != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v101 = v86*int32(104) + v71<<(uint(int32(3))%32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_json_parse_manifest_incremental_chunk[0])))
	if v104 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_json_parse_manifest_incremental_chunk[1])))
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	base.MemoryCopy(m, v85, v104, v109)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v112 = v111 + v109
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v112
	v369 = v112
	v371 = v71
	goto L16
L30:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v339&int32(4) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L31:
	;
	switch v117 + int32(-64) {
	case 0:
		goto L44
	case 1:
		goto L43
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	case 5:
		goto L39
	case 6:
		goto L38
	case 7:
		goto L37
	case 8:
		goto L36
	case 9:
		goto L35
	case 10:
		goto L34
	default:
		v365 = v71
		goto L17
	}
L32:
	;
	goto L33
L33:
	;
	switch v117 - int32(1) {
	case 0:
		goto L111
	default:
		goto L104
	case 3, 35:
		goto L110
	case 5, 33:
		v327 = int32(3)
		goto L103
	case 6:
		goto L109
	case 7:
		goto L108
	case 11:
		goto L107
	case 32:
		goto L106
	case 34:
		goto L105
	}
L34:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	if v282 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L96
	}
L35:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v37)+36))
	v247 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v247
	if v246 == v247 {
		v365 = v71
		goto L17
	} else {
		goto L83
	}
L36:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	if v233 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L80
	}
L37:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v224 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v220+v221))) = uint8(v224)
	if v218 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L77
	}
L38:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v200 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L74
	}
L39:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v184 = base.B2i32(v71 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v180+v181))) = uint8(v184)
	if v178 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L71
	}
L40:
	;
	v164 = int32(0)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v165|v166 == v164 {
		v338 = v164
		goto L30
	} else {
		goto L67
	}
L41:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L63
	}
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(_a_F_json_parse_manifest_incremental_chunk_3) < v144 {
		v429 = int32(3)
		goto L1
	} else {
		goto L56
	}
L43:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	F_dec_lex_level(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L52
	}
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if int32(_a_F_json_parse_manifest_incremental_chunk_3) < v123 {
		v429 = int32(3)
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v126 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v128 = m.T0[v126].(func(*base.Module, int32) int32)(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	if v128 != 0 {
		v397 = v128
		goto L5
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v365 = v71
	goto L17
L52:
	;
	if v133 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v139 = m.T0[v133].(func(*base.Module, int32) int32)(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v139 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v397 = v139
	goto L5
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v147 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v149 = m.T0[v147].(func(*base.Module, int32) int32)(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_inc_lex_level(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	if v149 != 0 {
		v397 = v149
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v365 = v71
	goto L17
L63:
	;
	if v154 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L64
	}
L64:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v160 = m.T0[v154].(func(*base.Module, int32) int32)(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	if v160 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L66
	}
L66:
	;
	v397 = v160
	goto L5
L67:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v170 != int32(1) {
		v338 = v164
		goto L30
	} else {
		goto L68
	}
L68:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = F_pstrdup(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	if v175 != 0 {
		v338 = v175
		goto L30
	} else {
		goto L70
	}
L70:
	;
	v429 = int32(16)
	goto L1
L71:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190+v191<<(uint(int32(2))%32))))
	v196 = m.T0[v178].(func(*base.Module, int32, int32, int32) int32)(m, v188, v195, v184)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	if v196 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L73
	}
L73:
	;
	v397 = v196
	goto L5
L74:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211+v206))))
	v214 = m.T0[v200].(func(*base.Module, int32, int32, int32) int32)(m, v203, v210, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	if v214 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L76
	}
L76:
	;
	v397 = v214
	goto L5
L77:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v229 = m.T0[v218].(func(*base.Module, int32, int32) int32)(m, v228, v224)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v229 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L79
	}
L79:
	;
	v397 = v229
	goto L5
L80:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238+v239))))
	v242 = m.T0[v233].(func(*base.Module, int32, int32) int32)(m, v236, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	if v242 == int32(0) {
		v365 = v71
		goto L17
	} else {
		goto L82
	}
L82:
	;
	v397 = v242
	goto L5
L83:
	;
	if v71 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v71
	v365 = v71
	goto L17
L85:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v253 != int32(1) {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v264 = v262 - v263
	v267 = F_palloc(m, v264+int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L91
	}
L88:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v258 = F_pstrdup(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v258
	if v258 != 0 {
		goto L84
	} else {
		goto L90
	}
L90:
	;
	v429 = int32(16)
	goto L1
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v267
	if v267 == int32(0) {
		v429 = int32(16)
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v264 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v267, v273, v264)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v275+v264))) = uint8(v277)
	goto L84
L96:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v288 = m.T0[v282].(func(*base.Module, int32, int32, int32) int32)(m, v285, v286, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v290&int32(4) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v301
	if v288 == v301 {
		v365 = v71
		goto L17
	} else {
		goto L102
	}
L99:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v295 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v397 = v288
	goto L5
L103:
	;
	v328 = int32(11)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v329 == int32(0) {
		v397 = v328
		goto L5
	} else {
		goto L115
	}
L104:
	;
	v327 = int32(0)
	goto L103
L105:
	;
	v327 = int32(4)
	goto L103
L106:
	;
	v327 = int32(2)
	goto L103
L107:
	;
	v327 = int32(8)
	goto L103
L108:
	;
	v327 = int32(5)
	goto L103
L109:
	;
	v316 = int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-v316))))
	if v318 == v316 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v327 = int32(6)
	goto L103
L111:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85-int32(1)))))
	v327 = base.B2i32(v310 == int32(8))
	goto L103
L112:
	;
	v321 = int32(6)
	goto L114
L113:
	;
	v321 = int32(3)
	goto L114
L114:
	;
	v327 = v321
	goto L103
L115:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v332 == int32(12) {
		v397 = v328
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v327<<(uint(int32(2))%32))+uint32(_c_F_json_parse_manifest_incremental_chunk[2])))
	v429 = v337
	goto L1
L117:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v357+v358<<(uint(int32(2))%32)))) = v338
	v365 = v71
	goto L17
L118:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v345+v346<<(uint(int32(2))%32))))
	if v350 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	F_pfree(m, v350)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	goto L15
L122:
	;
	if v4 != 0 {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	goto L124
L124:
	;
	v822 = F_json_errdetail(m, v429, l0)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L6
	} else {
		goto L218
	}
L125:
	;
	m.G0 = v21 + int32(32)
	return
L126:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v452 = int32(0)
	v453 = m.G0
	v455 = v453 - int32(112)
	m.G0 = v455
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if l2 == v452 {
		goto L136
	} else {
		goto L137
	}
L127:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v433 == int32(14) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v443 = F_pg_cryptohash_update(m, v442, l1, l2)
	mBase = m.M
	if int32(0) <= v443 {
		goto L125
	} else {
		goto L132
	}
L130:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_json_parse_manifest_incremental_chunk_4)
	m.T0[v436].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_5), v21)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	m.T0[v448].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_6), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	goto L125
L134:
	;
	goto L125
L135:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+32)) = int32(_a_F_json_parse_manifest_incremental_chunk_7)
	m.T0[v793].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_5), v455+int32(32))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L6
	} else {
		goto L217
	}
L136:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455))) = int32(_a_F_json_parse_manifest_incremental_chunk_8)
	m.T0[v787].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_5), v455)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L6
	} else {
		goto L216
	}
L137:
	;
	if l2 != int32(1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if base.Ui32(v548) <= base.Ui32(int32(1)) {
		goto L136
	} else {
		goto L164
	}
L139:
	;
	v471 = v452
	v475 = v5
	v476 = v5
	v477 = v5
	v482 = v5
	goto L142
L140:
	;
	v516 = v5
	v517 = v5
	v518 = v5
	v523 = v5
	goto L141
L141:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v517))))
	v528 = base.B2i32(v526 == int32(10))
	if v526 == int32(10) {
		goto L158
	} else {
		goto L159
	}
L142:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v476))))
	v487 = base.B2i32(v485 == int32(10))
	if v485 == int32(10) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if l2&int32(1) == int32(0) {
		v541 = v497
		v543 = v496
		v548 = v499
		goto L138
	} else {
		goto L157
	}
L144:
	;
	v488 = v476
	goto L146
L145:
	;
	v488 = v475
	goto L146
L146:
	;
	if v485 == int32(10) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v489 = v475
	goto L149
L148:
	;
	v489 = v477
	goto L149
L149:
	;
	v491 = v476 | int32(1)
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v491))))
	v495 = base.B2i32(v493 == int32(10))
	if v493 == int32(10) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v496 = v488
	goto L152
L151:
	;
	v496 = v489
	goto L152
L152:
	;
	if v493 == int32(10) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v497 = v491
	goto L155
L154:
	;
	v497 = v488
	goto L155
L155:
	;
	v499 = v487 + v482 + v495
	v500 = int32(2)
	v501 = v476 + v500
	v503 = v471 + v500
	if v503 != l2&int32(-2) {
		v471 = v503
		v475 = v497
		v476 = v501
		v477 = v496
		v482 = v499
		goto L142
	} else {
		goto L156
	}
L156:
	;
	goto L143
L157:
	;
	v516 = v497
	v517 = v501
	v518 = v496
	v523 = v499
	goto L141
L158:
	;
	v529 = v516
	goto L160
L159:
	;
	v529 = v518
	goto L160
L160:
	;
	if v526 == int32(10) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v530 = v517
	goto L163
L162:
	;
	v530 = v516
	goto L163
L163:
	;
	v541 = v530
	v543 = v529
	v548 = v528 + v523
	goto L138
L164:
	;
	if v541 != l2-int32(1) {
		goto L135
	} else {
		goto L165
	}
L165:
	;
	if v451 != 0 {
		v575 = v451
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v576 = F_pg_cryptohash_update(m, v575, l1, v543+int32(1))
	mBase = m.M
	if v576 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L167:
	;
	v558 = F_pg_cryptohash_create(m, int32(3))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v558 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v564].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_9), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v567 = F_pg_cryptohash_init(m, v558)
	mBase = m.M
	if int32(0) <= v567 {
		v575 = v558
		goto L166
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v572].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_10), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L6
	} else {
		goto L174
	}
L174:
	;
	v575 = v558
	goto L166
L175:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v581].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_6), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L6
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v587 = F_pg_cryptohash_final(m, v575, v455+int32(80), int32(32))
	mBase = m.M
	if v587 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v592].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_11), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L6
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	if v595 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	goto L181
L183:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v601].(func(*base.Module, int32, int32, int32))(m, v598, int32(_a_F_json_parse_manifest_incremental_chunk_12), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L6
	} else {
		goto L186
	}
L184:
	;
	v605 = v595
	goto L185
L185:
	;
	v606 = F_strlen(m, v605)
	mBase = m.M
	if v606 != int32(64) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v605 = v604
	goto L185
L187:
	;
	v740 = *(*int64)(unsafe.Add(mBase, uint32(v455)+80))
	v741 = *(*int64)(unsafe.Add(mBase, uint32(v455)+48))
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v455)+88))
	v744 = *(*int64)(unsafe.Add(mBase, uint32(v455)+56))
	v747 = *(*int64)(unsafe.Add(mBase, uint32(v455)+96))
	v750 = *(*int64)(unsafe.Add(mBase, uint32(v455-int32(-64))))
	v752 = *(*int64)(unsafe.Add(mBase, uint32(v455)+104))
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v455)+72))
	if v740^v741|(v743^v744)|(v747^v750|(v752^v753)) != int64(0) {
		goto L211
	} else {
		goto L212
	}
L188:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v455)+16)) = v605
	m.T0[v715].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_13), v455+int32(16))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L210
	}
L189:
	;
	v620 = int32(0)
	goto L190
L190:
	;
	v630 = v605 + v620<<(uint(int32(1))%32)
	v631 = int32(*(*int8)(unsafe.Add(mBase, uint32(v630))))
	v633 = v631 - int32(48)
	if base.Ui32(v633&int32(255)) <= base.Ui32(int32(9)) {
		v656 = v633
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L187
L192:
	;
	v657 = int32(*(*int8)(unsafe.Add(mBase, uint32(v630)+1)))
	v659 = v657 - int32(48)
	if base.Ui32(v659&int32(255)) <= base.Ui32(int32(9)) {
		v682 = v659
		goto L200
	} else {
		goto L201
	}
L193:
	;
	if base.Ui32((v631-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v656 = v631 - int32(87)
	goto L192
L195:
	;
	goto L196
L196:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v631-int32(65))&int32(255)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v655 = int32(-1)
	goto L199
L198:
	;
	v655 = v631 - int32(55)
	goto L199
L199:
	;
	v656 = v655
	goto L192
L200:
	;
	if v682|v656 < int32(0) {
		goto L188
	} else {
		goto L208
	}
L201:
	;
	if base.Ui32((v657-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v682 = v657 - int32(87)
	goto L200
L203:
	;
	goto L204
L204:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v657-int32(65))&int32(255)) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v681 = int32(-1)
	goto L207
L206:
	;
	v681 = v657 - int32(55)
	goto L207
L207:
	;
	v682 = v681
	goto L200
L208:
	;
	v691 = v682 + v656<<(uint(int32(4))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v455+int32(48)+v620))) = uint8(v691)
	v694 = v620 + int32(1)
	if v694 != int32(32) {
		v620 = v694
		goto L190
	} else {
		goto L209
	}
L209:
	;
	goto L191
L210:
	;
	goto L187
L211:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v457)+20))
	m.T0[v761].(func(*base.Module, int32, int32, int32))(m, v457, int32(_a_F_json_parse_manifest_incremental_chunk_14), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	F_pg_cryptohash_free(m, v575)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L6
	} else {
		goto L215
	}
L214:
	;
	goto L213
L215:
	;
	m.G0 = v455 + int32(112)
	goto L134
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v822
	m.T0[v824].(func(*base.Module, int32, int32, int32))(m, v24, int32(_a_F_json_parse_manifest_incremental_chunk_5), v21+int32(16))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_populate_record(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(1)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_json_populate_record_0), v3, v3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_json_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v9 = v6 + int32(8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_pq_getmsgtext(m, v10, v11-v12, v6+int32(76))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_json_recv[0]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
		v25 = F_makeJsonLexContextCstringLen(m, v9, v16, v20, v23, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v29 = F_pg_parse_json_or_errsave(m, v9, int32(_a_F_json_recv_0), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+76))
				v32 = F_cstring_to_text_with_len(m, v16, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					m.G0 = v6 + int32(80)
					return base.I64_extend_i32_u(v32)
				}
			}
		}
	}
}
func F_makeJsonLexContextCstringLen(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v5 = l4
	if l0 == int32(0) {
		v9 = F_palloc0(m, int32(68))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 == int32(0) {
				return int32(_a_F_makeJsonLexContextCstringLen_0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v17 | int32(1)
				v24 = v9
				*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+56)) = uint8(v5)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l2
				if v5 != 0 {
					v35 = F_makeStringInfo(m)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v35
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v38 | int32(2)
						return v24
					}
				} else {
					return v24
				}
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(68))
		v24 = l0
		*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
		*(*uint8)(unsafe.Add(mBase, uint32(v24)+56)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l2
		if v5 != 0 {
			v35 = F_makeStringInfo(m)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v35
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v38 | int32(2)
				return v24
			}
		} else {
			return v24
		}
	}
}
func F_printJsonPathItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_printJsonPathItem[0]))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v18 {
	case 0:
		goto L8
	case 1:
		goto L55
	case 2:
		goto L54
	case 3:
		goto L53
	case 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 41:
		goto L52
	case 6:
		goto L51
	case 7:
		goto L50
	case 19, 20:
		goto L49
	case 21:
		goto L48
	case 22:
		goto L47
	case 23:
		goto L46
	case 24:
		goto L45
	case 25:
		goto L44
	case 26:
		goto L43
	case 27:
		goto L42
	case 28:
		goto L41
	case 29:
		goto L40
	case 30:
		goto L39
	case 31:
		goto L38
	case 32:
		goto L37
	case 33:
		goto L36
	case 34:
		goto L35
	case 35:
		goto L34
	case 36:
		goto L33
	case 37:
		goto L32
	case 38:
		goto L31
	default:
		goto L9
	case 40:
		goto L30
	case 42:
		goto L29
	case 43:
		goto L28
	case 44:
		goto L27
	case 45:
		goto L26
	case 46:
		goto L25
	case 47:
		goto L24
	case 48:
		goto L23
	case 49:
		goto L22
	case 50:
		goto L21
	case 51:
		goto L20
	case 52:
		goto L19
	case 53:
		goto L18
	case 54:
		goto L17
	case 55:
		goto L16
	case 56:
		goto L15
	case 57:
		goto L13
	case 58:
		goto L12
	case 59:
		goto L11
	case 60:
		goto L10
	case 61:
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v806 {
		goto L342
	} else {
		goto L343
	}
L8:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L341
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L338
	}
L10:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_1))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L337
	}
L11:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_2))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L330
	}
L12:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_3))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L323
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_4))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L316
	}
L14:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_5))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L309
	}
L15:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_6))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L308
	}
L16:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_7))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L1
	} else {
		goto L307
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_8))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L300
	}
L18:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_9))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L293
	}
L19:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_10))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L286
	}
L20:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_11))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L279
	}
L21:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_12))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L272
	}
L22:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_13))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L271
	}
L23:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_14))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L270
	}
L24:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_15))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L269
	}
L25:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_16))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L256
	}
L26:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_17))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L255
	}
L27:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_18))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L254
	}
L28:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_19))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L253
	}
L29:
	;
	if l3 != 0 {
		goto L210
	} else {
		goto L211
	}
L30:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_20))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L209
	}
L31:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_21))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L208
	}
L32:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_22))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L201
	}
L33:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_23))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L200
	}
L34:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_24))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L199
	}
L35:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_25))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L198
	}
L36:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_26))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L197
	}
L37:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_27))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L196
	}
L38:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_28))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L195
	}
L39:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_29))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L191
	}
L40:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_30))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L187
	}
L41:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L185
	}
L42:
	;
	F_appendStringInfoChar(m, l0, int32(36))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L184
	}
L43:
	;
	F_appendStringInfoChar(m, l0, int32(64))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L183
	}
L44:
	;
	if l2 != 0 {
		goto L178
	} else {
		goto L179
	}
L45:
	;
	if l2 != 0 {
		goto L155
	} else {
		goto L156
	}
L46:
	;
	F_appendStringInfoChar(m, l0, int32(91))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L128
	}
L47:
	;
	if l2 != 0 {
		goto L123
	} else {
		goto L124
	}
L48:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_31))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L122
	}
L49:
	;
	if l3 != 0 {
		goto L104
	} else {
		goto L105
	}
L50:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L100
	}
L51:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_32))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L96
	}
L52:
	;
	if l3 != 0 {
		goto L70
	} else {
		goto L71
	}
L53:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v44 != 0 {
		goto L65
	} else {
		goto L66
	}
L54:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v23 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L7
L57:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v31 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+12)))
	v32 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	F_appendStringInfoString(m, l0, base.I32_wrap_i64(v32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 <= int32(0) {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L7
L65:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_33))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_34))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L7
L69:
	;
	goto L7
L70:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v56, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v63 = v61 - int32(4)
	if base.Ui32(v63) <= base.Ui32(int32(37)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v69 = v68
	goto L77
L76:
	;
	v69 = int32(6)
	goto L77
L77:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v75 = v73 - int32(4)
	if base.Ui32(v75) <= base.Ui32(int32(37)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v82 = v80
	goto L80
L79:
	;
	v82 = int32(6)
	goto L80
L80:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v69) <= base.Ui32(v82)))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v90 = F_jspOperationName(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_appendStringInfoString(m, l0, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_appendStringInfoChar(m, l0, int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v10+int32(100), v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v106 = v104 - int32(4)
	if base.Ui32(v106) <= base.Ui32(int32(37)) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v112 = v111
	goto L89
L88:
	;
	v112 = int32(6)
	goto L89
L89:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v118 = v116 - int32(4)
	if base.Ui32(v118) <= base.Ui32(int32(37)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v125 = v123
	goto L92
L91:
	;
	v125 = int32(6)
	goto L92
L92:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v112) <= base.Ui32(v125)))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L7
L96:
	;
	v138 = v10 + int32(100)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v138, v139, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v143 = int32(0)
	F_printJsonPathItem(m, l0, v138, v143, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L7
L100:
	;
	v154 = v10 + int32(100)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v154, v155, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v159 = int32(0)
	F_printJsonPathItem(m, l0, v154, v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_35))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L7
L104:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v172 = v18
	goto L106
L106:
	;
	if v172 == int32(19) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v172 = v171
	goto L106
L108:
	;
	v175 = int32(43)
	goto L110
L109:
	;
	v175 = int32(45)
	goto L110
L110:
	;
	F_appendStringInfoChar(m, l0, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v10+int32(100), v180, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v187 = v185 - int32(4)
	if base.Ui32(v187) <= base.Ui32(int32(37)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v193 = v192
	goto L115
L114:
	;
	v193 = int32(6)
	goto L115
L115:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v199 = v197 - int32(4)
	if base.Ui32(v199) <= base.Ui32(int32(37)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v206 = v204
	goto L118
L117:
	;
	v206 = int32(6)
	goto L118
L118:
	;
	F_printJsonPathItem(m, l0, v10+int32(100), int32(0), base.B2i32(base.Ui32(v193) <= base.Ui32(v206)))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	goto L7
L122:
	;
	goto L7
L123:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_appendStringInfoChar(m, l0, int32(42))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	goto L7
L128:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v227 <= int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L154
	}
L130:
	;
	v231 = v10 + int32(100)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	F_jspInitByBuffer(m, v231, v232, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v238 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v240 = v10 + int32(72)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v240, v241, v238)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v254 = v10 + int32(100)
	goto L134
L134:
	;
	v255 = int32(0)
	F_printJsonPathItem(m, l0, v254, v255, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v244 = int32(0)
	F_printJsonPathItem(m, l0, v231, v244, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_36))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v254 = v240
	goto L134
L138:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v259 < int32(2) {
		goto L129
	} else {
		goto L139
	}
L139:
	;
	v266 = int32(1)
	goto L140
L140:
	;
	v271 = v10 + int32(100)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v274 = v266 << (uint(int32(3)) % 32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v274+v275)))
	F_jspInitByBuffer(m, v271, v272, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L129
L142:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280+v274)+4))
	if v282 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v305 = int32(0)
	F_printJsonPathItem(m, l0, v304, v305, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L152
	}
L144:
	;
	v284 = v10 + int32(72)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v284, v285, v282)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L151
	}
L147:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v291 = int32(0)
	F_printJsonPathItem(m, l0, v271, v291, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_36))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v304 = v284
	goto L143
L151:
	;
	v304 = v10 + int32(100)
	goto L143
L152:
	;
	v310 = v266 + int32(1)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v310 < v311 {
		v266 = v310
		goto L140
	} else {
		goto L153
	}
L153:
	;
	goto L141
L154:
	;
	goto L7
L155:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v327 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L158:
	;
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v327
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_37), v10+int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L177
	}
L160:
	;
	if v327 == int32(-1) {
		goto L171
	} else {
		goto L172
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v327
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_38), v10+int32(32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L170
	}
L162:
	;
	switch v326 + int32(1) {
	case 0:
		goto L165
	case 1:
		goto L161
	default:
		goto L159
	}
L163:
	;
	goto L164
L164:
	;
	if v327 != v326 {
		goto L160
	} else {
		goto L167
	}
L165:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_39))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	goto L7
L167:
	;
	if v327 != int32(-1) {
		goto L161
	} else {
		goto L168
	}
L168:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_40))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	goto L7
L170:
	;
	goto L7
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v326
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_41), v10+int32(48))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v326 != int32(-1) {
		goto L159
	} else {
		goto L175
	}
L174:
	;
	goto L7
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v327
	F_appendStringInfo(m, l0, int32(_a_F_printJsonPathItem_42), v10-int32(-64))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	goto L7
L177:
	;
	goto L7
L178:
	;
	F_appendStringInfoChar(m, l0, int32(46))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v373, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L182
	}
L181:
	;
	goto L180
L182:
	;
	goto L7
L183:
	;
	goto L7
L184:
	;
	goto L7
L185:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_escape_json_with_len(m, l0, v386, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L7
L187:
	;
	v394 = v10 + int32(100)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v394, v395, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v399 = int32(0)
	F_printJsonPathItem(m, l0, v394, v399, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L7
L191:
	;
	v410 = v10 + int32(100)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v410, v411, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v415 = int32(0)
	F_printJsonPathItem(m, l0, v410, v415, v415)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	goto L7
L195:
	;
	goto L7
L196:
	;
	goto L7
L197:
	;
	goto L7
L198:
	;
	goto L7
L199:
	;
	goto L7
L200:
	;
	goto L7
L201:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v443 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v445 = v10 + int32(100)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v445, v446, v443)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	v449 = int32(0)
	F_printJsonPathItem(m, l0, v445, v449, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	goto L7
L208:
	;
	goto L7
L209:
	;
	goto L7
L210:
	;
	F_appendStringInfoChar(m, l0, int32(40))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v467 = v10 + int32(100)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v467, v468, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v475 = v473 - int32(4)
	if base.Ui32(v475) <= base.Ui32(int32(37)) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v485 = v483 - int32(4)
	if base.Ui32(v485) <= base.Ui32(int32(37)) {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v475<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v482 = v480
	goto L218
L217:
	;
	v482 = int32(6)
	goto L218
L218:
	;
	goto L215
L219:
	;
	F_printJsonPathItem(m, l0, v467, int32(0), base.B2i32(base.Ui32(v482) <= base.Ui32(v492)))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L223
	}
L220:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v485<<(uint(int32(2))%32))+uint32(_c_F_printJsonPathItem[1])))
	v492 = v490
	goto L222
L221:
	;
	v492 = int32(6)
	goto L222
L222:
	;
	goto L219
L223:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_43))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_escape_json_with_len(m, l0, v499, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v503 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_printJsonPathItem_44))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L251
	}
L229:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v507&int32(1) != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	F_appendStringInfoChar(m, l0, int32(105))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	v514 = v507
	goto L232
L232:
	;
	if v514&int32(2) != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v514 = v513
	goto L232
L234:
	;
	F_appendStringInfoChar(m, l0, int32(115))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	v521 = v514
	goto L236
L236:
	;
	if v521&int32(4) != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v521 = v520
	goto L236
L238:
	;
	F_appendStringInfoChar(m, l0, int32(109))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	v528 = v521
	goto L240
L240:
	;
	if v528&int32(8) != 0 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v528 = v527
	goto L240
L242:
	;
	F_appendStringInfoChar(m, l0, int32(120))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L245
	}
L243:
	;
	v535 = v528
	goto L244
L244:
	;
	if v535&int32(16) != 0 {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v535 = v534
	goto L244
L246:
	;
	F_appendStringInfoChar(m, l0, int32(113))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L250
	}
L249:
	;
	goto L248
L250:
	;
	goto L228
L251:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	goto L7
L253:
	;
	goto L7
L254:
	;
	goto L7
L255:
	;
	goto L7
L256:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v562 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v564 = v10 + int32(100)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v564, v565, v562)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v573 != 0 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	v568 = int32(0)
	F_printJsonPathItem(m, l0, v564, v568, v568)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L268
	}
L265:
	;
	v578 = v10 + int32(100)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v578, v579, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v583 = int32(0)
	F_printJsonPathItem(m, l0, v578, v583, v583)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	goto L264
L268:
	;
	goto L7
L269:
	;
	goto L7
L270:
	;
	goto L7
L271:
	;
	goto L7
L272:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v603 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v605 = v10 + int32(100)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v605, v606, v603)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	v609 = int32(0)
	F_printJsonPathItem(m, l0, v605, v609, v609)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	goto L7
L279:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v620 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v622 = v10 + int32(100)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v622, v623, v620)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	v626 = int32(0)
	F_printJsonPathItem(m, l0, v622, v626, v626)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	goto L7
L286:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v637 != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v639 = v10 + int32(100)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v639, v640, v637)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L292
	}
L290:
	;
	v643 = int32(0)
	F_printJsonPathItem(m, l0, v639, v643, v643)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L1
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	goto L7
L293:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v654 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v656 = v10 + int32(100)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v656, v657, v654)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	v660 = int32(0)
	F_printJsonPathItem(m, l0, v656, v660, v660)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	goto L296
L299:
	;
	goto L7
L300:
	;
	v672 = v10 + int32(100)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v672, v673, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v677 = int32(0)
	F_printJsonPathItem(m, l0, v672, v677, v677)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v672, v684, v685)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	v688 = int32(0)
	F_printJsonPathItem(m, l0, v672, v688, v688)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	goto L7
L307:
	;
	goto L7
L308:
	;
	goto L7
L309:
	;
	v705 = v10 + int32(100)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_jspInitByBuffer(m, v705, v706, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v710 = int32(0)
	F_printJsonPathItem(m, l0, v705, v710, v710)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_appendStringInfoChar(m, l0, int32(44))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_jspInitByBuffer(m, v705, v717, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	v721 = int32(0)
	F_printJsonPathItem(m, l0, v705, v721, v721)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	goto L7
L316:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v731 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v733 = v10 + int32(100)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v733, v734, v731)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L322
	}
L320:
	;
	v737 = int32(0)
	F_printJsonPathItem(m, l0, v733, v737, v737)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	goto L319
L322:
	;
	goto L7
L323:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v748 != 0 {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v750 = v10 + int32(100)
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v750, v751, v748)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L1
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L329
	}
L327:
	;
	v754 = int32(0)
	F_printJsonPathItem(m, l0, v750, v754, v754)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	goto L7
L330:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v765 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v767 = v10 + int32(100)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v767, v768, v765)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	F_appendStringInfoChar(m, l0, int32(41))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	v771 = int32(0)
	F_printJsonPathItem(m, l0, v767, v771, v771)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	goto L7
L337:
	;
	goto L7
L338:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v786
	F_errmsg_internal(m, int32(_a_F_printJsonPathItem_45), v10)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(_a_F_printJsonPathItem_46), int32(897), int32(_a_F_printJsonPathItem_47))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L341:
	;
	goto L7
L342:
	;
	v810 = v10 + int32(100)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_jspInitByBuffer(m, v810, v811, v806)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	m.G0 = v10 + int32(128)
	return
L345:
	;
	v814 = int32(1)
	F_printJsonPathItem(m, l0, v810, v814, v814)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	goto L344
}
