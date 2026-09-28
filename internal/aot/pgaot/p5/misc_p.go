package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ParamsErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l0 == int32(0) {
		m.G0 = v6 + int32(32)
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v10 == int32(0) {
			m.G0 = v6 + int32(32)
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v13 == int32(0) {
				m.G0 = v6 + int32(32)
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v16 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
						F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v6 + int32(32)
							return
						}
					}
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
					if v19 == int32(0) {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_0), v6)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v6 + int32(32)
								return
							}
						}
					} else {
						F_set_errcontext_domain(m, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v27
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v25
							F_errcontext_msg(m, int32(_a_F_ParamsErrorCallback_1), v6+int32(16))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
func F_PartConstraintImpliedByRelConstraint(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v11 == v3 {
		v85 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = F_ConstraintImpliedByRelConstraint(m, l0, l1, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L15
	}
L2:
	;
	v14 = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v15 != v14 {
		v85 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v18 <= int32(0) {
		v85 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v14
	v26 = v3
	goto L5
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = v25 - int32(1)
	v35 = v30 + v32<<(uint(int32(3))%32)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+35)))
	if v36 != int32(118) {
		v75 = v26
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v85 = v75
	goto L1
L7:
	;
	v78 = v25 + int32(1)
	if v78 <= v18 {
		v25 = v78
		v26 = v75
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+34)))
	if v39&int32(4) != 0 {
		v75 = v26
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v44 = F_palloc0(m, int32(20))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(52)
	v57 = v30 + v42<<(uint(int32(3))%32) + v32*int32(100)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+96))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+104))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+124))
	v62 = F_makeVar(m, int32(1), base.I32_extend16_s(v25), v58, v59, v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = int32(-1)
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+12)) = uint8(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v62
	v71 = F_lappend(m, v26, v44)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v75 = v71
	goto L7
L14:
	;
	goto L6
L15:
	;
	return v89
}
func F_PopActiveSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v8 - int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v13 != 0 {
		v19 = v5
		F_pfree(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
			if v6 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
				if v25 == int32(0) {
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v29
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
					*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v29
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
					v38 = int32(3)
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(48))))
					if base.B2i32(base.Ui32(v37) < base.Ui32(v38))|base.B2i32(base.Ui32(v44) < base.Ui32(v38)) == int32(0) {
						if v37-v44 < int32(0) {
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
						} else {
						}
					} else {
						if base.Ui32(v44) <= base.Ui32(v37) {
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
						}
					}
					return
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		if v14 != 0 {
			v19 = v5
			F_pfree(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
				if v6 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
					if v25 == int32(0) {
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v29
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v29
						return
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
						v38 = int32(3)
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(48))))
						if base.B2i32(base.Ui32(v37) < base.Ui32(v38))|base.B2i32(base.Ui32(v44) < base.Ui32(v38)) == int32(0) {
							if v37-v44 < int32(0) {
								*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
							} else {
							}
						} else {
							if base.Ui32(v44) <= base.Ui32(v37) {
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
							}
						}
						return
					}
				}
			}
		} else {
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0]))
				v19 = v18
				F_pfree(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[0])) = v6
					if v6 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
						if v25 == int32(0) {
							v29 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v29
							v32 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v29
							return
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[3]))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
							v38 = int32(3)
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[1]))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v41-int32(48))))
							if base.B2i32(base.Ui32(v37) < base.Ui32(v38))|base.B2i32(base.Ui32(v44) < base.Ui32(v38)) == int32(0) {
								if v37-v44 < int32(0) {
									*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
								} else {
								}
							} else {
								if base.Ui32(v44) <= base.Ui32(v37) {
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_PopActiveSnapshot[2])) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v44
								}
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_PreventCommandDuringRecovery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+308))
		v17 = base.B2i32(v15 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandDuringRecovery[0])) = uint8(v17)
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v19 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandDuringRecovery_0), v5)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandDuringRecovery_1), int32(450), int32(_a_F_PreventCommandDuringRecovery_2))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ProcessCommittedInvalidationMessages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1 <= int32(0) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v15 = F_errstart(m, int32(11), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
				if l2 != 0 {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_0)
				} else {
					v20 = int32(_a_F_ProcessCommittedInvalidationMessages_1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v20
				F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_2), v9+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1143), int32(_a_F_ProcessCommittedInvalidationMessages_4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						if l2 != 0 {
							v34 = F_errstart(m, int32(11), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								if v34 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
									F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return
										} else {
											if l3 != 0 {
												v46 = F_GetDatabasePath(m, l3, l4)
												mBase = m.M
												v47 = m.ExcPending
												if v47 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
													F_RelationCacheInitFilePreInvalidate(m)
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
														F_pfree(m, v52)
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
															F_SIInsertDataEntries(m, l0, l1)
															mBase = m.M
															v61 = m.ExcPending
															if v61 != 0 {
																return
															} else {
																F_RelationCacheInitFilePostInvalidate(m)
																mBase = m.M
																v63 = m.ExcPending
																if v63 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												}
											} else {
												F_RelationCacheInitFilePreInvalidate(m)
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						} else {
							F_SIInsertDataEntries(m, l0, l1)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v34 = F_errstart(m, int32(11), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
							F_errmsg_internal(m, int32(_a_F_ProcessCommittedInvalidationMessages_5), v9)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ProcessCommittedInvalidationMessages_3), int32(1147), int32(_a_F_ProcessCommittedInvalidationMessages_4))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									if l3 != 0 {
										v46 = F_GetDatabasePath(m, l3, l4)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
											F_RelationCacheInitFilePreInvalidate(m)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
												F_pfree(m, v52)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
													F_SIInsertDataEntries(m, l0, l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														F_RelationCacheInitFilePostInvalidate(m)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															m.G0 = v9 + int32(32)
															return
														}
													}
												}
											}
										}
									} else {
										F_RelationCacheInitFilePreInvalidate(m)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						} else {
							if l3 != 0 {
								v46 = F_GetDatabasePath(m, l3, l4)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = v46
									F_RelationCacheInitFilePreInvalidate(m)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0]))
										F_pfree(m, v52)
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_ProcessCommittedInvalidationMessages[0])) = int32(0)
											F_SIInsertDataEntries(m, l0, l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												F_RelationCacheInitFilePostInvalidate(m)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													m.G0 = v9 + int32(32)
													return
												}
											}
										}
									}
								}
							} else {
								F_RelationCacheInitFilePreInvalidate(m)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_SIInsertDataEntries(m, l0, l1)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_RelationCacheInitFilePostInvalidate(m)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_SIInsertDataEntries(m, l0, l1)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_ProcessInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
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
	var v325 int32
	_ = v325
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v416 int64
	_ = v416
	var v420 int64
	_ = v420
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v686 int32
	_ = v686
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int64
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int64
	_ = v829
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1495 int32
	_ = v1495
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1541 int32
	_ = v1541
	var v1545 int32
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1595 int32
	_ = v1595
	var v1600 int32
	_ = v1600
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L16
	} else {
		goto L441
	}
L2:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L16
	} else {
		goto L436
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L16
	} else {
		goto L432
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L16
	} else {
		goto L428
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L16
	} else {
		goto L424
	}
L6:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L16
	} else {
		goto L419
	}
L7:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L16
	} else {
		goto L414
	}
L8:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L16
	} else {
		goto L409
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	F_LockErrorCleanup(m)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L16
	} else {
		goto L404
	}
L10:
	;
	m.G0 = v14 + int32(128)
	return
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	if v19 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4]))
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v25 = int32(_a_F_ProcessInterrupts_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5]))
	v27 = int32(_a_F_ProcessInterrupts_1)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[6]))
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[4])) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[5])) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[6])) = v30
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = v30
	F_LockErrorCleanup(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7]))
	if v236 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L16:
	;
	return
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[8])))
	if v44 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v44 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[9]))
	if v48 != int32(2) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[9])) = int32(0)
	goto L1
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	if v55 == int32(4) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[11]))
	if v80 != 0 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_2), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if v26 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3498), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[12]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[13]))
	goto L44
L36:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_6), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	if v26 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14+int32(16))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3503), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L16
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	if v106 == v108 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v112 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	switch v134 - int32(5) {
	case 0:
		goto L61
	default:
		goto L59
	case 7:
		goto L60
	case 9:
		goto L62
	}
L48:
	;
	if v112 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_7), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L16
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L58
	}
L52:
	;
	if v26 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L16
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3508), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L51
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L90
	}
L60:
	;
	v191 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L16
	} else {
		goto L79
	}
L61:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L16
	} else {
		goto L71
	}
L62:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_8), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	if v26 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14-int32(-64))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L16
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3520), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[14]))
	v169 = int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v168 + v169
	F_errmsg(m, int32(_a_F_ProcessInterrupts_9), v14+v169)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	if v26 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14+int32(80))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L16
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3526), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	if v191 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_10), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L16
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L16
	} else {
		goto L89
	}
L83:
	;
	if v26 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14+int32(112))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3531), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L16
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	goto L82
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L16
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_11), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L16
	} else {
		goto L92
	}
L92:
	;
	if v26 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v26
	F_errdetail_log(m, int32(_a_F_ProcessInterrupts_3), v14+int32(48))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L16
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3539), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15]))
	if v361 != 0 {
		goto L9
	} else {
		goto L122
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[7])) = int32(0)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16])))
	if v243 != 0 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[17]))
	if v245 <= int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v248 = m.G0
	v250 = v248 - int32(48)
	m.G0 = v250
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18]))
	v254 = int32(0)
	F_ModifyWaitEvent(m, v253, v254, int32(128), v254)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L16
	} else {
		goto L102
	}
L102:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18]))
	v262 = int32(0)
	v265 = F_WaitEventSetWait(m, v261, v262, v250, int32(3), v262)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L16
	} else {
		goto L104
	}
L103:
	;
	m.G0 = v250 + int32(48)
	if v325 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L104:
	;
	if v265 <= int32(0) {
		v325 = int32(1)
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v273 = v265
	goto L106
L106:
	;
	v282 = int32(0)
	goto L108
L107:
	;
	v325 = int32(1)
	goto L103
L108:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v250+v282<<(uint(int32(4))%32))+4))
	v297 = v295 & int32(128)
	v299 = base.B2i32(v297 == int32(0))
	if v297 != 0 {
		v325 = v299
		goto L103
	} else {
		goto L110
	}
L109:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	v309 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v309
	v314 = base.AtomicRmwOr32(m, v309, int32(_a_F_ProcessInterrupts_12), v309)
	goto L115
L110:
	;
	if v295&int32(1) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v305 = v282 + int32(1)
	if v305 == v273 {
		v325 = v299
		goto L103
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L109
L114:
	;
	v282 = v305
	goto L108
L115:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[18]))
	v318 = int32(0)
	v321 = F_WaitEventSetWait(m, v317, v318, v250, int32(3), v318)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	if int32(0) < v321 {
		v273 = v321
		goto L106
	} else {
		goto L117
	}
L117:
	;
	goto L107
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[15])) = int32(1)
	goto L98
L119:
	;
	goto L120
L120:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[17]))
	F_enable_timeout_after(m, int32(11), v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	goto L98
L122:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v363 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v432 = int32(0)
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+340))
	if v435 == v432 {
		goto L147
	} else {
		goto L148
	}
L124:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1]))
	if v374 == int32(0) {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[21]))
	if v367 == int32(0) {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[3])) = int32(1)
	goto L123
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[1])) = int32(0)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[22])))
	if v384&int32(1) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	if v407 != 0 {
		goto L8
	} else {
		goto L144
	}
L129:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[23])))
	if v399&int32(1) != 0 {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[22])) = uint8(v389)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v392 = v384 & int32(1)
	goto L129
L133:
	;
	v408 = int32(0)
	if base.B2i32(v392 == int32(0))|base.B2i32(v407 == v408) == v408 {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v404 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[23])) = uint8(v404)
	goto L136
L135:
	;
	goto L136
L136:
	;
	v407 = v399 & int32(1)
	goto L133
L137:
	;
	v416 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[24]))
	goto L140
L138:
	;
	goto L139
L139:
	;
	if v392 != 0 {
		goto L2
	} else {
		goto L143
	}
L140:
	;
	v420 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessInterrupts[25]))
	goto L141
L141:
	;
	if v416 < v420 {
		goto L128
	} else {
		goto L142
	}
L142:
	;
	goto L2
L143:
	;
	goto L128
L144:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	if v423 == int32(4) {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16])))
	if v427 == int32(0) {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	goto L123
L147:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26]))
	if v550 != 0 {
		goto L180
	} else {
		goto L181
	}
L148:
	;
	v438 = int32(0)
	v440 = base.AtomicRmwOr32(m, v434, int32(340), v438)
	if v440 == v438 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v443 = v432
	goto L150
L150:
	;
	v455 = int32(1) << (uint(v443) % 32)
	if v455&v440 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L147
L152:
	;
	v535 = v443 + int32(1)
	if v535 != int32(8) {
		v443 = v535
		goto L150
	} else {
		goto L179
	}
L153:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[20]))
	v464 = base.AtomicRmwAnd32(m, v460, int32(340), v455^int32(-1))
	switch v443 - int32(1) {
	case 0, 1, 2:
		goto L156
	case 3:
		goto L154
	case 4:
		goto L157
	case 5:
		goto L159
	case 6:
		goto L158
	default:
		goto L155
	}
L154:
	;
	F_report_recovery_conflict(m, int32(4))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L16
	} else {
		goto L178
	}
L155:
	;
	F_pgstat_report_recovery_conflict(m, int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L16
	} else {
		goto L172
	}
L156:
	;
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[27]))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v501)+24))
	goto L169
L157:
	;
	v493 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L16
	} else {
		goto L166
	}
L158:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[28]))
	if v483 == int32(0) {
		goto L147
	} else {
		goto L162
	}
L159:
	;
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[28]))
	if v468 == int32(0) {
		goto L152
	} else {
		goto L160
	}
L160:
	;
	v472 = int32(_a_F_ProcessInterrupts_13)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[29]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[30])) = int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[19]))
	F_SetLatch(m, v478)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[29])) = v473
	goto L161
L161:
	;
	goto L152
L162:
	;
	v486 = F_HoldingBufferPinThatDelaysRecovery(m)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L16
	} else {
		goto L163
	}
L163:
	;
	if v486 == int32(0) {
		goto L147
	} else {
		goto L164
	}
L164:
	;
	F_report_recovery_conflict(m, int32(7))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L16
	} else {
		goto L165
	}
L165:
	;
	goto L147
L166:
	;
	if v493 == int32(0) {
		goto L152
	} else {
		goto L167
	}
L167:
	;
	F_report_recovery_conflict(m, int32(5))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L16
	} else {
		goto L168
	}
L168:
	;
	goto L152
L169:
	;
	if base.B2i32(v502 != int32(0)) == int32(0) {
		goto L152
	} else {
		goto L170
	}
L170:
	;
	F_report_recovery_conflict(m, v443)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L16
	} else {
		goto L171
	}
L171:
	;
	goto L152
L172:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L16
	} else {
		goto L173
	}
L173:
	;
	F_errcode(m, int32(67240389))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L16
	} else {
		goto L174
	}
L174:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_14), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L16
	} else {
		goto L175
	}
L175:
	;
	F_errdetail_recovery_conflict(m, int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L16
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3324), int32(_a_F_ProcessInterrupts_15))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L16
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	goto L152
L179:
	;
	goto L151
L180:
	;
	v552 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[26])) = v552
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[31]))
	if v552 < v555 {
		goto L5
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v559 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[32]))
	if v559 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	goto L182
L184:
	;
	v561 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[32])) = v561
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[33]))
	if v561 < v564 {
		goto L4
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[34]))
	if v568 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[34])) = v570
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[35]))
	if v570 < v573 {
		goto L3
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[36]))
	v578 = int32(0)
	v581 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[16])))
	if base.B2i32(v577 == v578)|base.B2i32(v581&int32(1) == v578) != 0 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L190
L192:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[37]))
	if v599 != 0 {
		goto L197
	} else {
		goto L198
	}
L193:
	;
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[27]))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+24))
	goto L194
L194:
	;
	if v589 != int32(0) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[36])) = int32(0)
	v596 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L16
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	F_ProcessProcSignalBarrier(m)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L16
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v603 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38]))
	if v603 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L199
L201:
	;
	v604 = m.G0
	v606 = v604 - int32(160)
	m.G0 = v606
	v608 = int32(_a_F_ProcessInterrupts_16)
	v610 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v610 + int32(1)
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39]))
	if v615 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	goto L203
L203:
	;
	v959 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[40]))
	if v959 != 0 {
		goto L280
	} else {
		goto L281
	}
L204:
	;
	v633 = int32(_a_F_ProcessInterrupts_17)
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v632
	v638 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[38])) = v638
	v641 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[42]))
	if base.B2i32(v641 == v638)|base.B2i32(v641 == int32(_a_F_ProcessInterrupts_18)) == v638 {
		goto L210
	} else {
		goto L211
	}
L205:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[43]))
	v625 = F_AllocSetContextCreateInternal(m, v620, int32(_a_F_ProcessInterrupts_19), int32(0), int32(_a_F_ProcessInterrupts_20), int32(_a_F_ProcessInterrupts_21))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L16
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	F_MemoryContextReset(m, v615)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L16
	} else {
		goto L209
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39])) = v625
	v632 = v625
	goto L204
L209:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39]))
	v632 = v631
	goto L204
L210:
	;
	v651 = v641
	goto L213
L211:
	;
	v924 = v632
	goto L212
L212:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v634
	F_MemoryContextReset(m, v924)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L16
	} else {
		goto L279
	}
L213:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v651)+56))
	if v660 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v922 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[39]))
	v924 = v922
	goto L212
L215:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	if v918 != int32(_a_F_ProcessInterrupts_18) {
		v651 = v918
		goto L213
	} else {
		goto L278
	}
L216:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v651)+20))
	if v663 <= int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v671 = int32(0)
	goto L218
L218:
	;
	v679 = v671 << (uint(int32(3)) % 32)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v651)+56))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v679+v680)+4))
	if v682 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L215
L220:
	;
	v904 = v671 + int32(1)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v651)+20))
	if v904 < v905 {
		v671 = v904
		goto L218
	} else {
		goto L277
	}
L221:
	;
	v686 = v682
	goto L222
L222:
	;
	v701 = F_shm_mq_receive(m, v686, v606+int32(56), v606+int32(52), int32(1))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L16
	} else {
		goto L224
	}
L223:
	;
	goto L220
L224:
	;
	if v701 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	if v701 == int32(1) {
		goto L220
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v722 = v606 + int32(36)
	F_initStringInfo(m, v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L16
	} else {
		goto L233
	}
L228:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L16
	} else {
		goto L229
	}
L229:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L16
	} else {
		goto L230
	}
L230:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_22), int32(0))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L16
	} else {
		goto L231
	}
L231:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_23), int32(1129), int32(_a_F_ProcessInterrupts_19))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L16
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
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v606)+52))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v606)+56))
	F_appendBinaryStringInfo(m, v722, v725, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L16
	} else {
		goto L234
	}
L234:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v651)+64))
	if v729 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v743 = F_pq_getmsgbyte(m, v606+int32(36))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L16
	} else {
		goto L244
	}
L236:
	;
	v732 = v729 + v671
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732))))
	if v733 != 0 {
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v734 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v732))) = uint8(v734)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v651)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v651)+60)) = v736 + v734
	goto L235
L238:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v606)+36))
	F_pfree(m, v886)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L16
	} else {
		goto L275
	}
L239:
	;
	v872 = v606 + int32(36)
	v874 = F_pq_getmsgint(m, v872, int32(4))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L16
	} else {
		goto L270
	}
L240:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L16
	} else {
		goto L267
	}
L241:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v651)+56))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v847+v679)+4))
	F_shm_mq_detach(m, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L16
	} else {
		goto L266
	}
L242:
	;
	v791 = v606 + int32(36)
	v793 = F_pq_getmsgint(m, v791, int32(4))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L16
	} else {
		goto L259
	}
L243:
	;
	F_pq_parse_errornotice(m, v606+int32(36), v606+int32(60))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L16
	} else {
		goto L245
	}
L244:
	;
	v745 = base.I32_extend8_s(v743)
	switch v745 - int32(65) {
	case 0:
		goto L239
	default:
		goto L240
	case 4, 13:
		goto L243
	case 15:
		goto L242
	case 23:
		goto L241
	}
L245:
	;
	v754 = int32(21)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v606)+60))
	if v754 <= v755 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v758 = v754
	goto L248
L247:
	;
	v758 = v755
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606)+60)) = v758
	v761 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[44]))
	if v761 != int32(2) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v606)+108))
	if v764 != 0 {
		goto L253
	} else {
		goto L254
	}
L250:
	;
	goto L251
L251:
	;
	v779 = int32(_a_F_ProcessInterrupts_24)
	v780 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45]))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v651)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45])) = v782
	F_ThrowErrorData(m, v606+int32(60))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L16
	} else {
		goto L258
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606)+108)) = v776
	goto L251
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606)+20)) = int32(_a_F_ProcessInterrupts_25)
	*(*int32)(unsafe.Add(mBase, uint32(v606)+16)) = v764
	v771 = F_psprintf(m, int32(_a_F_ProcessInterrupts_26), v606+int32(16))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L16
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v774 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_25))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L16
	} else {
		goto L257
	}
L256:
	;
	v776 = v771
	goto L252
L257:
	;
	v776 = v774
	goto L252
L258:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45])) = v780
	goto L238
L259:
	;
	v795 = F_pq_getmsgint64(m, v791)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L16
	} else {
		goto L260
	}
L260:
	;
	F_pq_getmsgend(m, v791)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L16
	} else {
		goto L261
	}
L261:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[46]))
	if v801 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L238
L263:
	;
	goto L262
L264:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[47])))
	if v805&int32(1) == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v810 = int32(_a_F_ProcessInterrupts_27)
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	v813 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v812 + v813
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	*(*int32)(unsafe.Add(mBase, uint32(v801))) = v816 + v813
	v820 = int32(0)
	v822 = int32(_a_F_ProcessInterrupts_28)
	v823 = base.AtomicRmwOr32(m, v820, v822, v820)
	v828 = v801 + v793<<(uint(int32(3))%32) + int32(232)
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v828)))
	*(*int64)(unsafe.Add(mBase, uint32(v828))) = v829 + v795
	v835 = base.AtomicRmwOr32(m, v820, v822, v820)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	*(*int32)(unsafe.Add(mBase, uint32(v801))) = v836 + v813
	v842 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[2])) = v842 - v813
	goto L263
L266:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v651)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v852+v679)+4)) = int32(0)
	goto L238
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606))) = v745
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v606)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+4)) = v861
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_29), v606)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L16
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_23), int32(1251), int32(_a_F_ProcessInterrupts_30))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L16
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	v876 = F_pq_getmsgrawstring(m, v872)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L16
	} else {
		goto L271
	}
L271:
	;
	v878 = F_pq_getmsgrawstring(m, v872)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L16
	} else {
		goto L272
	}
L272:
	;
	F_pq_endmessage(m, v872)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L16
	} else {
		goto L273
	}
L273:
	;
	F_NotifyMyFrontEnd(m, v876, v878, v874)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L16
	} else {
		goto L274
	}
L274:
	;
	goto L238
L275:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v651)+56))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v889+v679)+4))
	if v891 != 0 {
		v686 = v891
		goto L222
	} else {
		goto L276
	}
L276:
	;
	goto L223
L277:
	;
	goto L219
L278:
	;
	goto L214
L279:
	;
	v938 = int32(_a_F_ProcessInterrupts_16)
	v940 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v940 - int32(1)
	m.G0 = v606 + int32(160)
	goto L203
L280:
	;
	F_ProcessLogMemoryContextInterrupt(m)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L16
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48]))
	if v963 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	goto L282
L284:
	;
	v964 = m.G0
	v966 = v964 - int32(176)
	m.G0 = v966
	v968 = int32(_a_F_ProcessInterrupts_16)
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v970 + int32(1)
	v975 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49]))
	if v975 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L285:
	;
	goto L286
L286:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50]))
	if v1183 != 0 {
		goto L335
	} else {
		goto L336
	}
L287:
	;
	v993 = int32(0)
	v994 = int32(_a_F_ProcessInterrupts_17)
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v992
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[48])) = v993
	v1002 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[51]))
	if v1002 != 0 {
		goto L293
	} else {
		goto L294
	}
L288:
	;
	v980 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[43]))
	v985 = F_AllocSetContextCreateInternal(m, v980, int32(_a_F_ProcessInterrupts_31), int32(0), int32(_a_F_ProcessInterrupts_20), int32(_a_F_ProcessInterrupts_21))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L16
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	F_MemoryContextReset(m, v975)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L16
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49])) = v985
	v992 = v985
	goto L287
L292:
	;
	v991 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49]))
	v992 = v991
	goto L287
L293:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+4))
	if int32(0) < v1003 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v1148 = v992
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v995
	F_MemoryContextReset(m, v1148)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L16
	} else {
		goto L334
	}
L296:
	;
	v1008 = v993
	goto L299
L297:
	;
	goto L298
L298:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[49]))
	v1148 = v1146
	goto L295
L299:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+12))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1017+v1008<<(uint(int32(2))%32))))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	if v1022 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	goto L298
L301:
	;
	v1131 = v1008 + int32(1)
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+4))
	if v1131 < v1132 {
		v1008 = v1131
		goto L299
	} else {
		goto L333
	}
L302:
	;
	v1030 = F_shm_mq_receive(m, v1022, v966+int32(72), v966+int32(68), int32(1))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L16
	} else {
		goto L306
	}
L303:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v966)+52))
	F_pfree(m, v1126)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L16
	} else {
		goto L332
	}
L304:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L16
	} else {
		goto L328
	}
L305:
	;
	v1033 = v966 + int32(52)
	F_initStringInfo(m, v1033)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L16
	} else {
		goto L307
	}
L306:
	;
	switch v1030 {
	case 0:
		goto L305
	case 1:
		goto L301
	default:
		goto L304
	}
L307:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v966)+68))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v966)+72))
	F_appendBinaryStringInfo(m, v1033, v1036, v1037)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L16
	} else {
		goto L308
	}
L308:
	;
	v1040 = F_pq_getmsgbyte(m, v1033)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L16
	} else {
		goto L311
	}
L309:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L16
	} else {
		goto L325
	}
L310:
	;
	F_pq_parse_errornotice(m, v966+int32(52), v966+int32(76))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L16
	} else {
		goto L312
	}
L311:
	;
	v1042 = base.I32_extend8_s(v1040)
	switch v1042 - int32(65) {
	case 0, 13:
		goto L303
	default:
		goto L309
	case 4:
		goto L310
	}
L312:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v966)+124))
	if v1051 != 0 {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[52]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[45])) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v966)+124)) = v1063
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L16
	} else {
		goto L319
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v966)+36)) = int32(_a_F_ProcessInterrupts_32)
	*(*int32)(unsafe.Add(mBase, uint32(v966)+32)) = v1051
	v1058 = F_psprintf(m, int32(_a_F_ProcessInterrupts_26), v966+int32(32))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L16
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v1061 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_32))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L16
	} else {
		goto L318
	}
L317:
	;
	v1063 = v1058
	goto L313
L318:
	;
	v1063 = v1061
	goto L313
L319:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L16
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_33), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L16
	} else {
		goto L321
	}
L321:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L16
	} else {
		goto L322
	}
L322:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v966)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v966)+16)) = v1083
	F_errcontext_msg(m, int32(_a_F_ProcessInterrupts_34), v966+int32(16))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L16
	} else {
		goto L323
	}
L323:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_35), int32(1060), int32(_a_F_ProcessInterrupts_36))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L16
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v966))) = v1042
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v966)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v966)+4)) = v1100
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_37), v966)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L16
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_35), int32(1074), int32(_a_F_ProcessInterrupts_36))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L16
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L16
	} else {
		goto L329
	}
L329:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_38), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L16
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_35), int32(1146), int32(_a_F_ProcessInterrupts_31))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L16
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	goto L301
L333:
	;
	goto L300
L334:
	;
	v1162 = int32(_a_F_ProcessInterrupts_16)
	v1164 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v1164 - int32(1)
	m.G0 = v966 + int32(176)
	goto L286
L335:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[50])) = int32(0)
	v1188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[10]))
	if v1188 == int32(7) {
		goto L338
	} else {
		goto L339
	}
L336:
	;
	goto L337
L337:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[53]))
	if v1226 == int32(0) {
		goto L10
	} else {
		goto L355
	}
L338:
	;
	v1193 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L16
	} else {
		goto L341
	}
L339:
	;
	goto L340
L340:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessInterrupts[54])))
	if v1208 != 0 {
		goto L348
	} else {
		goto L349
	}
L341:
	;
	if v1193 != 0 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_39), int32(0))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L16
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L16
	} else {
		goto L347
	}
L345:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_40), int32(1400), int32(_a_F_ProcessInterrupts_41))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L16
	} else {
		goto L346
	}
L346:
	;
	goto L344
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L16
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	goto L337
L351:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L16
	} else {
		goto L352
	}
L352:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_42), int32(0))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L16
	} else {
		goto L353
	}
L353:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_40), int32(1414), int32(_a_F_ProcessInterrupts_41))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L16
	} else {
		goto L354
	}
L354:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L355:
	;
	v1229 = m.G0
	v1231 = v1229 - int32(160)
	m.G0 = v1231
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[55]))
	if v1234 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L356:
	;
	goto L10
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L16
	} else {
		goto L399
	}
L358:
	;
	m.G0 = v1231 + int32(160)
	goto L356
L359:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1234)+8))
	if v1237 == int32(0) {
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1240 = int32(_a_F_ProcessInterrupts_16)
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v1242 + int32(1)
	v1247 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[56]))
	if v1247 == int32(0) {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	v1265 = int32(_a_F_ProcessInterrupts_17)
	v1266 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v1264
	v1270 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[53])) = v1270
	v1273 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[55]))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+8))
	if v1274 == v1270 {
		goto L367
	} else {
		goto L368
	}
L362:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[43]))
	v1257 = F_AllocSetContextCreateInternal(m, v1252, int32(_a_F_ProcessInterrupts_43), int32(0), int32(_a_F_ProcessInterrupts_20), int32(_a_F_ProcessInterrupts_21))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L16
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	F_MemoryContextReset(m, v1247)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L16
	} else {
		goto L366
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[56])) = v1257
	v1264 = v1257
	goto L361
L366:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[56]))
	v1264 = v1263
	goto L361
L367:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[41])) = v1266
	v1384 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[56]))
	F_MemoryContextReset(m, v1384)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L16
	} else {
		goto L398
	}
L368:
	;
	v1278 = v1274
	goto L369
L369:
	;
	v1293 = F_shm_mq_receive(m, v1278, v1231+int32(56), v1231+int32(52), int32(1))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L16
	} else {
		goto L373
	}
L370:
	;
	goto L367
L371:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[55]))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+8))
	if v1369 != 0 {
		v1278 = v1369
		goto L369
	} else {
		goto L397
	}
L372:
	;
	v1296 = v1231 + int32(36)
	F_initStringInfo(m, v1296)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L16
	} else {
		goto L374
	}
L373:
	;
	switch v1293 {
	case 0:
		goto L372
	case 1:
		goto L367
	case 2:
		goto L357
	default:
		goto L371
	}
L374:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+52))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+56))
	F_appendBinaryStringInfo(m, v1296, v1299, v1300)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L16
	} else {
		goto L375
	}
L375:
	;
	v1303 = F_pq_getmsgbyte(m, v1296)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L16
	} else {
		goto L380
	}
L376:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+36))
	F_pfree(m, v1363)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L16
	} else {
		goto L396
	}
L377:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[55]))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+8))
	F_shm_mq_detach(m, v1355)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L16
	} else {
		goto L395
	}
L378:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L16
	} else {
		goto L392
	}
L379:
	;
	F_pq_parse_errornotice(m, v1231+int32(36), v1231+int32(60))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L16
	} else {
		goto L381
	}
L380:
	;
	v1305 = base.I32_extend8_s(v1303)
	switch v1305 - int32(69) {
	case 0, 9:
		goto L379
	default:
		goto L378
	case 19:
		goto L377
	}
L381:
	;
	v1314 = int32(21)
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+60))
	if v1314 <= v1315 {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1318 = v1314
	goto L384
L383:
	;
	v1318 = v1315
	goto L384
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+60)) = v1318
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+108))
	if v1320 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+108)) = v1332
	F_ThrowErrorData(m, v1231+int32(60))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L16
	} else {
		goto L391
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+20)) = int32(_a_F_ProcessInterrupts_44)
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+16)) = v1320
	v1327 = F_psprintf(m, int32(_a_F_ProcessInterrupts_26), v1231+int32(16))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L16
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	v1330 = F_pstrdup(m, int32(_a_F_ProcessInterrupts_44))
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L16
	} else {
		goto L390
	}
L389:
	;
	v1332 = v1327
	goto L385
L390:
	;
	v1332 = v1330
	goto L385
L391:
	;
	goto L376
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1231))) = v1305
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+4)) = v1343
	F_errmsg_internal(m, int32(_a_F_ProcessInterrupts_45), v1231)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L16
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_46), int32(_a_F_ProcessInterrupts_47), int32(_a_F_ProcessInterrupts_48))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L16
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[55]))
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+8)) = int32(0)
	goto L376
L396:
	;
	goto L371
L397:
	;
	goto L370
L398:
	;
	v1387 = int32(_a_F_ProcessInterrupts_16)
	v1389 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[0])) = v1389 - int32(1)
	goto L358
L399:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L16
	} else {
		goto L400
	}
L400:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_49), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L16
	} else {
		goto L401
	}
L401:
	;
	F_errhint(m, int32(_a_F_ProcessInterrupts_50), int32(0))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L16
	} else {
		goto L402
	}
L402:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_46), int32(_a_F_ProcessInterrupts_51), int32(_a_F_ProcessInterrupts_43))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L16
	} else {
		goto L403
	}
L403:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L404:
	;
	v1447 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessInterrupts[9])) = v1447
	F_errstart_cold(m, int32(22), v1447)
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L16
	} else {
		goto L405
	}
L405:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L16
	} else {
		goto L406
	}
L406:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_52), int32(0))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L16
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3570), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L16
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L16
	} else {
		goto L410
	}
L410:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L16
	} else {
		goto L411
	}
L411:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_53), int32(0))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L16
	} else {
		goto L412
	}
L412:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3629), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L16
	} else {
		goto L413
	}
L413:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L414:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L16
	} else {
		goto L415
	}
L415:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L16
	} else {
		goto L416
	}
L416:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_54), int32(0))
	mBase = m.M
	v1495 = m.ExcPending
	if v1495 != 0 {
		goto L16
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3636), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L16
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L16
	} else {
		goto L420
	}
L420:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L16
	} else {
		goto L421
	}
L421:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_55), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L16
	} else {
		goto L422
	}
L422:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3649), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		goto L16
	} else {
		goto L423
	}
L423:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L424:
	;
	F_errcode(m, int32(50463042))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L16
	} else {
		goto L425
	}
L425:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_56), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L16
	} else {
		goto L426
	}
L426:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3670), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L16
	} else {
		goto L427
	}
L427:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L428:
	;
	F_errcode(m, int32(67240258))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L16
	} else {
		goto L429
	}
L429:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_57), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L16
	} else {
		goto L430
	}
L430:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3683), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L16
	} else {
		goto L431
	}
L431:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L432:
	;
	F_errcode(m, int32(84017605))
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L16
	} else {
		goto L433
	}
L433:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_58), int32(0))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L16
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3696), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L16
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L16
	} else {
		goto L437
	}
L437:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L16
	} else {
		goto L438
	}
L438:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_59), int32(0))
	mBase = m.M
	v1579 = m.ExcPending
	if v1579 != 0 {
		goto L16
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3622), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L16
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(67371461))
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L16
	} else {
		goto L442
	}
L442:
	;
	F_errmsg(m, int32(_a_F_ProcessInterrupts_60), int32(0))
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L16
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_ProcessInterrupts_4), int32(3493), int32(_a_F_ProcessInterrupts_5))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L16
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessNotifyInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	goto L2
L1:
	;
	return
L2:
	;
	if v5 != int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L5
L5:
	;
	v15 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1])) = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[2]))
	if v18 == v15 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[1]))
	if v142 != 0 {
		goto L5
	} else {
		goto L33
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v22)+808))
	if v24 != int64(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v89 == int64(0) {
		goto L7
	} else {
		goto L13
	}
L10:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)+752))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v22)+728))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v22)+704))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v22)+680))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v22)+656))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v22)+632))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v22)+608))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v22)+584))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v22)+560))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v22)+536))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v22)+512))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v22)+488))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v22)+464))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v22)+440))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v22)+416))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v22)+392))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v22)+368))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v22)+344))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v22)+320))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v22)+296))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v22)+272))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v22)+248))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v22)+224))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v22)+200))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v22)+176))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v22)+152))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v22)+128))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v22)+104))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v22)+80))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v22)+56))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v22)+32))
	v89 = v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + v23))))))))))))))))))))))))))))))
	goto L12
L11:
	;
	v89 = v23
	goto L12
L12:
	;
	goto L9
L13:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v93 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L16
	} else {
		goto L21
	}
L15:
	;
	v98 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return
L17:
	;
	if v98 == int32(0) {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_0), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(3086), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	if l0 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[4]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v120 = m.T0[v119].(func(*base.Module) int32)(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessNotifyInterrupt[3])))
	if v123 != int32(1) {
		goto L7
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v128 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	if v128 == int32(0) {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errmsg_internal(m, int32(_a_F_ProcessNotifyInterrupt_2), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_ProcessNotifyInterrupt_1), int32(3110), int32(_a_F_ProcessNotifyInterrupt_0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	goto L7
L33:
	;
	goto L6
}
func F_PromoteIsTriggered(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])))
	if v5 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
		v12 = base.AtomicRmwXchg32(m, v9, int32(96), int32(1))
		if v12 != 0 {
			F_s_lock(m, v9+int32(96), int32(_a_F_PromoteIsTriggered_0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
				*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v23)
				v25 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
				v28 = v23
				return v28 & int32(1)
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[1]))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
			*(*uint8)(unsafe.Add(mBase, _c_F_PromoteIsTriggered[0])) = uint8(v23)
			v25 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+96)), uint32(v25))
			v28 = v23
			return v28 & int32(1)
		}
	} else {
		v28 = int32(1)
		return v28 & int32(1)
	}
}
func F_p_isasclet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v6 != int32(1) {
		v41 = v2
		return v41
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9+v10))))
		if v12 < int32(0) {
			v41 = v2
			return v41
		} else {
			v15 = F_pg_database_locale(m)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v22 = int32(2)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21<<(uint(v22)%32))))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v26 < v22 {
					v36 = F_pg_database_locale(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = F_pg_iswalpha(m, v25, v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = v38
							return v41
						}
					}
				} else {
					v29 = int32(1)
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)))
					if v30 != v29 {
						v36 = F_pg_database_locale(m)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = F_pg_iswalpha(m, v25, v36)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v41 = v38
								return v41
							}
						}
					} else {
						if base.Ui32(int32(127)) < base.Ui32(v25) {
							v41 = v29
							return v41
						} else {
							v36 = F_pg_database_locale(m)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = F_pg_iswalpha(m, v25, v36)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v41 = v38
									return v41
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_p_isnotalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	v4 = F_pg_database_locale(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(v11)%32))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 < v11 {
			v24 = F_pg_database_locale(m)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = F_pg_iswalnum(m, v14, v24)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v30 = v26 ^ int32(1)
					return v30
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			if v18 != int32(1) {
				v24 = F_pg_database_locale(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_pg_iswalnum(m, v14, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v30 = v26 ^ int32(1)
						return v30
					}
				}
			} else {
				if base.Ui32(int32(127)) < base.Ui32(v14) {
					v30 = int32(0)
					return v30
				} else {
					v24 = F_pg_database_locale(m)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_pg_iswalnum(m, v14, v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v30 = v26 ^ int32(1)
							return v30
						}
					}
				}
			}
		}
	}
}
func F_pad(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(256)
	m.G0 = v9
	if l4&int32(_a_F_pad_0)|base.B2i32(l2 <= l3) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = l2 - l3
	v18 = int32(256)
	v20 = base.B2i32(base.Ui32(v17) < base.Ui32(v18))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v9 + int32(256)
	return
L4:
	;
	v21 = v17
	goto L6
L5:
	;
	v21 = v18
	goto L6
L6:
	;
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v20 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v2)
	v28 = v9 + v21
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(1)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(3)))) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(2)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(7)) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)) = uint8(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v28-int32(4)))) = uint8(v2)
	if base.Ui32(v21) < base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v53 = (int32(0) - v9) & int32(3)
	v54 = v9 + v53
	v58 = v2 & int32(255) * int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v58
	v62 = (v21 - v53) & int32(-4)
	v63 = v54 + v62
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(4)))) = v58
	if base.Ui32(v62) < base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(8)))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(12)))) = v58
	if base.Ui32(v62) < base.Ui32(int32(25)) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(16)))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(20)))) = v58
	v89 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v63-v89))) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v63-int32(28)))) = v58
	v98 = v54&int32(4) | v89
	v99 = v62 - v98
	if base.Ui32(v99) < base.Ui32(int32(32)) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v104 = base.I64_extend_i32_u(v58) * int64(4294967297)
	v107 = v98 + v54
	v108 = v99
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v107)+24)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107)+16)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = v104
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = v104
	v116 = int32(32)
	v119 = v108 - v116
	if base.Ui32(int32(31)) < base.Ui32(v119) {
		v107 = v107 + v116
		v108 = v119
		goto L16
	} else {
		goto L18
	}
L17:
	;
	goto L8
L18:
	;
	goto L17
L19:
	;
	v133 = v17
	goto L22
L20:
	;
	v146 = v17
	goto L21
L21:
	;
	F_out(m, l0, v9, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L24
	} else {
		goto L27
	}
L22:
	;
	F_out(m, l0, v9, int32(256))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v146 = v140
	goto L21
L24:
	;
	return
L25:
	;
	v140 = v133 - int32(256)
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		v133 = v140
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L3
}
func F_pair_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = l0
	goto L1
L1:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v23-int32(9)))&base.B2i32(v23 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v37 = v14 + base.B2i32(v23 == int32(40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v37
	v41 = F_float8in_internal(m, v37, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v14 = v14 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l1))) = v41
	if l6 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v12 + int32(16)
	return v146
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != int32(44) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v48 != int32(453) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v146 = int32(0)
	goto L8
L13:
	;
	v128 = int32(0)
	v129 = F_errsave_start(m, l6)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L30
	}
L14:
	;
	v64 = F_float8in_internal(m, v57, v12+int32(12), l4, l5, l6)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l2))) = v64
	if l6 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v23 != int32(40) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v69 != int32(453) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v72 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v146 = int32(0)
	goto L8
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v80 = v78 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v80
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 != int32(41) {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v85 = v80
	goto L23
L23:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v94-int32(9)))&base.B2i32(v94 != int32(32)) != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v103 = v85 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v103
	v85 = v103
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v114
	v146 = int32(1)
	goto L8
L27:
	;
	goto L28
L28:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v146 = int32(1)
	goto L8
L30:
	;
	if v129 == int32(0) {
		v146 = v128
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
	F_errmsg(m, int32(_a_F_pair_decode_0), v12)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, l6, int32(_a_F_pair_decode_1), int32(252), int32(_a_F_pair_decode_2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v146 = v128
	goto L8
}
func F_pairingheap_remove(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_pairingheap_remove_first(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v18 != l1 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	v20 = int32(4)
	goto L8
L7:
	;
	v20 = int32(0)
	goto L8
L8:
	;
	v21 = v15 + v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v24 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22
	if v22 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v89
	if v22 == int32(0) {
		goto L9
	} else {
		goto L48
	}
L14:
	;
	v89 = v23
	goto L13
L15:
	;
	goto L16
L16:
	;
	v29 = int32(0)
	v30 = v23
	goto L17
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v37 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v29 == int32(0) {
		v89 = v57
		goto L13
	} else {
		goto L34
	}
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v29
	v57 = v30
	goto L19
L21:
	;
	goto L22
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32) int32)(m, v30, v37, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v47 = base.B2i32(v44 < int32(0))
	if v44 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v48 = v30
	goto L26
L25:
	;
	v48 = v37
	goto L26
L26:
	;
	if v44 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v49 = v37
	goto L29
L28:
	;
	v49 = v30
	goto L29
L29:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v48
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v48
	if v41 != 0 {
		v29 = v49
		v30 = v41
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v57 = v49
	goto L19
L34:
	;
	v64 = v57
	v66 = v29
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = m.T0[v74].(func(*base.Module, int32, int32, int32) int32)(m, v64, v66, v73)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	v89 = v80
	goto L13
L37:
	;
	v78 = base.B2i32(v75 < int32(0))
	if v75 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v79 = v64
	goto L40
L39:
	;
	v79 = v66
	goto L40
L40:
	;
	if v75 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v80 = v66
	goto L43
L42:
	;
	v80 = v64
	goto L43
L43:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v79
	goto L46
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v79
	if v72 != 0 {
		v64 = v80
		v66 = v72
		goto L35
	} else {
		goto L47
	}
L47:
	;
	goto L36
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v89
	return
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v108
	goto L9
}
func F_paramlist_parser_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(863)
	return
}
func F_parse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v371 int32
	_ = v371
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = m.T0[v16].(func(*base.Module) int32)(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v13 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v25 = v23
	goto L8
L7:
	;
	v25 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
	return int32(0)
L9:
	;
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v48)+4)) = v49
	v51 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(v51)
	v53 = int32(380)
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v48)+20)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(281479271677952)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v63 != 0 {
		v371 = v6
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v29
	v48 = v13
	goto L9
L11:
	;
	goto L12
L12:
	;
	v33 = F_palloc_extended(m, int32(88), int32(2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v33 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+84)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v33
	v48 = v33
	goto L9
L17:
	;
	v41 = v39
	goto L19
L18:
	;
	v41 = int32(12)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v41
	return int32(0)
L20:
	;
	return v371
L21:
	;
	v65 = v48 + int32(20)
	v77 = v6
	goto L23
L22:
	;
	if l1 != v278 {
		goto L94
	} else {
		goto L95
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v79 = F_newstate(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v278 = v277
	goto L22
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = F_newstate(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v89 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v92 <= v93 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_parse[0]))
	if v169 != 0 {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	F_createarc(m, v87, int32(110), int32(0), l3, v79)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L55
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v95 == int32(0) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v117 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L39:
	;
	v103 = v95
	goto L40
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	if v110 != v79 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L35
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v116 != 0 {
		v103 = v116
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)))
	if v112 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v113 == int32(110) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	v125 = v117
	goto L48
L48:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
	if v132 != l3 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L35
L50:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+24))
	if v138 != 0 {
		v125 = v138
		goto L48
	} else {
		goto L54
	}
L51:
	;
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125)+4)))
	if v134 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v135 == int32(110) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L49
L55:
	;
	goto L34
L56:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v172 <= v173 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L58
L60:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v247 != 0 {
		goto L82
	} else {
		goto L83
	}
L61:
	;
	F_createarc(m, v167, int32(110), int32(0), v82, l4)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L81
	}
L62:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v175 == int32(0) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v197 == int32(0) {
		goto L61
	} else {
		goto L73
	}
L65:
	;
	v183 = v175
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v190 != l4 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L61
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	if v196 != 0 {
		v183 = v196
		goto L66
	} else {
		goto L72
	}
L69:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+4)))
	if v192 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v193 == int32(110) {
		goto L60
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L67
L73:
	;
	v205 = v197
	goto L74
L74:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v212 != v82 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L61
L76:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+24))
	if v218 != 0 {
		v205 = v218
		goto L74
	} else {
		goto L80
	}
L77:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v205)+4)))
	if v214 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v215 == int32(110) {
		goto L60
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	goto L75
L81:
	;
	goto L60
L82:
	;
	return int32(0)
L83:
	;
	goto L84
L84:
	;
	v250 = int32(0)
	v252 = F_parsebranch(m, l0, l1, l2, v79, v82, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 != 0 {
		v371 = v250
		goto L20
	} else {
		goto L86
	}
L86:
	;
	if v77 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v261 = v257 | v260
	v270 = v257&int32(28) | v261<<(uint(int32(1))%32)&(v261<<(uint(int32(2))%32))&int32(4) | v260
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v272 != int32(124) {
		v278 = v272
		goto L22
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+24)) = v252
	goto L87
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v252
	goto L87
L91:
	;
	v275 = F_next(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v275 != 0 {
		v77 = v252
		goto L23
	} else {
		goto L93
	}
L93:
	;
	goto L24
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v283 != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v288 == v252 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v285 = v283
	goto L99
L98:
	;
	v285 = int32(8)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v285
	goto L96
L100:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
	if v290 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v316&int32(28) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
	F_pfree(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v302 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v302)
	v304 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v304
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v304
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v308 != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	F_pfree(m, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
	F_pfree(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = int32(0)
	goto L105
L109:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v48
	return v252
L110:
	;
	goto L111
L111:
	;
	F_pfree(m, v48)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	return v252
L113:
	;
	if v288 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v371 = v48
	goto L20
L116:
	;
	v328 = v288
	goto L119
L117:
	;
	goto L118
L118:
	;
	v348 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v348)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = int32(0)
	goto L115
L119:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	F_freesubre(m, l0, v328)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	if v333 != 0 {
		v328 = v333
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
}
func F_parseCheckAggregates(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
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
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v688 int32
	_ = v688
	var v689 int64
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v738 int32
	_ = v738
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	v3 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v3
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v25 == v3 {
		v255 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v262 == int32(0) {
		v368 = v3
		v373 = v3
		goto L41
	} else {
		goto L42
	}
L2:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+104)))
	v30 = F_expand_grouping_sets(m, v25, v28, int32(_a_F_parseCheckAggregates_0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	if v30 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v57 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_errcode(m, int32(16777477))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v45
	goto L13
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v47 = v46
	goto L13
L13:
	;
	v48 = F_exprLocation(m, v47)
	mBase = m.M
	F_parser_errposition(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1173), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v236 != int32(1) {
		v255 = v229
		goto L1
	} else {
		goto L39
	}
L17:
	;
	v229 = int32(0)
	goto L16
L18:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v61 <= v60 {
		v229 = v57
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v67 = v60
	v75 = v57
	goto L20
L20:
	;
	v82 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v67<<(uint(int32(2))%32))))
	if base.B2i32(v75 == v82)|base.B2i32(v90 == v82) != 0 {
		v179 = v82
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v229 = v179
	goto L16
L22:
	;
	if v179 == int32(0) {
		goto L17
	} else {
		goto L37
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v94 <= int32(0) {
		v179 = v82
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v101 = v82
	v104 = v82
	v107 = v94
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v115 <= int32(0) {
		v158 = v101
		v164 = v107
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v179 = v158
	goto L22
L27:
	;
	v173 = v104 + int32(1)
	if v173 < v164 {
		v101 = v158
		v104 = v173
		v107 = v164
		goto L25
	} else {
		goto L36
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v104<<(uint(int32(2))%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v131 = int32(0)
	goto L29
L29:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v123+v131<<(uint(int32(2))%32))))
	if v122 != v146 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v151 = F_lappend_int(m, v101, v122)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L35
	}
L31:
	;
	v149 = v131 + int32(1)
	if v149 != v115 {
		v131 = v149
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	v158 = v101
	v164 = v107
	goto L27
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v158 = v151
	v164 = v153
	goto L27
L36:
	;
	goto L26
L37:
	;
	v196 = v67 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v196 < v197 {
		v67 = v196
		v75 = v179
		goto L20
	} else {
		goto L38
	}
L38:
	;
	goto L21
L39:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v239 == int32(0) {
		v255 = v229
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = int32(0)
	v255 = v229
	goto L1
L41:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v377 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L42:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v265 <= int32(0) {
		v368 = v3
		v373 = v3
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v268 = int32(0)
	if v265 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v272 = int32(0)
	if v272 < v265 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v332 = v268
	v333 = v268
	v343 = v3
	goto L46
L46:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v349 = int32(2)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v332<<(uint(v349)%32))))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	switch v353 - v349 {
	case 0:
		v368 = int32(1)
		v373 = v343
		goto L41
	default:
		v358 = v343
		goto L60
	case 4:
		goto L61
	}
L47:
	;
	v275 = v265
	goto L49
L48:
	;
	v275 = v272
	goto L49
L49:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v284 = v268
	v286 = v3
	v295 = v3
	v296 = v3
	goto L50
L50:
	;
	v299 = int32(1)
	v301 = int32(2)
	v303 = v280 + v284<<(uint(v301)%32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	switch v305 - v301 {
	case 0:
		v311 = v299
		v312 = v295
		goto L52
	default:
		v310 = v295
		goto L53
	case 4:
		goto L54
	}
L51:
	;
	if v275&int32(1) == int32(0) {
		v368 = v320
		v373 = v321
		goto L41
	} else {
		goto L59
	}
L52:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	switch v314 - int32(2) {
	case 0:
		v320 = v299
		v321 = v312
		goto L55
	default:
		v319 = v312
		goto L56
	case 4:
		goto L57
	}
L53:
	;
	v311 = v286
	v312 = v310
	goto L52
L54:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+92)))
	v310 = v308 | v295
	goto L53
L55:
	;
	v322 = int32(2)
	v323 = v284 + v322
	v325 = v296 + v322
	if v325 != v275&int32(2147483646) {
		v284 = v323
		v286 = v320
		v295 = v321
		v296 = v325
		goto L50
	} else {
		goto L58
	}
L56:
	;
	v320 = v311
	v321 = v319
	goto L55
L57:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+92)))
	v319 = v317 | v312
	goto L56
L58:
	;
	goto L51
L59:
	;
	v332 = v323
	v333 = v320
	v343 = v321
	goto L46
L60:
	;
	v368 = v333
	v373 = v358
	goto L41
L61:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+92)))
	v358 = v356 | v343
	goto L60
L62:
	;
	if v368&int32(1) != 0 {
		goto L105
	} else {
		goto L106
	}
L63:
	;
	v558 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v381 = int32(0)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v381 < v383 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v390 = int32(0)
	v393 = v381
	goto L69
L67:
	;
	v426 = v381
	goto L68
L68:
	;
	if v426 == int32(0) {
		v558 = v381
		goto L62
	} else {
		goto L77
	}
L69:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v377)+12))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v405+v390<<(uint(int32(2))%32))))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v411 = F_get_sortgroupclause_tle(m, v409, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L71
	}
L70:
	;
	v426 = v415
	goto L68
L71:
	;
	if v411 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v413 = F_lappend(m, v393, v411)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L3
	} else {
		goto L75
	}
L73:
	;
	v415 = v393
	goto L74
L74:
	;
	v417 = v390 + int32(1)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
	if v417 < v418 {
		v390 = v417
		v393 = v415
		goto L69
	} else {
		goto L76
	}
L75:
	;
	v415 = v413
	goto L74
L76:
	;
	goto L70
L77:
	;
	v440 = int32(0)
	v445 = F_palloc0(m, int32(136))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = int32(9)
	*(*int64)(unsafe.Add(mBase, uint32(v445))) = int64(101)
	v453 = F_makeAlias(m, int32(_a_F_parseCheckAggregates_4), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	if v426 == int32(0) {
		v522 = v381
		v525 = v440
		v528 = v440
		v533 = v440
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v536 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v445)+124)) = uint16(v536)
	*(*int32)(unsafe.Add(mBase, uint32(v445)+120)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v453
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v541 = F_lappend(m, v540, v445)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L3
	} else {
		goto L100
	}
L81:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	if v457 <= int32(0) {
		v522 = v381
		v525 = v440
		v528 = v440
		v533 = v440
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v464 = v381
	v467 = v440
	v468 = v440
	v470 = v440
	v475 = v440
	goto L83
L83:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v426)+12))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478+v468<<(uint(int32(2))%32))))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+12))
	if v483 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v522 = v502
	v525 = v497
	v528 = v512
	v533 = v507
	goto L80
L85:
	;
	v484 = F_pstrdup(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L3
	} else {
		goto L88
	}
L86:
	;
	v487 = int32(_a_F_parseCheckAggregates_5)
	goto L87
L87:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v489 = F_makeString(m, v487)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L3
	} else {
		goto L89
	}
L88:
	;
	v487 = v484
	goto L87
L89:
	;
	v491 = F_lappend(m, v488, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = v491
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v495 = F_copyObjectImpl(m, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v497 = F_lappend(m, v467, v495)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v500 = F_exprType(m, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	v502 = F_lappend_oid(m, v464, v500)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v505 = F_exprTypmod(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	v507 = F_lappend_int(m, v475, v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L96
	}
L96:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v510 = F_exprCollation(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	v512 = F_lappend_oid(m, v470, v510)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L3
	} else {
		goto L98
	}
L98:
	;
	v515 = v468 + int32(1)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	if v515 < v516 {
		v464 = v502
		v467 = v497
		v468 = v515
		v470 = v512
		v475 = v507
		goto L83
	} else {
		goto L99
	}
L99:
	;
	goto L84
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v541
	if v541 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v541)+4))
	v546 = v544
	goto L103
L102:
	;
	v546 = int32(0)
	goto L103
L103:
	;
	v547 = F_buildNSItemFromLists(m, v445, v546, v522, v533, v528)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v547
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v551 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v551)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v550
	v558 = v426
	goto L62
L105:
	;
	v575 = F_flatten_join_alias_for_parser(m, l1, v558, int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L3
	} else {
		goto L108
	}
L106:
	;
	v577 = v558
	goto L107
L107:
	;
	if v577 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v577 = v575
	goto L107
L109:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v689 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v689
	v691 = int32(1)
	v692 = v675 & v691
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v692)
	v694 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v694
	*(*int64)(unsafe.Add(mBase, uint32(v21)+20)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v577
	v700 = v368 & v691
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v700)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v694)
	v708 = F_finalize_grouping_exprs_walker(m, v688, v21+int32(4))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L3
	} else {
		goto L140
	}
L110:
	;
	v675 = v668
	v676 = int32(0)
	goto L109
L111:
	;
	v668 = int32(0)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v581 = int32(0)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v582 <= v581 {
		v668 = v581
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v585 = int32(0)
	v590 = v585
	v592 = v581
	v593 = v585
	goto L115
L115:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v577)+12))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v605+v590<<(uint(int32(2))%32))))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	if v611 != int32(6) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v675 = v662
	v676 = v663
	goto L109
L117:
	;
	v665 = v590 + int32(1)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v577)+4))
	if v665 < v666 {
		v590 = v665
		v592 = v662
		v593 = v663
		goto L115
	} else {
		goto L139
	}
L118:
	;
	v662 = int32(1)
	v663 = v593
	goto L117
L119:
	;
	goto L120
L120:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v615 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v609)+16))
	v617 = int32(0)
	if v255 == v617 {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v659 = v610
	goto L123
L123:
	;
	v660 = F_lappend(m, v593, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L138
	}
L124:
	;
	if v655 == int32(0) {
		v662 = v592
		v663 = v593
		goto L117
	} else {
		goto L137
	}
L125:
	;
	v655 = int32(0)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v623 <= int32(0) {
		v649 = v617
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v655 = v649
	goto L124
L129:
	;
	v626 = int32(0)
	if v626 < v623 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v629 = v623
	goto L132
L131:
	;
	v629 = v626
	goto L132
L132:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v632 = int32(0)
	goto L133
L133:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v630+v632<<(uint(int32(2))%32))))
	v641 = base.B2i32(v640 == v616)
	if v640 == v616 {
		v649 = v641
		goto L128
	} else {
		goto L135
	}
L134:
	;
	v649 = v641
	goto L128
L135:
	;
	v643 = v632 + int32(1)
	if v643 != v629 {
		v632 = v643
		goto L133
	} else {
		goto L136
	}
L136:
	;
	goto L134
L137:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	v659 = v658
	goto L123
L138:
	;
	v662 = v592
	v663 = v660
	goto L117
L139:
	;
	goto L116
L140:
	;
	if v700 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v711 = F_flatten_join_alias_for_parser(m, l1, v688, int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L3
	} else {
		goto L144
	}
L142:
	;
	v713 = v688
	goto L143
L143:
	;
	v714 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v714)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v714
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v692)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v714
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v577
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v714)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21
	v730 = v21 + int32(4)
	v731 = F_substitute_grouped_columns_mutator(m, v713, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L3
	} else {
		goto L145
	}
L144:
	;
	v713 = v711
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v731
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v735 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+36)) = v735
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v692)
	v738 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v738
	*(*int64)(unsafe.Add(mBase, uint32(v21)+20)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v577
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v700)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v738)
	v748 = F_finalize_grouping_exprs_walker(m, v734, v730)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L146
	}
L146:
	;
	if v700 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v751 = F_flatten_join_alias_for_parser(m, l1, v734, int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L3
	} else {
		goto L150
	}
L148:
	;
	v753 = v734
	goto L149
L149:
	;
	v754 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)) = uint8(v754)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v754
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+32)) = uint8(v692)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v754
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v676
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v577
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v754)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l0
	v771 = F_substitute_grouped_columns_mutator(m, v753, v21+int32(4))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L3
	} else {
		goto L151
	}
L150:
	;
	v753 = v751
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v771
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v774&v373&int32(1) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L3
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	m.G0 = v21 + int32(48)
	return
L155:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L3
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(_a_F_parseCheckAggregates_6), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v790 = F_locate_agg_of_level(m, l1, int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	F_parser_errposition(m, l0, v790)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L3
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_parseCheckAggregates_2), int32(1334), int32(_a_F_parseCheckAggregates_3))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_datetime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v115 int32
	_ = v115
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int64
	_ = v246
	var v247 int64
	_ = v247
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v411 int32
	_ = v411
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v414 int64
	_ = v414
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v424 int64
	_ = v424
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int64
	_ = v460
	var v461 int64
	_ = v461
	var v462 int64
	_ = v462
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v468 int64
	_ = v468
	var v469 int64
	_ = v469
	var v472 int64
	_ = v472
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v561 int64
	_ = v561
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v18 = v13 + int32(52)
	v27 = F_do_to_timestamp(m, l0, l1, int32(100), int32(1), v18, v13+int32(40), v13+int32(44), v13+int32(36), v13+int32(32), l5)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return int64(0)
	} else {
		if v27 == int32(0) {
			v561 = v10
			m.G0 = v13 + int32(96)
			return v561
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
			if v33 != 0 {
				v35 = v33
			} else {
				v35 = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
			v39 = v37 & int32(2)
			if v37&int32(1) != 0 {
				v43 = v37 & int32(4)
				if v39 != 0 {
					if v43 != 0 {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)))
						if v44 == int32(1) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l4))) = v47
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
							v51 = v13 + int32(24)
							v58 = m.G0
							v60 = v58 - int32(16)
							m.G0 = v60
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							if v62 <= int32(-4713) {
								if v62 != int32(-4713) {
									*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
									v139 = int32(-1)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
									if int32(10) < v67 {
										v78 = v67
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
										v80 = F_date2j(m, v62, v78, v79)
										mBase = m.M
										v83 = base.I64_extend_i32_s(v80 - int32(_a_F_parse_datetime_0))
										v84 = int64(63)
										F___multi3(m, v60, v83, v83>>(uint(v84)%64), int64(86400000000), int64(0))
										mBase = m.M
										v89 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
										v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										if v89 != v90>>(uint(v84)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
											v139 = int32(-1)
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
											v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
											v98 = int32(60)
											v107 = base.I64_extend_i32_s(v49) + base.I64_extend_i32_s(v95+(v96+v97*v98)*v98)*int64(1000000)
											v108 = v90 + v107
											*(*int64)(unsafe.Add(mBase, uint32(v51))) = v108
											if base.B2i32(v107 < int64(0))^base.B2i32(v108 < v90) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
												v139 = int32(-1)
											} else {
												if l4 != 0 {
													v115 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
													v120 = base.I64_extend_i32_s(int32(0)-v115)*int64(-1000000) + v108
													*(*int64)(unsafe.Add(mBase, uint32(v51))) = v120
													v122 = v120
												} else {
													v122 = v108
												}
												if base.Ui64(v122+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v139 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
													v139 = int32(-1)
												}
											}
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
										v139 = int32(-1)
									}
								}
							} else {
								if v62 <= int32(_a_F_parse_datetime_1) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
									v78 = v72
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
									v80 = F_date2j(m, v62, v78, v79)
									mBase = m.M
									v83 = base.I64_extend_i32_s(v80 - int32(_a_F_parse_datetime_0))
									v84 = int64(63)
									F___multi3(m, v60, v83, v83>>(uint(v84)%64), int64(86400000000), int64(0))
									mBase = m.M
									v89 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
									v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									if v89 != v90>>(uint(v84)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
										v139 = int32(-1)
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
										v98 = int32(60)
										v107 = base.I64_extend_i32_s(v49) + base.I64_extend_i32_s(v95+(v96+v97*v98)*v98)*int64(1000000)
										v108 = v90 + v107
										*(*int64)(unsafe.Add(mBase, uint32(v51))) = v108
										if base.B2i32(v107 < int64(0))^base.B2i32(v108 < v90) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
											v139 = int32(-1)
										} else {
											if l4 != 0 {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
												v120 = base.I64_extend_i32_s(int32(0)-v115)*int64(-1000000) + v108
												*(*int64)(unsafe.Add(mBase, uint32(v51))) = v120
												v122 = v120
											} else {
												v122 = v108
											}
											if base.Ui64(v122+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v139 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
												v139 = int32(-1)
											}
										}
									}
								} else {
									if v62 != int32(_a_F_parse_datetime_2) {
										*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
										v139 = int32(-1)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
										if int32(5) < v75 {
											*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
											v139 = int32(-1)
										} else {
											v78 = v75
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
											v80 = F_date2j(m, v62, v78, v79)
											mBase = m.M
											v83 = base.I64_extend_i32_s(v80 - int32(_a_F_parse_datetime_0))
											v84 = int64(63)
											F___multi3(m, v60, v83, v83>>(uint(v84)%64), int64(86400000000), int64(0))
											mBase = m.M
											v89 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
											v90 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
											if v89 != v90>>(uint(v84)%64) {
												*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
												v139 = int32(-1)
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
												v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
												v98 = int32(60)
												v107 = base.I64_extend_i32_s(v49) + base.I64_extend_i32_s(v95+(v96+v97*v98)*v98)*int64(1000000)
												v108 = v90 + v107
												*(*int64)(unsafe.Add(mBase, uint32(v51))) = v108
												if base.B2i32(v107 < int64(0))^base.B2i32(v108 < v90) != 0 {
													*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
													v139 = int32(-1)
												} else {
													if l4 != 0 {
														v115 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
														v120 = base.I64_extend_i32_s(int32(0)-v115)*int64(-1000000) + v108
														*(*int64)(unsafe.Add(mBase, uint32(v51))) = v120
														v122 = v120
													} else {
														v122 = v108
													}
													if base.Ui64(v122+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
														v139 = int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
														v139 = int32(-1)
													}
												}
											}
										}
									}
								}
							}
							m.G0 = v60 + int32(16)
							if v139 == int32(0) {
								v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								v180 = F_AdjustTimestampForTypmod(m, v13+int32(24), v179, l5)
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1184)
									v184 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
									v561 = v184
									m.G0 = v13 + int32(96)
									return v561
								}
							} else {
								v145 = F_errsave_start(m, l5)
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int64(0)
								} else {
									if v145 == int32(0) {
										v561 = v10
										m.G0 = v13 + int32(96)
										return v561
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_3), int32(0))
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_5), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return int64(0)
												} else {
													v561 = v10
													m.G0 = v13 + int32(96)
													return v561
												}
											}
										}
									}
								}
							}
						} else {
							v161 = F_errsave_start(m, l5)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int64(0)
							} else {
								if v161 == int32(0) {
									v561 = v10
									m.G0 = v13 + int32(96)
									return v561
								} else {
									F_errcode(m, int32(117440642))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_7), int32(0))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_8), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v176 = m.ExcPending
											if v176 != 0 {
												return int64(0)
											} else {
												v561 = v10
												m.G0 = v13 + int32(96)
												return v561
											}
										}
									}
								}
							}
						}
					} else {
						v186 = v13 + int32(52)
						v187 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
						v190 = v13 + int32(24)
						v197 = m.G0
						v199 = v197 - int32(16)
						m.G0 = v199
						v201 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						if v201 <= int32(-4713) {
							if v201 != int32(-4713) {
								*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
								v278 = int32(-1)
							} else {
								v206 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
								if int32(10) < v206 {
									v217 = v206
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
									v219 = F_date2j(m, v201, v217, v218)
									mBase = m.M
									v222 = base.I64_extend_i32_s(v219 - int32(_a_F_parse_datetime_0))
									v223 = int64(63)
									F___multi3(m, v199, v222, v222>>(uint(v223)%64), int64(86400000000), int64(0))
									mBase = m.M
									v228 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
									v229 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
									if v228 != v229>>(uint(v223)%64) {
										*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
										v278 = int32(-1)
									} else {
										v234 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
										v235 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
										v236 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
										v237 = int32(60)
										v246 = base.I64_extend_i32_s(v187) + base.I64_extend_i32_s(v234+(v235+v236*v237)*v237)*int64(1000000)
										v247 = v229 + v246
										*(*int64)(unsafe.Add(mBase, uint32(v190))) = v247
										if base.B2i32(v246 < int64(0))^base.B2i32(v247 < v229) != 0 {
											*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
											v278 = int32(-1)
										} else {
											if base.Ui64(v247+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v278 = int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
												v278 = int32(-1)
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
									v278 = int32(-1)
								}
							}
						} else {
							if v201 <= int32(_a_F_parse_datetime_1) {
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
								v217 = v211
								v218 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
								v219 = F_date2j(m, v201, v217, v218)
								mBase = m.M
								v222 = base.I64_extend_i32_s(v219 - int32(_a_F_parse_datetime_0))
								v223 = int64(63)
								F___multi3(m, v199, v222, v222>>(uint(v223)%64), int64(86400000000), int64(0))
								mBase = m.M
								v228 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
								v229 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
								if v228 != v229>>(uint(v223)%64) {
									*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
									v278 = int32(-1)
								} else {
									v234 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
									v235 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
									v236 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
									v237 = int32(60)
									v246 = base.I64_extend_i32_s(v187) + base.I64_extend_i32_s(v234+(v235+v236*v237)*v237)*int64(1000000)
									v247 = v229 + v246
									*(*int64)(unsafe.Add(mBase, uint32(v190))) = v247
									if base.B2i32(v246 < int64(0))^base.B2i32(v247 < v229) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
										v278 = int32(-1)
									} else {
										if base.Ui64(v247+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v278 = int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
											v278 = int32(-1)
										}
									}
								}
							} else {
								if v201 != int32(_a_F_parse_datetime_2) {
									*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
									v278 = int32(-1)
								} else {
									v214 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
									if int32(5) < v214 {
										*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
										v278 = int32(-1)
									} else {
										v217 = v214
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
										v219 = F_date2j(m, v201, v217, v218)
										mBase = m.M
										v222 = base.I64_extend_i32_s(v219 - int32(_a_F_parse_datetime_0))
										v223 = int64(63)
										F___multi3(m, v199, v222, v222>>(uint(v223)%64), int64(86400000000), int64(0))
										mBase = m.M
										v228 = *(*int64)(unsafe.Add(mBase, uint32(v199)+8))
										v229 = *(*int64)(unsafe.Add(mBase, uint32(v199)))
										if v228 != v229>>(uint(v223)%64) {
											*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
											v278 = int32(-1)
										} else {
											v234 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
											v235 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
											v236 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
											v237 = int32(60)
											v246 = base.I64_extend_i32_s(v187) + base.I64_extend_i32_s(v234+(v235+v236*v237)*v237)*int64(1000000)
											v247 = v229 + v246
											*(*int64)(unsafe.Add(mBase, uint32(v190))) = v247
											if base.B2i32(v246 < int64(0))^base.B2i32(v247 < v229) != 0 {
												*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
												v278 = int32(-1)
											} else {
												if base.Ui64(v247+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
													v278 = int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(0)
													v278 = int32(-1)
												}
											}
										}
									}
								}
							}
						}
						m.G0 = v199 + int32(16)
						if v278 != 0 {
							v282 = F_errsave_start(m, l5)
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return int64(0)
							} else {
								if v282 == int32(0) {
									v561 = v10
									m.G0 = v13 + int32(96)
									return v561
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v288 = m.ExcPending
									if v288 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_parse_datetime_9), int32(0))
										mBase = m.M
										v292 = m.ExcPending
										if v292 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_10), int32(_a_F_parse_datetime_6))
											mBase = m.M
											v297 = m.ExcPending
											if v297 != 0 {
												return int64(0)
											} else {
												v561 = v10
												m.G0 = v13 + int32(96)
												return v561
											}
										}
									}
								}
							}
						} else {
							v300 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v301 = F_AdjustTimestampForTypmod(m, v13+int32(24), v300, l5)
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1114)
								v305 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
								v561 = v305
								m.G0 = v13 + int32(96)
								return v561
							}
						}
					}
				} else {
					if v43 != 0 {
						v306 = F_errsave_start(m, l5)
						mBase = m.M
						v307 = m.ExcPending
						if v307 != 0 {
							return int64(0)
						} else {
							if v306 == int32(0) {
								v561 = v10
								m.G0 = v13 + int32(96)
								return v561
							} else {
								F_errcode(m, int32(117440642))
								mBase = m.M
								v312 = m.ExcPending
								if v312 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_parse_datetime_11), int32(0))
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int64(0)
									} else {
										F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_12), int32(_a_F_parse_datetime_6))
										mBase = m.M
										v321 = m.ExcPending
										if v321 != 0 {
											return int64(0)
										} else {
											v561 = v10
											m.G0 = v13 + int32(96)
											return v561
										}
									}
								}
							}
						}
					} else {
						v322 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
						if v322 <= int32(-4713) {
							if v322 != int32(-4713) {
								v339 = F_errsave_start(m, l5)
								mBase = m.M
								v340 = m.ExcPending
								if v340 != 0 {
									return int64(0)
								} else {
									if v339 == int32(0) {
										v561 = v10
										m.G0 = v13 + int32(96)
										return v561
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v345 = m.ExcPending
										if v345 != 0 {
											return int64(0)
										} else {
											v346 = F_text_to_cstring(m, l0)
											mBase = m.M
											v347 = m.ExcPending
											if v347 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v346
												F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
												mBase = m.M
												v353 = m.ExcPending
												if v353 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
													mBase = m.M
													v358 = m.ExcPending
													if v358 != 0 {
														return int64(0)
													} else {
														v561 = v10
														m.G0 = v13 + int32(96)
														return v561
													}
												}
											}
										}
									}
								}
							} else {
								v327 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
								if v327 <= int32(10) {
									v339 = F_errsave_start(m, l5)
									mBase = m.M
									v340 = m.ExcPending
									if v340 != 0 {
										return int64(0)
									} else {
										if v339 == int32(0) {
											v561 = v10
											m.G0 = v13 + int32(96)
											return v561
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v345 = m.ExcPending
											if v345 != 0 {
												return int64(0)
											} else {
												v346 = F_text_to_cstring(m, l0)
												mBase = m.M
												v347 = m.ExcPending
												if v347 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v346
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
													mBase = m.M
													v353 = m.ExcPending
													if v353 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v358 = m.ExcPending
														if v358 != 0 {
															return int64(0)
														} else {
															v561 = v10
															m.G0 = v13 + int32(96)
															return v561
														}
													}
												}
											}
										}
									}
								} else {
									v497 = v327
									v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
									v503 = base.B2i32(int32(2) < v497)
									if int32(2) < v497 {
										v504 = int32(_a_F_parse_datetime_15)
									} else {
										v504 = int32(_a_F_parse_datetime_16)
									}
									v505 = v504 + v322
									v510 = base.I32_div_s(v505, int32(4))
									v513 = base.I32_div_s(v505, int32(-100))
									v516 = base.I32_div_s(v505, int32(400))
									if int32(2) < v497 {
										v520 = int32(1)
									} else {
										v520 = int32(13)
									}
									v525 = base.I32_div_s((v520+v497)*int32(_a_F_parse_datetime_17), int32(256))
									v528 = v498 + v505*int32(365) + v510 + v513 + v516 + v525 - int32(_a_F_parse_datetime_18)
									if base.Ui32(int32(2147483494)) <= base.Ui32(v528) {
										v531 = F_errsave_start(m, l5)
										mBase = m.M
										v532 = m.ExcPending
										if v532 != 0 {
											return int64(0)
										} else {
											if v531 == int32(0) {
												v561 = v10
												m.G0 = v13 + int32(96)
												return v561
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v537 = m.ExcPending
												if v537 != 0 {
													return int64(0)
												} else {
													v538 = F_text_to_cstring(m, l0)
													mBase = m.M
													v539 = m.ExcPending
													if v539 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13))) = v538
														F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
														mBase = m.M
														v543 = m.ExcPending
														if v543 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v548 = m.ExcPending
															if v548 != 0 {
																return int64(0)
															} else {
																v561 = v10
																m.G0 = v13 + int32(96)
																return v561
															}
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
										v561 = base.I64_extend_i32_s(v528 - int32(_a_F_parse_datetime_0))
										m.G0 = v13 + int32(96)
										return v561
									}
								}
							}
						} else {
							if v322 <= int32(_a_F_parse_datetime_1) {
								v332 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
								v497 = v332
								v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
								v503 = base.B2i32(int32(2) < v497)
								if int32(2) < v497 {
									v504 = int32(_a_F_parse_datetime_15)
								} else {
									v504 = int32(_a_F_parse_datetime_16)
								}
								v505 = v504 + v322
								v510 = base.I32_div_s(v505, int32(4))
								v513 = base.I32_div_s(v505, int32(-100))
								v516 = base.I32_div_s(v505, int32(400))
								if int32(2) < v497 {
									v520 = int32(1)
								} else {
									v520 = int32(13)
								}
								v525 = base.I32_div_s((v520+v497)*int32(_a_F_parse_datetime_17), int32(256))
								v528 = v498 + v505*int32(365) + v510 + v513 + v516 + v525 - int32(_a_F_parse_datetime_18)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v528) {
									v531 = F_errsave_start(m, l5)
									mBase = m.M
									v532 = m.ExcPending
									if v532 != 0 {
										return int64(0)
									} else {
										if v531 == int32(0) {
											v561 = v10
											m.G0 = v13 + int32(96)
											return v561
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v537 = m.ExcPending
											if v537 != 0 {
												return int64(0)
											} else {
												v538 = F_text_to_cstring(m, l0)
												mBase = m.M
												v539 = m.ExcPending
												if v539 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13))) = v538
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
													mBase = m.M
													v543 = m.ExcPending
													if v543 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v548 = m.ExcPending
														if v548 != 0 {
															return int64(0)
														} else {
															v561 = v10
															m.G0 = v13 + int32(96)
															return v561
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
									v561 = base.I64_extend_i32_s(v528 - int32(_a_F_parse_datetime_0))
									m.G0 = v13 + int32(96)
									return v561
								}
							} else {
								if v322 != int32(_a_F_parse_datetime_2) {
									v339 = F_errsave_start(m, l5)
									mBase = m.M
									v340 = m.ExcPending
									if v340 != 0 {
										return int64(0)
									} else {
										if v339 == int32(0) {
											v561 = v10
											m.G0 = v13 + int32(96)
											return v561
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v345 = m.ExcPending
											if v345 != 0 {
												return int64(0)
											} else {
												v346 = F_text_to_cstring(m, l0)
												mBase = m.M
												v347 = m.ExcPending
												if v347 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v346
													F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
													mBase = m.M
													v353 = m.ExcPending
													if v353 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
														mBase = m.M
														v358 = m.ExcPending
														if v358 != 0 {
															return int64(0)
														} else {
															v561 = v10
															m.G0 = v13 + int32(96)
															return v561
														}
													}
												}
											}
										}
									}
								} else {
									v335 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
									if v335 < int32(6) {
										v497 = v335
										v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
										v503 = base.B2i32(int32(2) < v497)
										if int32(2) < v497 {
											v504 = int32(_a_F_parse_datetime_15)
										} else {
											v504 = int32(_a_F_parse_datetime_16)
										}
										v505 = v504 + v322
										v510 = base.I32_div_s(v505, int32(4))
										v513 = base.I32_div_s(v505, int32(-100))
										v516 = base.I32_div_s(v505, int32(400))
										if int32(2) < v497 {
											v520 = int32(1)
										} else {
											v520 = int32(13)
										}
										v525 = base.I32_div_s((v520+v497)*int32(_a_F_parse_datetime_17), int32(256))
										v528 = v498 + v505*int32(365) + v510 + v513 + v516 + v525 - int32(_a_F_parse_datetime_18)
										if base.Ui32(int32(2147483494)) <= base.Ui32(v528) {
											v531 = F_errsave_start(m, l5)
											mBase = m.M
											v532 = m.ExcPending
											if v532 != 0 {
												return int64(0)
											} else {
												if v531 == int32(0) {
													v561 = v10
													m.G0 = v13 + int32(96)
													return v561
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v537 = m.ExcPending
													if v537 != 0 {
														return int64(0)
													} else {
														v538 = F_text_to_cstring(m, l0)
														mBase = m.M
														v539 = m.ExcPending
														if v539 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13))) = v538
															F_errmsg(m, int32(_a_F_parse_datetime_13), v13)
															mBase = m.M
															v543 = m.ExcPending
															if v543 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_19), int32(_a_F_parse_datetime_6))
																mBase = m.M
																v548 = m.ExcPending
																if v548 != 0 {
																	return int64(0)
																} else {
																	v561 = v10
																	m.G0 = v13 + int32(96)
																	return v561
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1082)
											v561 = base.I64_extend_i32_s(v528 - int32(_a_F_parse_datetime_0))
											m.G0 = v13 + int32(96)
											return v561
										}
									} else {
										v339 = F_errsave_start(m, l5)
										mBase = m.M
										v340 = m.ExcPending
										if v340 != 0 {
											return int64(0)
										} else {
											if v339 == int32(0) {
												v561 = v10
												m.G0 = v13 + int32(96)
												return v561
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v345 = m.ExcPending
												if v345 != 0 {
													return int64(0)
												} else {
													v346 = F_text_to_cstring(m, l0)
													mBase = m.M
													v347 = m.ExcPending
													if v347 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v346
														F_errmsg(m, int32(_a_F_parse_datetime_13), v13+int32(16))
														mBase = m.M
														v353 = m.ExcPending
														if v353 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_14), int32(_a_F_parse_datetime_6))
															mBase = m.M
															v358 = m.ExcPending
															if v358 != 0 {
																return int64(0)
															} else {
																v561 = v10
																m.G0 = v13 + int32(96)
																return v561
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
				}
			} else {
				if v39 != 0 {
					if v37&int32(4) != 0 {
						v362 = F_palloc(m, int32(16))
						mBase = m.M
						v363 = m.ExcPending
						if v363 != 0 {
							return int64(0)
						} else {
							v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+44)))
							if v364 == int32(1) {
								v367 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(l4))) = v367
								v369 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
								v371 = v13 + int32(52)
								v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
								v373 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v362)+8)) = v367
								v377 = int32(60)
								*(*int64)(unsafe.Add(mBase, uint32(v362))) = base.I64_extend_i32_s(v369) + base.I64_extend_i32_s(v372+(v373+v374*v377)*v377)*int64(1000000)
								v404 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
								if base.Ui32(v404) <= base.Ui32(int32(6)) {
									v411 = v404 << (uint(int32(3)) % 32)
									v412 = *(*int64)(unsafe.Add(mBase, uint32(v411)+uint32(_c_F_parse_datetime[0])))
									v413 = *(*int64)(unsafe.Add(mBase, uint32(v411)+uint32(_c_F_parse_datetime[1])))
									v414 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
									if int64(0) <= v414 {
										v417 = v413 + v414
										v418 = base.I64_rem_s(v417, v412)
										v424 = v417 - v418
									} else {
										v420 = v413 - v414
										v421 = base.I64_rem_s(v420, v412)
										v424 = v421 - v420
									}
									*(*int64)(unsafe.Add(mBase, uint32(v362))) = v424
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1266)
								v561 = base.I64_extend_i32_u(v362)
								m.G0 = v13 + int32(96)
								return v561
							} else {
								v388 = F_errsave_start(m, l5)
								mBase = m.M
								v389 = m.ExcPending
								if v389 != 0 {
									return int64(0)
								} else {
									if v388 == int32(0) {
										v561 = v10
										m.G0 = v13 + int32(96)
										return v561
									} else {
										F_errcode(m, int32(117440642))
										mBase = m.M
										v394 = m.ExcPending
										if v394 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_datetime_20), int32(0))
											mBase = m.M
											v398 = m.ExcPending
											if v398 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_21), int32(_a_F_parse_datetime_6))
												mBase = m.M
												v403 = m.ExcPending
												if v403 != 0 {
													return int64(0)
												} else {
													v561 = v10
													m.G0 = v13 + int32(96)
													return v561
												}
											}
										}
									}
								}
							}
						}
					} else {
						v434 = v13 + int32(24)
						v435 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+40)))
						v437 = v13 + int32(52)
						v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
						v439 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
						v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
						v441 = int32(60)
						*(*int64)(unsafe.Add(mBase, uint32(v434))) = v435 + base.I64_extend_i32_s(v438+(v439+v440*v441)*v441)*int64(1000000)
						v452 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						if base.Ui32(v452) <= base.Ui32(int32(6)) {
							v459 = v452 << (uint(int32(3)) % 32)
							v460 = *(*int64)(unsafe.Add(mBase, uint32(v459)+uint32(_c_F_parse_datetime[0])))
							v461 = *(*int64)(unsafe.Add(mBase, uint32(v459)+uint32(_c_F_parse_datetime[1])))
							v462 = *(*int64)(unsafe.Add(mBase, uint32(v434)))
							if int64(0) <= v462 {
								v465 = v461 + v462
								v466 = base.I64_rem_s(v465, v460)
								v472 = v465 - v466
							} else {
								v468 = v461 - v462
								v469 = base.I64_rem_s(v468, v460)
								v472 = v469 - v468
							}
							*(*int64)(unsafe.Add(mBase, uint32(v434))) = v472
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1083)
						v480 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
						v561 = v480
						m.G0 = v13 + int32(96)
						return v561
					}
				} else {
					v481 = F_errsave_start(m, l5)
					mBase = m.M
					v482 = m.ExcPending
					if v482 != 0 {
						return int64(0)
					} else {
						if v481 == int32(0) {
							v561 = v10
							m.G0 = v13 + int32(96)
							return v561
						} else {
							F_errcode(m, int32(117440642))
							mBase = m.M
							v487 = m.ExcPending
							if v487 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_parse_datetime_22), int32(0))
								mBase = m.M
								v491 = m.ExcPending
								if v491 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, l5, int32(_a_F_parse_datetime_4), int32(_a_F_parse_datetime_23), int32(_a_F_parse_datetime_6))
									mBase = m.M
									v496 = m.ExcPending
									if v496 != 0 {
										return int64(0)
									} else {
										v561 = v10
										m.G0 = v13 + int32(96)
										return v561
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
func F_parse_ident_line(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v18 != 0 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		v21 = v19
	} else {
		v21 = int32(0)
	}
	v22 = int32(16)
	v23 = l0 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = F_palloc0(m, v22)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v17
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		if int32(2) <= v32 {
			v35 = int32(0)
			v37 = F_errstart(m, l1, v35)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 != 0 {
					F_errcode(m, int32(22))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = v17
								F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2576), int32(_a_F_parse_ident_line_3))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v60 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23))) = v60
											v262 = v35
											m.G0 = v15 + int32(80)
											return v262
										}
									}
								}
							}
						}
					}
				} else {
					v60 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v23))) = v60
						v262 = v35
						m.G0 = v15 + int32(80)
						return v262
					}
				}
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
			v66 = F_pstrdup(m, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v66
				v70 = v21 + int32(4)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
				if base.Ui32(v72+v73<<(uint(int32(2))%32)) <= base.Ui32(v70) {
					v78 = int32(0)
					v80 = F_errstart(m, l1, v78)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						if v80 != 0 {
							F_errcode(m, int32(22))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_set_errcontext_domain(m, int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v24
										*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v17
										F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15-int32(-64))
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2582), int32(_a_F_parse_ident_line_3))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v105 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v23))) = v105
													v262 = v78
													m.G0 = v15 + int32(80)
													return v262
												}
											}
										}
									}
								}
							}
						} else {
							v105 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23))) = v105
								v262 = v78
								m.G0 = v15 + int32(80)
								return v262
							}
						}
					}
				} else {
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					if int32(2) <= v109 {
						v112 = int32(0)
						v114 = F_errstart(m, l1, v112)
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							if v114 != 0 {
								F_errcode(m, int32(22))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										F_set_errcontext_domain(m, int32(0))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v24
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v17
											F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(16))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2584), int32(_a_F_parse_ident_line_3))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													v139 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v23))) = v139
														v262 = v112
														m.G0 = v15 + int32(80)
														return v262
													}
												}
											}
										}
									}
								}
							} else {
								v139 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23))) = v139
									v262 = v112
									m.G0 = v15 + int32(80)
									return v262
								}
							}
						}
					} else {
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
						v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+4)))
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						v146 = F_strlen(m, v145)
						mBase = m.M
						v149 = F_palloc0(m, v146+int32(13))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v149)+4)) = uint8(v144)
							v155 = v149 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v149))) = v155
							v158 = v146 + int32(1)
							if v158 != 0 {
								base.MemoryCopy(m, v155, v145, v158)
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v149
							v162 = v21 + int32(8)
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
							v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
							if base.Ui32(v164+v165<<(uint(int32(2))%32)) <= base.Ui32(v162) {
								v170 = int32(0)
								v172 = F_errstart(m, l1, v170)
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return int32(0)
								} else {
									if v172 != 0 {
										F_errcode(m, int32(22))
										mBase = m.M
										v176 = m.ExcPending
										if v176 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_parse_ident_line_4), int32(0))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												F_set_errcontext_domain(m, int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v24
													*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v17
													F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(48))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2592), int32(_a_F_parse_ident_line_3))
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															v197 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
															mBase = m.M
															v198 = m.ExcPending
															if v198 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v23))) = v197
																v262 = v170
																m.G0 = v15 + int32(80)
																return v262
															}
														}
													}
												}
											}
										}
									} else {
										v197 = F_pstrdup(m, int32(_a_F_parse_ident_line_4))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v23))) = v197
											v262 = v170
											m.G0 = v15 + int32(80)
											return v262
										}
									}
								}
							} else {
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
								if int32(2) <= v201 {
									v204 = int32(0)
									v206 = F_errstart(m, l1, v204)
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int32(0)
									} else {
										if v206 != 0 {
											F_errcode(m, int32(22))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_parse_ident_line_0), int32(0))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													F_set_errcontext_domain(m, int32(0))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v24
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v17
														F_errcontext_msg(m, int32(_a_F_parse_ident_line_1), v15+int32(32))
														mBase = m.M
														v224 = m.ExcPending
														if v224 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_parse_ident_line_2), int32(2594), int32(_a_F_parse_ident_line_3))
															mBase = m.M
															v229 = m.ExcPending
															if v229 != 0 {
																return int32(0)
															} else {
																v231 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
																mBase = m.M
																v232 = m.ExcPending
																if v232 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v231
																	v262 = v204
																	m.G0 = v15 + int32(80)
																	return v262
																}
															}
														}
													}
												}
											}
										} else {
											v231 = F_pstrdup(m, int32(_a_F_parse_ident_line_0))
											mBase = m.M
											v232 = m.ExcPending
											if v232 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v23))) = v231
												v262 = v204
												m.G0 = v15 + int32(80)
												return v262
											}
										}
									}
								} else {
									v234 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
									v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
									v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+4)))
									v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
									v239 = F_strlen(m, v238)
									mBase = m.M
									v242 = F_palloc0(m, v239+int32(13))
									mBase = m.M
									v243 = m.ExcPending
									if v243 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v242)+4)) = uint8(v236)
										v248 = v242 + int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(v242))) = v248
										v251 = v239 + int32(1)
										if v251 != 0 {
											base.MemoryCopy(m, v248, v238, v251)
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v242
										v254 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
										v255 = F_regcomp_auth_token(m, v254, v24, v17, v23, l1)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											if v255 != 0 {
												v262 = int32(0)
												m.G0 = v15 + int32(80)
												return v262
											} else {
												v258 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
												v259 = F_regcomp_auth_token(m, v258, v24, v17, v23, l1)
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return int32(0)
												} else {
													if v259 != 0 {
														v261 = int32(0)
													} else {
														v261 = v26
													}
													v262 = v261
													m.G0 = v15 + int32(80)
													return v262
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
	}
}
func F_pathkeys_count_contained_in(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	if l0 == l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v76
	return int32(1)
L2:
	;
	if l0 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	return int32(1)
L7:
	;
	goto L8
L8:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v23
	return v23
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = int32(0)
	if v28 < v27 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = v27
	goto L14
L13:
	;
	v31 = v28
	goto L14
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = int32(0)
	goto L15
L15:
	;
	if v37 < v32 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v44 + v37<<(uint(int32(2))%32)
	goto L19
L18:
	;
	v48 = int32(0)
	goto L19
L19:
	;
	if v37 == v31 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
	return base.B2i32(v48 == int32(0))
L21:
	;
	goto L22
L22:
	;
	v55 = base.B2i32(v48 == int32(0))
	if v48 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return v55
L24:
	;
	goto L25
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v60 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return v55
L27:
	;
	goto L28
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60+v37<<(uint(int32(2))%32))))
	if v65 != v69 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v37
	return int32(0)
L30:
	;
	v37 = v37 + int32(1)
	goto L15
}
func F_pgarch_waken_stop(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_pgarch_waken_stop_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_pgarch_waken_stop_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgarch_waken_stop[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_pgl_atexit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_atexit[0]))
	if v3 <= int32(31) {
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_atexit[0])) = v3 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_c_F_pgl_atexit[1]))) = int32(1196)
	} else {
	}
	return
}
func F_pgl_freopen(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	switch l2 {
	case 0:
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[0]))
		v8 = F_freopen(m, l0, l1, v7)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[1])) = v8
			return v8
		}
	case 1:
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[2]))
		v17 = F_freopen(m, l0, l1, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[3])) = v17
			return v17
		}
	case 2:
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_freopen[4]))
		v23 = F_freopen(m, l0, l1, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = v23
			return v25
		}
	default:
		v25 = int32(0)
		return v25
	}
}
func F_pglz_decompress_datum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v217 int32
	_ = v217
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_palloc(m, v3&int32(1073741823)+int32(4))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(8)
	v13 = l0 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(base.Ui32(v14)>>(uint(int32(2))%32)) - v12
	v20 = v8 + int32(4)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = v21 & int32(1073741823)
	v25 = int32(0)
	v33 = v20 + v23
	v34 = v13 + v18
	if base.B2i32(v18 <= v25)|base.B2i32(v23 <= v25) == v25 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v217 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L4:
	;
	goto L3
L5:
	;
	goto L46
L6:
	;
	v42 = v13
	v45 = v20
	goto L9
L7:
	;
	goto L8
L8:
	;
	v191 = v13
	v194 = v20
	goto L5
L9:
	;
	v56 = v42 + int32(1)
	if base.Ui32(v34) <= base.Ui32(v56) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v191 = v176
	v194 = v179
	goto L5
L11:
	;
	if base.Ui32(v34) <= base.Ui32(v176) {
		v191 = v176
		v194 = v179
		goto L5
	} else {
		goto L44
	}
L12:
	;
	v176 = v56
	v179 = v45
	goto L11
L13:
	;
	goto L14
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v60 = v56
	v63 = v45
	v69 = v58
	v70 = int32(0)
	goto L15
L15:
	;
	if v69&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v176 = v151
	v179 = v164
	goto L11
L17:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v70))|base.B2i32(base.Ui32(v34) <= base.Ui32(v151)) != 0 {
		v176 = v151
		v179 = v164
		goto L11
	} else {
		goto L42
	}
L18:
	;
	v75 = int32(-1)
	v77 = v60 + int32(2)
	if base.Ui32(v34) < base.Ui32(v77) {
		v217 = v75
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v145)
	v147 = int32(1)
	v151 = v60 + v147
	v164 = v63 + v147
	goto L17
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v84 = v80&int32(15) + int32(3)
	if v84 != int32(18) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v93 = v84
	v94 = v77
	goto L24
L23:
	;
	if base.Ui32(v34) <= base.Ui32(v77) {
		v217 = v75
		goto L4
	} else {
		goto L25
	}
L24:
	;
	v99 = v80<<(uint(int32(4))%32)&int32(3840) | v79
	if base.B2i32(v99 == int32(0))|base.B2i32(v63-v20 < v99) != 0 {
		v217 = v75
		goto L4
	} else {
		goto L26
	}
L25:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v93 = v88 + int32(18)
	v94 = v60 + int32(3)
	goto L24
L26:
	;
	v105 = v33 - v63
	if v93 < v105 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = v93
	goto L29
L28:
	;
	v107 = v105
	goto L29
L29:
	;
	if v99 < v107 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v110 = v99
	v112 = v63
	v114 = v107
	goto L33
L31:
	;
	v130 = v99
	v132 = v63
	v134 = v107
	goto L32
L32:
	;
	if v134 != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	if v110 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v130 = v127
	v132 = v124
	v134 = v125
	goto L32
L35:
	;
	base.MemoryCopy(m, v112, v112-v110, v110)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v124 = v110 + v112
	v125 = v114 - v110
	v127 = v110 << (uint(int32(1)) % 32)
	if v127 < v125 {
		v110 = v127
		v112 = v124
		v114 = v125
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	base.MemoryCopy(m, v132, v132-v130, v134)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v151 = v94
	v164 = v132 + v134
	goto L17
L42:
	;
	v169 = int32(1)
	if base.Ui32(v164) < base.Ui32(v33) {
		v60 = v151
		v63 = v164
		v69 = int32(base.Ui32(v69&int32(254)) >> (uint(v169) % 32))
		v70 = v70 + v169
		goto L15
	} else {
		goto L43
	}
L43:
	;
	goto L16
L44:
	;
	if base.Ui32(v179) < base.Ui32(v33) {
		v42 = v176
		v45 = v179
		goto L9
	} else {
		goto L45
	}
L45:
	;
	goto L10
L46:
	;
	if base.B2i32(v191 != v34)|base.B2i32(v194 != v33) != 0 {
		v217 = int32(-1)
		goto L4
	} else {
		goto L49
	}
L48:
	;
	v217 = v194 - v20
	goto L4
L49:
	;
	goto L48
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v217<<(uint(int32(2))%32) + int32(16)
	return v8
L53:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg_internal(m, int32(_a_F_pglz_decompress_datum_0), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pglz_decompress_datum_1), int32(98), int32(_a_F_pglz_decompress_datum_2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstatginindex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstatginindex_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatginindex_1), int32(510), int32(_a_F_pgstatginindex_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v27 = F_pgstatginindex_internal(m, base.I32_wrap_i64(v3), l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				return v27
			}
		}
	}
}
func F_pgstathashindex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int64
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int64
	_ = v248
	var v256 int64
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v387 int64
	_ = v387
	var v388 int64
	_ = v388
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int64
	_ = v413
	var v421 int64
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v448 int64
	_ = v448
	var v449 int64
	_ = v449
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v491 int64
	_ = v491
	var v492 int64
	_ = v492
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v577 int64
	_ = v577
	var v578 int64
	_ = v578
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v608 int32
	_ = v608
	var v613 int64
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v623 int64
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int64
	_ = v633
	var v644 float64
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int64
	_ = v672
	var v673 int32
	_ = v673
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	v2 = int32(0)
	v20 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(160)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v20
	v38 = F_relation_open(m, v34, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L6
	} else {
		goto L117
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L6
	} else {
		goto L109
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L6
	} else {
		goto L105
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L101
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L6
	} else {
		goto L97
	}
L6:
	;
	return int64(0)
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v43 != int32(105) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
	if v46 != int32(405) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	if v49 == int32(116) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
	if v52 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+192))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+18)))
	if v56 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v62 = F__hash_getbuf(m, v38, int32(0), int32(1), int32(8))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v82 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v81)+42)))
	v83 = int64(*(*int32)(unsafe.Add(mBase, uint32(v81)+28)))
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L20
	}
L16:
	;
	if v62 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v62^int32(-1))<<(uint(int32(2))%32))))
	v81 = v73
	goto L15
L18:
	;
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[1]))
	v81 = v75 + v62<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L20:
	;
	v88 = F_RelationGetNumberOfBlocksInFork(m, v38, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v91 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = int32(1)
	v97 = int32(0)
	v102 = F_read_stream_begin_relation(m, int32(12), v91, v38, v97, int32(3), v32-int32(-64), v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v88) < base.Ui32(int32(2)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v608 = v2
	v613 = v20
	v614 = v20
	v615 = v20
	v619 = v20
	v620 = v20
	v623 = int64(0)
	goto L26
L25:
	;
	v118 = v2
	v121 = v2
	v122 = int32(1)
	v124 = v2
	v125 = v2
	v126 = v20
	v127 = v20
	v128 = v20
	goto L27
L26:
	;
	F_read_stream_end(m, v102)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L6
	} else {
		goto L87
	}
L27:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[2]))
	if v137 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v608 = v571
	v613 = v576
	v614 = v577
	v615 = v578
	v619 = base.I64_extend_i32_u(v575)
	v620 = base.I64_extend_i32_u(v574)
	v623 = base.I64_extend_i32_u(v568)
	goto L26
L29:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v141 = F_read_stream_next_buffer(m, v102, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	F_LockBufferInternal(m, v141, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	if v141 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_UnlockReleaseBuffer(m, v141)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L6
	} else {
		goto L85
	}
L36:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
	if v164 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[0]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149+(v141^int32(-1))<<(uint(int32(2))%32))))
	v163 = v155
	goto L36
L38:
	;
	goto L39
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[1]))
	v163 = v157 + v141<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	v568 = v118 + int32(1)
	v571 = v121
	v574 = v124
	v575 = v125
	v576 = v126
	v577 = v127
	v578 = v128
	goto L35
L41:
	;
	goto L42
L42:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+19)))
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+16)))
	if (v169<<(uint(int32(8))%32)-v172)&int32(_a_F_pgstathashindex_0) != int32(16) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v178 = v172 + v163
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+12)))
	switch v179 & int32(15) {
	case 0:
		goto L46
	case 1:
		goto L47
	case 2:
		goto L48
	default:
		goto L45
	case 4:
		goto L44
	}
L44:
	;
	v568 = v118
	v571 = v121 + int32(1)
	v574 = v124
	v575 = v125
	v576 = v126
	v577 = v127
	v578 = v128
	goto L35
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L6
	} else {
		goto L77
	}
L46:
	;
	v568 = v118 + int32(1)
	v571 = v121
	v574 = v124
	v575 = v125
	v576 = v126
	v577 = v127
	v578 = v128
	goto L35
L47:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+12)))
	if base.Ui32(v347) < base.Ui32(int32(25)) {
		v491 = v126
		v492 = v127
		goto L63
	} else {
		goto L64
	}
L48:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+12)))
	if base.Ui32(v182) < base.Ui32(int32(25)) {
		v326 = v126
		v327 = v127
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+12)))
	v340 = v338 - v339
	v341 = int32(0)
	if v341 < v340 {
		goto L60
	} else {
		goto L61
	}
L50:
	;
	v188 = int32(base.Ui32(v182+int32(_a_F_pgstathashindex_1)) >> (uint(int32(2)) % 32))
	v190 = v188 & int32(_a_F_pgstathashindex_0)
	if v190 == int32(0) {
		v326 = v126
		v327 = v127
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v193 = int32(1)
	v195 = v163 + int32(20)
	if v190 != v193 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v205 = v193
	v208 = int32(0)
	v222 = v126
	v223 = v127
	goto L55
L53:
	;
	v266 = v193
	v283 = v126
	v284 = v127
	goto L54
L54:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v195+v266<<(uint(int32(2))%32))))
	v297 = int32(_a_F_pgstathashindex_2)
	v298 = v296 & v297
	v326 = v283 + base.I64_extend_i32_u(base.B2i32(v298 == v297))
	v327 = v284 + base.I64_extend_i32_u(base.B2i32(v298 != v297))
	goto L49
L55:
	;
	v232 = int32(2)
	v234 = v195 + v205<<(uint(v232)%32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = int32(_a_F_pgstathashindex_2)
	v237 = v235 & v236
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v244 = v242 & v236
	v248 = v222 + base.I64_extend_i32_u(base.B2i32(v237 == v236)) + base.I64_extend_i32_u(base.B2i32(v244 == v236))
	v256 = base.I64_extend_i32_u(base.B2i32(v244 != v236)) + (v223 + base.I64_extend_i32_u(base.B2i32(v237 != v236)))
	v258 = v205 + v232
	v260 = v208 + v232
	if v260 != v188&int32(_a_F_pgstathashindex_3) {
		v205 = v258
		v208 = v260
		v222 = v248
		v223 = v256
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v188&int32(1) == int32(0) {
		v326 = v248
		v327 = v256
		goto L49
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v266 = v258
	v283 = v248
	v284 = v256
	goto L54
L59:
	;
	v568 = v118
	v571 = v121
	v574 = v124 + int32(1)
	v575 = v125
	v576 = v326
	v577 = v327
	v578 = v128 + base.I64_extend_i32_u(v344)
	goto L35
L60:
	;
	v344 = v340
	goto L62
L61:
	;
	v344 = v341
	goto L62
L62:
	;
	goto L59
L63:
	;
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+14)))
	v504 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+12)))
	v505 = v503 - v504
	v506 = int32(0)
	if v506 < v505 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v353 = int32(base.Ui32(v347+int32(_a_F_pgstathashindex_1)) >> (uint(int32(2)) % 32))
	v355 = v353 & int32(_a_F_pgstathashindex_0)
	if v355 == int32(0) {
		v491 = v126
		v492 = v127
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v358 = int32(1)
	v360 = v163 + int32(20)
	if v355 != v358 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v370 = v358
	v373 = int32(0)
	v387 = v126
	v388 = v127
	goto L69
L67:
	;
	v431 = v358
	v448 = v126
	v449 = v127
	goto L68
L68:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v360+v431<<(uint(int32(2))%32))))
	v462 = int32(_a_F_pgstathashindex_2)
	v463 = v461 & v462
	v491 = v448 + base.I64_extend_i32_u(base.B2i32(v463 == v462))
	v492 = v449 + base.I64_extend_i32_u(base.B2i32(v463 != v462))
	goto L63
L69:
	;
	v397 = int32(2)
	v399 = v360 + v370<<(uint(v397)%32)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v401 = int32(_a_F_pgstathashindex_2)
	v402 = v400 & v401
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v409 = v407 & v401
	v413 = v387 + base.I64_extend_i32_u(base.B2i32(v402 == v401)) + base.I64_extend_i32_u(base.B2i32(v409 == v401))
	v421 = base.I64_extend_i32_u(base.B2i32(v409 != v401)) + (v388 + base.I64_extend_i32_u(base.B2i32(v402 != v401)))
	v423 = v370 + v397
	v425 = v373 + v397
	if v425 != v353&int32(_a_F_pgstathashindex_3) {
		v370 = v423
		v373 = v425
		v387 = v413
		v388 = v421
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v353&int32(1) == int32(0) {
		v491 = v413
		v492 = v421
		goto L63
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v431 = v423
	v448 = v413
	v449 = v421
	goto L68
L73:
	;
	v568 = v118
	v571 = v121
	v574 = v124
	v575 = v125 + int32(1)
	v576 = v491
	v577 = v492
	v578 = v128 + base.I64_extend_i32_u(v509)
	goto L35
L74:
	;
	v509 = v505
	goto L76
L75:
	;
	v509 = v506
	goto L76
L76:
	;
	goto L73
L77:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+12)))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	if v141 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v522 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v521
	F_errmsg(m, int32(_a_F_pgstathashindex_4), v32)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L6
	} else {
		goto L83
	}
L80:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[3]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526+(v141^int32(-1))*int32(56))+16))
	v541 = v532
	goto L79
L81:
	;
	goto L82
L82:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[4]))
	v535 = int32(56)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534+v141*v535-v535)+16))
	v541 = v540
	goto L79
L83:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(737), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v589 = v122 + int32(1)
	if v589 != v88 {
		v118 = v568
		v121 = v571
		v122 = v589
		v124 = v574
		v125 = v575
		v126 = v576
		v127 = v577
		v128 = v578
		goto L27
	} else {
		goto L86
	}
L86:
	;
	goto L28
L87:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v633 = base.I64_extend_i32_u(v88+(v608^int32(-1))) * v82
	if v633 == int64(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v644 = float64(0)
	goto L91
L90:
	;
	v644 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v615+v82*v623), float64(100)), base.F64_convert_i64_u(v633))
	goto L91
L91:
	;
	v648 = F_get_call_result_type(m, l0, int32(0), v32+int32(156))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	if v648 != int32(1) {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v32)+156))
	v653 = F_BlessTupleDesc(m, v652)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+136)) = v644
	*(*int64)(unsafe.Add(mBase, uint32(v32)+128)) = v613
	*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v614
	*(*int64)(unsafe.Add(mBase, uint32(v32)+112)) = v623
	*(*int64)(unsafe.Add(mBase, uint32(v32)+104)) = base.I64_extend_i32_u(v608)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+96)) = v619
	*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v620
	*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v32)+156)) = v653
	v669 = F_heap_form_tuple(m, v653, v32+int32(80), v32+int32(72))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v669)+16))
	v672 = F_HeapTupleHeaderGetDatum(m, v671)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	m.G0 = v32 + int32(160)
	return v672
L97:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v685 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_7), v32+int32(48))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(638), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_pgstathashindex_8), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(648), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v722 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_9), v32+int32(32))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(655), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	if v141 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v762
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v743 + int32(4)
	F_errmsg(m, int32(_a_F_pgstathashindex_10), v32+int32(16))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L6
	} else {
		goto L115
	}
L112:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[3]))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v747+(v141^int32(-1))*int32(56))+16))
	v762 = v753
	goto L111
L113:
	;
	goto L114
L114:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_pgstathashindex[4]))
	v756 = int32(56)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v755+v141*v756-v756)+16))
	v762 = v761
	goto L111
L115:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(709), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L6
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errmsg_internal(m, int32(_a_F_pgstathashindex_11), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_pgstathashindex_5), int32(767), int32(_a_F_pgstathashindex_6))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstatindex(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_superuser(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pgstatindex_0), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstatindex_1), int32(152), int32(_a_F_pgstatindex_2))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v28 = F_textToQualifiedNameList(m, v4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = F_makeRangeVarFromNameList(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v33 = F_relation_openrv(m, v30, int32(1))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = F_pgstatindex_impl(m, v33, l0)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_pgstattuple_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int64(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int64(0)
				} else {
					v14 = F_pgstat_relation(m, v12, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int64(0)
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_pkt_stream_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc(m, int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(70368744177664)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v5
		return int32(_a_F_pkt_stream_init_0)
	}
}
func F_pktreader_free(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	base.MemoryFill(m, l0, int32(0), int32(8))
	F_pfree(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_placeChar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	if l0 == int32(0) {
		v11 = F_palloc0_mul(m, int32(12), int32(256))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = v11
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v19 = v15 + v16*int32(12)
			if l2 <= int32(1) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				if v22 != 0 {
					v25 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							return v15
						} else {
							F_errcode(m, int32(22))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_placeChar_0), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_placeChar_1), int32(74), int32(_a_F_placeChar_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										return v15
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l4
					v43 = F_palloc(m, l4)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v43
						if l4 == int32(0) {
						} else {
							base.MemoryCopy(m, v43, l3, l4)
						}
						return v15
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v52 = int32(1)
				v56 = F_placeChar(m, v51, l1+v52, l2-v52, l3, l4)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19))) = v56
					return v15
				}
			}
		}
	} else {
		v15 = l0
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v19 = v15 + v16*int32(12)
		if l2 <= int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			if v22 != 0 {
				v25 = F_errstart(m, int32(19), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						return v15
					} else {
						F_errcode(m, int32(22))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_placeChar_0), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_placeChar_1), int32(74), int32(_a_F_placeChar_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									return v15
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l4
				v43 = F_palloc(m, l4)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v43
					if l4 == int32(0) {
					} else {
						base.MemoryCopy(m, v43, l3, l4)
					}
					return v15
				}
			}
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v52 = int32(1)
			v56 = F_placeChar(m, v51, l1+v52, l2-v52, l3, l4)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v56
				return v15
			}
		}
	}
}
func F_points_box(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v20 float64
	_ = v20
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v43 float64
	_ = v43
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = F_palloc(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		if base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			if base.F64_lt(v14, v20) != 0 {
				v28 = v20
				v29 = v8
			} else {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807)) {
					v28 = v20
					v29 = v8
				} else {
					v28 = v14
					v29 = v7
				}
			}
		} else {
			v28 = v14
			v29 = v7
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10))) = v28
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v31
		v34 = v7 + int32(8)
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		if base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v34)))
			if base.F64_gt(v43, v35)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v43)&int64(9223372036854775807))) != 0 {
				v52 = v43
				v53 = v8 + int32(8)
			} else {
				v52 = v35
				v53 = v34
			}
		} else {
			v52 = v35
			v53 = v34
		}
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v52
		v55 = *(*float64)(unsafe.Add(mBase, uint32(v53)))
		*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v55
		return base.I64_extend_i32_u(v10)
	}
}
func F_polish_ISO_8859_2_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 < v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v123 = v8 + int32(2)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v123 <= v124 {
		goto L40
	} else {
		goto L41
	}
L2:
	;
	if v57 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v19 = v8
	goto L5
L4:
	;
	v19 = v6
	goto L5
L5:
	;
	v26 = v8
	goto L7
L6:
	;
	v57 = v37
	goto L2
L7:
	;
	if v26 == v19 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v57 = int32(-1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v26))))
	if int32(243) < v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
	v26 = v49
	goto L7
L13:
	;
	v34 = v32 - int32(97)
	if v34 < int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v34)>>(uint(int32(3))%32)))+uint32(_c_F_polish_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v41)>>(uint(v34&int32(7))%32))&v37 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L12
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = v60 + v57
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v72 < v61 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v115 < int32(0) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v74 = v61
	goto L21
L20:
	;
	v74 = v72
	goto L21
L21:
	;
	v80 = v61
	goto L23
L22:
	;
	v115 = int32(1)
	goto L18
L23:
	;
	if v80 == v74 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v115 = int32(-1)
	goto L18
L26:
	;
	goto L27
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v80))))
	if int32(243) < v89 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v91 = v89 - int32(97)
	if v91 < int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v91)>>(uint(int32(3))%32)))+uint32(_c_F_polish_ISO_8859_2_stem[0]))))
	if int32(base.Ui32(v97)>>(uint(v91&int32(7))%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v106 = v80 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v80 = v106
	goto L23
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v118 + v115
	goto L1
L33:
	;
	return v275
L34:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v272
	v275 = int32(1)
	goto L33
L35:
	;
	v264 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L45
	} else {
		goto L88
	}
L36:
	;
	v258 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L45
	} else {
		goto L86
	}
L37:
	;
	v252 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_2))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L45
	} else {
		goto L84
	}
L38:
	;
	v246 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_3))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L45
	} else {
		goto L82
	}
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v177
	switch v155 - int32(1) {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	default:
		goto L34
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v123
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v124 < v128 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v158 = v124
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v163 = int32(0)
	v167 = F_find_among_b(m, l0, int32(_a_F_polish_ISO_8859_2_stem_4), int32(4), v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L45
	} else {
		goto L53
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v149
	v155 = F_find_among_b(m, l0, int32(_a_F_polish_ISO_8859_2_stem_5), int32(118), int32(_a_F_polish_ISO_8859_2_stem_6))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L45
	} else {
		goto L51
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v128
	v135 = F_find_among_b(m, l0, int32(_a_F_polish_ISO_8859_2_stem_7), int32(5), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	return int32(0)
L46:
	;
	if v135 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v123
	goto L43
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v123
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v143
	v145 = F_slice_del(m, l0)
	mBase = m.M
	if v145 < int32(0) {
		v275 = v145
		goto L33
	} else {
		goto L50
	}
L50:
	;
	goto L43
L51:
	;
	if v155 != 0 {
		goto L39
	} else {
		goto L52
	}
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v158 = v157
	goto L42
L53:
	;
	if v167 == int32(0) {
		v275 = v163
		goto L33
	} else {
		goto L54
	}
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v171 <= v173 {
		v275 = v163
		goto L33
	} else {
		goto L55
	}
L55:
	;
	switch v167 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L37
	case 2:
		goto L36
	case 3:
		goto L35
	default:
		goto L34
	}
L56:
	;
	v207 = F_slice_del(m, l0)
	mBase = m.M
	if v207 < int32(0) {
		v275 = v207
		goto L33
	} else {
		goto L72
	}
L57:
	;
	v203 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_8))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L45
	} else {
		goto L70
	}
L58:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v190 <= v177 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v186 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_9))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L45
	} else {
		goto L62
	}
L60:
	;
	v181 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v181 {
		goto L34
	} else {
		goto L61
	}
L61:
	;
	v275 = v181
	goto L33
L62:
	;
	if int32(0) <= v186 {
		goto L34
	} else {
		goto L63
	}
L63:
	;
	v275 = v186
	goto L33
L64:
	;
	v192 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v192 {
		goto L34
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v197 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_10))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L45
	} else {
		goto L68
	}
L67:
	;
	v275 = v192
	goto L33
L68:
	;
	if int32(0) <= v197 {
		goto L34
	} else {
		goto L69
	}
L69:
	;
	v275 = v197
	goto L33
L70:
	;
	if int32(0) <= v203 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	v275 = v203
	goto L33
L72:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v210
	v213 = v210 - int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v213 <= v214 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+v213))))
	if base.B2i32(v218 != int32(122))&base.B2i32(v218 != int32(99)) != 0 {
		goto L34
	} else {
		goto L74
	}
L74:
	;
	v227 = F_find_among_b(m, l0, int32(_a_F_polish_ISO_8859_2_stem_11), int32(5), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L45
	} else {
		goto L75
	}
L75:
	;
	if v227 == int32(0) {
		goto L34
	} else {
		goto L76
	}
L76:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v231
	switch v227 - int32(1) {
	case 0:
		goto L78
	case 1:
		goto L77
	default:
		goto L34
	}
L77:
	;
	v240 = F_slice_from_s(m, l0, int32(1), int32(_a_F_polish_ISO_8859_2_stem_12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L45
	} else {
		goto L80
	}
L78:
	;
	v235 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v235 {
		goto L34
	} else {
		goto L79
	}
L79:
	;
	v275 = v235
	goto L33
L80:
	;
	if int32(0) <= v240 {
		goto L34
	} else {
		goto L81
	}
L81:
	;
	v275 = v240
	goto L33
L82:
	;
	if int32(0) <= v246 {
		goto L34
	} else {
		goto L83
	}
L83:
	;
	v275 = v246
	goto L33
L84:
	;
	if int32(0) <= v252 {
		goto L34
	} else {
		goto L85
	}
L85:
	;
	v275 = v252
	goto L33
L86:
	;
	if int32(0) <= v258 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	v275 = v258
	goto L33
L88:
	;
	if v264 < int32(0) {
		v275 = v264
		goto L33
	} else {
		goto L89
	}
L89:
	;
	goto L34
}
func F_pqinitmask(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[0])) = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = int64(-15032385537)
	*(*int64)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = int64(-15032385537)
	v31 = int32(_a_F_pqinitmask_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v32 & base.I32_rotl(int32(-2), int32(4))
	v59 = int32(_a_F_pqinitmask_1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v60 & base.I32_rotl(int32(-2), int32(4))
	v87 = int32(_a_F_pqinitmask_0)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v88 & base.I32_rotl(int32(-2), int32(5))
	v115 = int32(_a_F_pqinitmask_1)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v116 & base.I32_rotl(int32(-2), int32(5))
	v143 = int32(_a_F_pqinitmask_0)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v144 & base.I32_rotl(int32(-2), int32(3))
	v171 = int32(_a_F_pqinitmask_1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v172 & base.I32_rotl(int32(-2), int32(3))
	v199 = int32(_a_F_pqinitmask_0)
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v200 & base.I32_rotl(int32(-2), int32(7))
	v227 = int32(_a_F_pqinitmask_1)
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v228 & base.I32_rotl(int32(-2), int32(7))
	v255 = int32(_a_F_pqinitmask_0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v256 & base.I32_rotl(int32(-2), int32(10))
	v283 = int32(_a_F_pqinitmask_1)
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v284 & base.I32_rotl(int32(-2), int32(10))
	v311 = int32(_a_F_pqinitmask_0)
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v312 & base.I32_rotl(int32(-2), int32(6))
	v339 = int32(_a_F_pqinitmask_1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v340 & base.I32_rotl(int32(-2), int32(6))
	v367 = int32(_a_F_pqinitmask_0)
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v368 & base.I32_rotl(int32(-2), int32(30))
	v395 = int32(_a_F_pqinitmask_1)
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v396 & base.I32_rotl(int32(-2), int32(30))
	v423 = int32(_a_F_pqinitmask_0)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[1])) = v424 & base.I32_rotl(int32(-2), int32(17))
	v451 = int32(_a_F_pqinitmask_1)
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v452 & base.I32_rotl(int32(-2), int32(17))
	v479 = int32(_a_F_pqinitmask_1)
	v480 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v480 & base.I32_rotl(int32(-2), int32(2))
	v507 = int32(_a_F_pqinitmask_1)
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v508 & base.I32_rotl(int32(-2), int32(14))
	v535 = int32(_a_F_pqinitmask_1)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_pqinitmask[2])) = v536 & base.I32_rotl(int32(-2), int32(13))
	return
}
func F_pqsignal_be(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = l1 + int32(2)
	switch v10 {
	case 0, 2:
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_pqsignal_be[0]))) = l1
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7+int32(16)))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(268435456)
	switch v10 {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(-2)
		v33 = int32(268435457)
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(268435460)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(_a_F_pqsignal_be_0)
		v33 = int32(268435461)
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v33 = int32(268435457)
	}
	if l0 == int32(17) {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v33
	} else {
	}
	v38 = v7 + int32(12)
	if base.Ui32(int32(65)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pqsignal_be[1])) = int32(28)
	} else {
		if v38 != 0 {
			v57 = l0 * int32(20)
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_pqsignal_be[2]))) = v58
			v60 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_pqsignal_be[3]))) = v60
			v62 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
			*(*int64)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_pqsignal_be[4]))) = v62
		} else {
		}
	}
	m.G0 = v7 + int32(32)
	return
}
func F_pre_format_elog_string(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pre_format_elog_string[0])) = l0
	return
}
func F_predicatelock_twophase_recover(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = base.I32_wrap_i64(l0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v13 {
	case 0:
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
		v19 = F_LWLockAcquire(m, v15+int32(3584), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			if base.B2i32(v23 == int32(0))|base.B2i32(v23 == v22) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v235 = m.ExcPending
				if v235 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_predicatelock_twophase_recover_0))
					mBase = m.M
					v238 = m.ExcPending
					if v238 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_predicatelock_twophase_recover_1), int32(0))
						mBase = m.M
						v242 = m.ExcPending
						if v242 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_predicatelock_twophase_recover_2), int32(_a_F_predicatelock_twophase_recover_3), int32(_a_F_predicatelock_twophase_recover_4))
							mBase = m.M
							v247 = m.ExcPending
							if v247 != 0 {
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
				v36 = v34 + int32(8)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				if v37 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v36
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v36
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v23
				v48 = v23 + int32(-64)
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v23-int32(60)))) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v23-int32(48)))) = int64(-1)
				v64 = int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(v23-int32(56)))) = v64
				v67 = v23 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v67
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v67
				*(*int64)(unsafe.Add(mBase, uint32(v23-int32(40)))) = v64
				*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v12
				*(*int64)(unsafe.Add(mBase, uint32(v23-int32(8)))) = int64(0)
				v82 = v23 - int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v23-int32(12)))) = v82
				*(*int32)(unsafe.Add(mBase, uint32(v82))) = v82
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v85
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v87
				if v87&int32(32) != 0 {
					v99 = v87
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = v93 + int32(1)
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
					v99 = v97
				}
				*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v99 | int32(1536)
				v106 = v23 - int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v23-int32(20)))) = v106
				v111 = v23 - int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v23-int32(28)))) = v111
				*(*int32)(unsafe.Add(mBase, uint32(v106))) = v106
				*(*int32)(unsafe.Add(mBase, uint32(v111))) = v111
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
				v117 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
				v123 = F_hash_search(m, v117, v10+int32(12), int32(1), v10+int32(11))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v48
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
					v128 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[1]))
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
					if v129 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v126
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
						v148 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
						v152 = F_LWLockAcquire(m, v148+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return
						} else {
							if v146 == int32(0) {
								v157 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = int64(0)
							} else {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
								if v162 == int32(1) {
									v167 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+308))
									v170 = base.B2i32(v168 != int32(2))
									*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v170)
									v172 = v170
								} else {
									v172 = int32(0)
								}
								v174 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
								if v172 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
									if v177 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
									} else {
										v180 = int32(3)
										if base.B2i32(base.Ui32(v146) < base.Ui32(v180))|base.B2i32(base.Ui32(v177) < base.Ui32(v180)) == int32(0) {
											if v146-v177 < int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
											} else {
											}
										} else {
											if base.Ui32(v177) <= base.Ui32(v146) {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
											}
										}
									}
								}
							}
							v196 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
							F_LWLockRelease(m, v196+int32(_a_F_predicatelock_twophase_recover_5))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								v257 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								F_LWLockRelease(m, v257+int32(3584))
								mBase = m.M
								v261 = m.ExcPending
								if v261 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					} else {
						v132 = int32(3)
						if base.B2i32(base.Ui32(v129) < base.Ui32(v132))|base.B2i32(base.Ui32(v126) < base.Ui32(v132)) == int32(0) {
							if int32(0) < v129-v126 {
								*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v126
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
								v148 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								v152 = F_LWLockAcquire(m, v148+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									if v146 == int32(0) {
										v157 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = int64(0)
									} else {
										v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
										if v162 == int32(1) {
											v167 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+308))
											v170 = base.B2i32(v168 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v170)
											v172 = v170
										} else {
											v172 = int32(0)
										}
										v174 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										if v172 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
										} else {
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
											if v177 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
											} else {
												v180 = int32(3)
												if base.B2i32(base.Ui32(v146) < base.Ui32(v180))|base.B2i32(base.Ui32(v177) < base.Ui32(v180)) == int32(0) {
													if v146-v177 < int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
													} else {
													}
												} else {
													if base.Ui32(v177) <= base.Ui32(v146) {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
													}
												}
											}
										}
									}
									v196 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									F_LWLockRelease(m, v196+int32(_a_F_predicatelock_twophase_recover_5))
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return
									} else {
										v257 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
										F_LWLockRelease(m, v257+int32(3584))
										mBase = m.M
										v261 = m.ExcPending
										if v261 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							} else {
								if v126 != v129 {
								} else {
									v249 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = v249 + int32(1)
								}
								v257 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								F_LWLockRelease(m, v257+int32(3584))
								mBase = m.M
								v261 = m.ExcPending
								if v261 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						} else {
							if base.Ui32(v129) <= base.Ui32(v126) {
								if v126 != v129 {
								} else {
									v249 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = v249 + int32(1)
								}
								v257 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								F_LWLockRelease(m, v257+int32(3584))
								mBase = m.M
								v261 = m.ExcPending
								if v261 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v128)+20)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v126
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
								v148 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
								v152 = F_LWLockAcquire(m, v148+int32(_a_F_predicatelock_twophase_recover_5), int32(0))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									if v146 == int32(0) {
										v157 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = int64(0)
									} else {
										v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])))
										if v162 == int32(1) {
											v167 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[5]))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+308))
											v170 = base.B2i32(v168 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[4])) = uint8(v170)
											v172 = v170
										} else {
											v172 = int32(0)
										}
										v174 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[3]))
										if v172 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
										} else {
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
											if v177 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
											} else {
												v180 = int32(3)
												if base.B2i32(base.Ui32(v146) < base.Ui32(v180))|base.B2i32(base.Ui32(v177) < base.Ui32(v180)) == int32(0) {
													if v146-v177 < int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
													} else {
													}
												} else {
													if base.Ui32(v177) <= base.Ui32(v146) {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v146
													}
												}
											}
										}
									}
									v196 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
									F_LWLockRelease(m, v196+int32(_a_F_predicatelock_twophase_recover_5))
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return
									} else {
										v257 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
										F_LWLockRelease(m, v257+int32(3584))
										mBase = m.M
										v261 = m.ExcPending
										if v261 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
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
	case 1:
		v202 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[6]))
		v204 = l2 + int32(4)
		v205 = F_get_hash_value(m, v202, v204)
		mBase = m.M
		v206 = m.ExcPending
		if v206 != 0 {
			return
		} else {
			v208 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
			v212 = F_LWLockAcquire(m, v208+int32(3584), int32(1))
			mBase = m.M
			v213 = m.ExcPending
			if v213 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v12
				v216 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[2]))
				v219 = int32(0)
				v221 = F_hash_search(m, v216, v10+int32(4), v219, v219)
				mBase = m.M
				v222 = m.ExcPending
				if v222 != 0 {
					return
				} else {
					v224 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_twophase_recover[0]))
					F_LWLockRelease(m, v224+int32(3584))
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
						return
					} else {
						v229 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
						F_CreatePredicateLock(m, v204, v205, v229)
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	default:
		m.G0 = v10 + int32(16)
		return
	}
}
func F_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v15
	v19 = int32(_a_F_prepare_cb_wrapper_1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v9 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v9 + int32(16)
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+147)) = uint8(v29)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+164)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+152)) = v33
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	if v37 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_prepare_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_prepare_cb_wrapper_3), v9)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_prepare_cb_wrapper_4), int32(1066), int32(_a_F_prepare_cb_wrapper_5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
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
		m.T0[v37].(func(*base.Module, int32, int32, int64))(m, v11, l1, l2)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_prepare_cb_wrapper[0])) = v60
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_printsimple(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v15 < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	m.T0[v18].(func(*base.Module, int32, int32))(m, l0, v14)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v24 = v9 + int32(-16)
	F_pq_beginmessage(m, v24, int32(68))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	F_enlargeStringInfo(m, v24, int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	v35 = int32(8)
	v39 = v28<<(uint(v35)%32) | int32(base.Ui32(v28)>>(uint(v35)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v32+v33))) = uint16(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v32 + int32(2)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v44 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = v44
	v51 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage(m, v9+int32(-16))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L70
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v51))))
	if v58 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v322 = v51 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v322 < v323 {
		v50 = v323
		v51 = v322
		goto L11
	} else {
		goto L69
	}
L14:
	;
	F_enlargeStringInfo(m, v9+int32(-16), int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = int32(3)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v74+v51<<(uint(v75)%32))))
	v84 = v13 + v50<<(uint(v75)%32) + v51*int32(100)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+96))
	switch v85 - int32(20) {
	case 0:
		goto L21
	default:
		goto L19
	case 3:
		goto L18
	case 5:
		goto L22
	case 6:
		goto L20
	}
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v66+v67))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v66 + int32(4)
	goto L13
L18:
	;
	v296 = v9 + int32(-48)
	v297 = base.I32_wrap_i64(v78)
	if int32(0) <= v297 {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L61
	}
L20:
	;
	v155 = v9 + int32(-48)
	v156 = base.I32_wrap_i64(v78)
	v157 = int32(0)
	if v156 == v157 {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v132 = v9 + int32(-48)
	if int64(0) <= v78 {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	v91 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v78))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v93 = int32(1)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	v97 = v95 & v93
	if v97 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = v93
	goto L26
L25:
	;
	v98 = int32(4)
	goto L26
L26:
	;
	if v95 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v91+v98, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L38
	}
L28:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v105 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v116 = int32(1)
	if v97 != 0 {
		v126 = int32(base.Ui32(v95)>>(uint(v116)%32)) - v116
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v108 = int32(16)
	goto L33
L32:
	;
	v108 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v105-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = int32(4)
	goto L36
L35:
	;
	v115 = v108
	goto L36
L36:
	;
	v126 = v115
	goto L27
L37:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v126 = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	goto L13
L39:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v132, v146)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L43
	}
L40:
	;
	v142 = v78
	v143 = int32(0)
	goto L42
L41:
	;
	v137 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v137)
	v142 = int64(0) - v78
	v143 = int32(1)
	goto L42
L42:
	;
	v145 = F_pg_ulltoa_n(m, v142, v132+v143)
	mBase = m.M
	v146 = v145 + v143
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132+v146))) = uint8(v148)
	goto L39
L43:
	;
	goto L13
L44:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v155, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L60
	}
L45:
	;
	v166 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v166)
	v276 = int32(1)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v172 = int32(1233)
	v177 = int32(base.Ui32((base.I32_clz(v156)^int32(31))*v172+v172) >> (uint(int32(12)) % 32))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177<<(uint(int32(2))%32))+uint32(_c_F_printsimple[0])))
	v182 = v177 + base.B2i32(base.Ui32(v180) <= base.Ui32(v156))
	if base.Ui32(int32(_a_F_printsimple_0)) <= base.Ui32(v156) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v186 = v156
	v188 = v157
	goto L51
L49:
	;
	v222 = v156
	v224 = v157
	goto L50
L50:
	;
	if base.Ui32(int32(100)) <= base.Ui32(v222) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v195 = v155 + v182 - v188
	v196 = int32(4)
	v199 = base.I32_div_u_s(v186, int32(_a_F_printsimple_0))
	v202 = v186 + v199*int32(-10000)
	v203 = int32(100)
	v204 = base.I32_div_u_s(v202, v203)
	v205 = int32(1)
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204<<(uint(v205)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v195-v196))) = uint16(v207)
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v202-v204*v203)<<(uint(v205)%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v195-int32(2)))) = uint16(v216)
	v219 = v188 + v196
	if base.Ui32(int32(99999999)) < base.Ui32(v186) {
		v186 = v199
		v188 = v219
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v222 = v199
	v224 = v219
	goto L50
L53:
	;
	goto L52
L54:
	;
	v235 = int32(2)
	v237 = int32(_a_F_printsimple_1)
	v239 = int32(100)
	v240 = base.I32_div_u_s(v222&v237, v239)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v222-v240*v239)&v237<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v155+v182-v224-v235))) = uint16(v248)
	v252 = v240
	v253 = v224 | v235
	goto L56
L55:
	;
	v252 = v222
	v253 = v224
	goto L56
L56:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v252) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252<<(uint(int32(1))%32))+uint32(_c_F_printsimple[1]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v155+v182-v253-int32(2)))) = uint16(v262)
	v276 = v182
	goto L44
L58:
	;
	goto L59
L59:
	;
	v265 = v252 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v265)
	v276 = v182
	goto L44
L60:
	;
	goto L13
L61:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v84)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v283
	F_errmsg_internal(m, int32(_a_F_printsimple_2), v11)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_printsimple_3), int32(137), int32(_a_F_printsimple_4))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_pq_sendcountedtext(m, v9+int32(-16), v296, v311)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L68
	}
L65:
	;
	v307 = v297
	v308 = int32(0)
	goto L67
L66:
	;
	v302 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v296))) = uint8(v302)
	v307 = int32(0) - v297
	v308 = int32(1)
	goto L67
L67:
	;
	v310 = F_pg_ultoa_n(m, v307, v296+v308)
	mBase = m.M
	v311 = v310 + v308
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v296+v311))) = uint8(v313)
	goto L64
L68:
	;
	goto L13
L69:
	;
	goto L12
L70:
	;
	m.G0 = v11 - int32(-64)
	return int32(1)
}
func F_privilege_to_string(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v1 = l0
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if v1&(v1-int64(1)) == int64(0) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(base.I64_ctz(v1))<<(uint(int32(2))%32))+uint32(_c_F_privilege_to_string[0])))
		m.G0 = v6 + int32(16)
		return v17
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(v6))) = uint32(v1)
			F_errmsg_internal(m, int32(_a_F_privilege_to_string_0), v6)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_privilege_to_string_1), int32(2631), int32(_a_F_privilege_to_string_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
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
func F_processPendingPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v5
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)) = uint16(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(-1)
	if base.Ui32(int32(25)) <= base.Ui32(v20) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(base.Ui32(v20+int32(_a_F_processPendingPage_0)) >> (uint(int32(2)) % 32))
	goto L3
L2:
	;
	v32 = v5
	goto L3
L3:
	;
	v34 = v32 & int32(_a_F_processPendingPage_1)
	if base.Ui32(l3) <= base.Ui32(v34) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = l3
	v46 = v5
	goto L7
L5:
	;
	v152 = v5
	v155 = v5
	goto L6
L6:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_ginInsertBAEntries(m, l0, v16+int32(8), v155&int32(_a_F_processPendingPage_1), v164, v165, v152)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L33
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2+int32(20)+v41&int32(_a_F_processPendingPage_1)<<(uint(int32(2))%32))))
	v60 = l2 + v57&int32(_a_F_processPendingPage_2)
	v61 = F_gintuple_get_attrnum(m, v51, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v140
	v155 = v102
	goto L6
L9:
	;
	return
L10:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)))
	if v63 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v107 = F_gintuple_get_key(m, v104, v60, v16+int32(7))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L26
	}
L12:
	;
	v65 = v16 + int32(8)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+2)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65))))
	v68 = int32(16)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	if v66|v67<<(uint(v68)%32) == v71|v72<<(uint(v68)%32) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+12)) = uint16(v97)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v99
	v102 = v61
	goto L11
L15:
	;
	v85 = v46 & int32(_a_F_processPendingPage_1)
	if v61 == v85 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+4)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)))
	if v78 == v79 {
		v82 = int32(1)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v82 = int32(0)
	goto L16
L20:
	;
	goto L19
L21:
	;
	v87 = v82
	goto L23
L22:
	;
	v87 = int32(0)
	goto L23
L23:
	;
	if v87 != 0 {
		v102 = v46
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_ginInsertBAEntries(m, l0, v65, v85, v88, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	goto L14
L26:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+7)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v111 <= v110 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v114 = v111 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v118 = F_repalloc_mul(m, v116, int32(8), v114)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	v128 = v110
	goto L29
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v129+v128<<(uint(int32(3))%32)))) = v107
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v135))) = uint8(v109)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v139 = int32(1)
	v140 = v138 + v139
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v140
	v143 = v41 + v139
	if base.Ui32(v143&int32(_a_F_processPendingPage_1)) <= base.Ui32(v34) {
		v41 = v143
		v46 = v102
		goto L7
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v124 = F_repalloc_mul(m, v121, int32(1), v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v128 = v127
	goto L29
L32:
	;
	goto L8
L33:
	;
	m.G0 = v16 + int32(16)
	return
}
func F_prsd_lextype(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v7 = F_palloc_mul(m, int32(12), int32(24))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = int32(1)
	goto L3
L3:
	;
	v16 = int32(12)
	v18 = v7 + v12*v16
	*(*int32)(unsafe.Add(mBase, uint32(v18-v16))) = v12
	v25 = v12 << (uint(int32(2)) % 32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_prsd_lextype[0])))
	v27 = F_pstrdup(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+276)) = int32(0)
	return base.I64_extend_i32_u(v7)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18-int32(8)))) = v27
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_prsd_lextype[1])))
	v33 = F_pstrdup(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18-int32(4)))) = v33
	v37 = v12 + int32(1)
	if v37 != int32(24) {
		v12 = v37
		goto L3
	} else {
		goto L7
	}
L7:
	;
	goto L4
}
func F_pt_contained_poly(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v11 = F_point_inside(m, v2, v8, v4+int32(40))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v11 != int32(0)))
		}
	}
}
func F_pull_up_union_leaf_queries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = l0
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v25 != int32(142) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L29
	}
L3:
	;
	goto L2
L4:
	;
	if v25 != int32(63) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_pull_up_union_leaf_queries(m, v123, l1, l2, l3, l4)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L28
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v32 = F_palloc0(m, int32(36))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+12)) = int64(0)
	v36 = v30 + l4
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(324)
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = v44
	goto L12
L11:
	;
	v45 = v41
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v45
	v49 = F_palloc0(m, v45<<(uint(int32(1))%32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v52 == int32(0) {
		v98 = v41
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v98
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+136))
	v108 = F_lappend(m, v107, v32)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L25
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v55 <= int32(0) {
		v98 = v41
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v59 = int32(0)
	v61 = v55
	v63 = v41
	goto L17
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v59<<(uint(int32(2))%32))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+26)))
	if v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v98 = v90
	goto L14
L19:
	;
	v77 = F_makeVarFromTargetEntry(m, v36, v73)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v89 = v61
	v90 = v63
	goto L21
L21:
	;
	v92 = v59 + int32(1)
	if v92 < v89 {
		v59 = v92
		v61 = v89
		v63 = v90
		goto L17
	} else {
		goto L24
	}
L22:
	;
	v79 = F_lappend(m, v63, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v81<<(uint(int32(1))%32)-int32(2)))) = uint16(v81)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v89 = v88
	v90 = v79
	goto L21
L24:
	;
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+136)) = v108
	v112 = F_palloc0(m, int32(8))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = int32(63)
	v118 = F_pull_up_subqueries_recurse(m, l1, v112, int32(0), v32)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v13 + int32(16)
	return
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v15 = v126
	goto L1
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v131
	F_errmsg_internal(m, int32(_a_F_pull_up_union_leaf_queries_0), v13)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pull_up_union_leaf_queries_1), int32(1905), int32(_a_F_pull_up_union_leaf_queries_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pull_varattnos_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = int32(0)
	if l0 == v3 {
		v27 = v3
		return v27
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(6) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v9 != v10 {
				v27 = v3
				return v27
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v12 != 0 {
					v27 = v3
					return v27
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
					v17 = F_bms_add_member(m, v13, v14+int32(7))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
						return int32(0)
					}
				}
			}
		} else {
			v25 = F_expression_tree_walker_impl(m, l0, int32(948), l1)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v25
				return v27
			}
		}
	}
}
func F_pullf_create_mbuf_reader(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14361(m, l0, l1, int32(_a_F_pullf_create_mbuf_reader_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pullf_read_fixed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pullf_read_max(m, l0, l1, v8+int32(12), l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 < int32(0) {
			v32 = v12
			m.G0 = v8 + int32(16)
			return v32
		} else {
			if v12 != l1 {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_px_debug(m, int32(_a_F_pullf_read_fixed_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v32 = int32(-100)
					m.G0 = v8 + int32(16)
					return v32
				}
			} else {
				v25 = int32(0)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				if base.B2i32(l1 == v25)|base.B2i32(v28 == l2) != 0 {
					v32 = v25
				} else {
					base.MemoryCopy(m, l2, v28, l1)
					v32 = v25
				}
				m.G0 = v8 + int32(16)
				return v32
			}
		}
	}
}
func F_pushValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	v4 = l3
	v5 = l4
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l2 <= int32(2047) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return
L2:
	;
	v174 = l2 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v177 = v175 - v176
	v178 = v174 + v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v179 <= v178 {
		goto L32
	} else {
		goto L33
	}
L3:
	;
	v154 = F_palloc0(m, int32(12))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L19
	} else {
		goto L30
	}
L4:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v136 = F_errsave_start(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L19
	} else {
		goto L25
	}
L7:
	;
	v18 = int32(-1)
	if l2 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v109 = int32(0)
	goto L9
L9:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v112 = v110 - v111
	if v112 < int32(_a_F_pushValue_0) {
		goto L3
	} else {
		goto L18
	}
L10:
	;
	v109 = v89 ^ int32(-1)
	goto L9
L11:
	;
	v30 = v18
	v31 = l1
	v32 = int32(0)
	goto L14
L12:
	;
	v68 = v18
	v69 = l1
	goto L13
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32((v74^int32(base.Ui32(v68)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_c_F_pushValue[0])))
	v89 = v80 ^ v68<<(uint(int32(8))%32)
	goto L10
L14:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v38 = int32(24)
	v41 = int32(2)
	v43 = *(*int32)(unsafe.Add(mBase, uint32((v37^int32(base.Ui32(v30)>>(uint(v38)%32)))<<(uint(v41)%32))+uint32(_c_F_pushValue[0])))
	v44 = int32(8)
	v46 = v43 ^ v30<<(uint(v44)%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32((v36^int32(base.Ui32(v46)>>(uint(v38)%32)))<<(uint(v41)%32))+uint32(_c_F_pushValue[0])))
	v55 = v52 ^ v46<<(uint(v44)%32)
	v57 = v31 + v41
	v59 = v32 + v41
	if v59 != l2&int32(-2) {
		v30 = v55
		v31 = v57
		v32 = v59
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if l2&int32(1) == int32(0) {
		v89 = v55
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v68 = v55
	v69 = v57
	goto L13
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v116 = F_errsave_start(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	if v116 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v123
	F_errmsg(m, int32(_a_F_pushValue_1), v14+int32(16))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errsave_finish(m, v115, int32(_a_F_pushValue_2), int32(555), int32(_a_F_pushValue_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L2
L25:
	;
	if v136 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v143
	F_errmsg(m, int32(_a_F_pushValue_4), v14)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v135, int32(_a_F_pushValue_2), int32(588), int32(_a_F_pushValue_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v109
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+2)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)) = uint8(v4)
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = l2&int32(4095) | v112<<(uint(int32(12))%32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v168 = F_lcons(m, v154, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v168
	goto L2
L32:
	;
	v186 = v179
	v187 = v176
	goto L35
L33:
	;
	v209 = v175
	goto L34
L34:
	;
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v193 = v186 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v193
	v195 = F_repalloc(m, v187, v193)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L37
	}
L36:
	;
	v209 = v198
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v195
	v198 = v177 + v195
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v200 <= v178 {
		v186 = v200
		v187 = v195
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	base.MemoryCopy(m, v209, l1, l2)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v215 = v214 + l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v215
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v219 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v174 + v223
	goto L1
}
func F_pushf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		v14 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l3, l2, v9+int32(12))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v14 {
				v22 = v14
				v24 = F_palloc0(m, int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
					if v22 != 0 {
						v32 = F_palloc(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
							v41 = v35
							m.G0 = v9 + int32(16)
							return v41
						}
					} else {
						v34 = int32(0)
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
						v41 = v35
						m.G0 = v9 + int32(16)
						return v41
					}
				}
			} else {
				v41 = v14
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		v22 = int32(0)
		v24 = F_palloc0(m, int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
			if v22 != 0 {
				v32 = F_palloc(m, v22)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v41 = v35
					m.G0 = v9 + int32(16)
					return v41
				}
			} else {
				v34 = int32(0)
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v41 = v35
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	}
}
func F_pushf_flush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v54
L2:
	;
	v6 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	v54 = int32(0)
	goto L1
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	if int32(0) < v11 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v38 != 0 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	if int32(0) < v27 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v21 = m.T0[v19].(func(*base.Module, int32, int32, int32, int32) int32)(m, v14, v20, v16, v15)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v25 = F_pushf_write(m, v14, v16, v15)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v27 = v21
	goto L10
L16:
	;
	v27 = v25
	goto L10
L17:
	;
	v30 = int32(-12)
	goto L19
L18:
	;
	v30 = v27
	goto L19
L19:
	;
	if v30 < int32(0) {
		v54 = v30
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L9
L21:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
	v41 = m.T0[v38].(func(*base.Module, int32, int32) int32)(m, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v46 != 0 {
		v6 = v46
		goto L5
	} else {
		goto L26
	}
L24:
	;
	if v41 < int32(0) {
		v54 = v41
		goto L1
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L6
}
func F_put_notnull_info(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
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
	var v21 int64
	_ = v21
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	if v12 <= l1 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v14 = int32(_a_F_put_notnull_info_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_put_notnull_info[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_put_notnull_info[0])) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v27 = v21
	goto L7
L5:
	;
	goto L6
L6:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v71 = v67 + base.I32_wrap_i64(int64(base.Ui64(l1)>>(uint(int64(2))%64)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v75 = int32(1)
	v78 = base.I32_wrap_i64(l1) << (uint(v75) % 32) & int32(6)
	if l2 != 0 {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v29 = base.I64_div_s(v27, int64(4))
	v30 = base.I32_wrap_i64(v29)
	if v30 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_put_notnull_info[0])) = v15
	goto L6
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	if v56 <= l1 {
		v27 = v56
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v34 = F_palloc0(m, int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v43 = F_repalloc0(m, v40, v30, v30<<(uint(int32(1))%32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L15
	}
L13:
	;
	return
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v34
	v52 = int64(128)
	goto L9
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v43
	v52 = v29 << (uint(int64(3)) % 64) & int64(4294967288)
	goto L9
L16:
	;
	goto L8
L17:
	;
	v85 = v75
	goto L19
L18:
	;
	v85 = int32(2)
	goto L19
L19:
	;
	v87 = v72&(int32(3)<<(uint(v78)%32)^int32(-1)) | v85<<(uint(v78)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v87)
	goto L3
}
