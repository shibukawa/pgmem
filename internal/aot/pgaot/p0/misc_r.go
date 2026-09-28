package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RegisterBuiltinShmemCallbacks(m *base.Module) {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_0))
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_1))
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_2))
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_3))
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_4))
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_5))
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_6))
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_7))
								v24 = m.ExcPending
								if v24 != 0 {
									return
								} else {
									F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_8))
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_9))
										v30 = m.ExcPending
										if v30 != 0 {
											return
										} else {
											F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_10))
											v33 = m.ExcPending
											if v33 != 0 {
												return
											} else {
												F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_11))
												v36 = m.ExcPending
												if v36 != 0 {
													return
												} else {
													F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_12))
													v39 = m.ExcPending
													if v39 != 0 {
														return
													} else {
														F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_13))
														v42 = m.ExcPending
														if v42 != 0 {
															return
														} else {
															F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_14))
															v45 = m.ExcPending
															if v45 != 0 {
																return
															} else {
																F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_15))
																v48 = m.ExcPending
																if v48 != 0 {
																	return
																} else {
																	F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_16))
																	v51 = m.ExcPending
																	if v51 != 0 {
																		return
																	} else {
																		F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_17))
																		v54 = m.ExcPending
																		if v54 != 0 {
																			return
																		} else {
																			F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_18))
																			v57 = m.ExcPending
																			if v57 != 0 {
																				return
																			} else {
																				F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_19))
																				v60 = m.ExcPending
																				if v60 != 0 {
																					return
																				} else {
																					F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_20))
																					v63 = m.ExcPending
																					if v63 != 0 {
																						return
																					} else {
																						F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_21))
																						v66 = m.ExcPending
																						if v66 != 0 {
																							return
																						} else {
																							F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_22))
																							v69 = m.ExcPending
																							if v69 != 0 {
																								return
																							} else {
																								F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_23))
																								v72 = m.ExcPending
																								if v72 != 0 {
																									return
																								} else {
																									F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_24))
																									v75 = m.ExcPending
																									if v75 != 0 {
																										return
																									} else {
																										F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_25))
																										v78 = m.ExcPending
																										if v78 != 0 {
																											return
																										} else {
																											F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_26))
																											v81 = m.ExcPending
																											if v81 != 0 {
																												return
																											} else {
																												F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_27))
																												v84 = m.ExcPending
																												if v84 != 0 {
																													return
																												} else {
																													F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_28))
																													v87 = m.ExcPending
																													if v87 != 0 {
																														return
																													} else {
																														F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_29))
																														v90 = m.ExcPending
																														if v90 != 0 {
																															return
																														} else {
																															F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_30))
																															v93 = m.ExcPending
																															if v93 != 0 {
																																return
																															} else {
																																F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_31))
																																v96 = m.ExcPending
																																if v96 != 0 {
																																	return
																																} else {
																																	F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_32))
																																	v99 = m.ExcPending
																																	if v99 != 0 {
																																		return
																																	} else {
																																		F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_33))
																																		v102 = m.ExcPending
																																		if v102 != 0 {
																																			return
																																		} else {
																																			F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_34))
																																			v105 = m.ExcPending
																																			if v105 != 0 {
																																				return
																																			} else {
																																				F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_35))
																																				v108 = m.ExcPending
																																				if v108 != 0 {
																																					return
																																				} else {
																																					F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_36))
																																					v111 = m.ExcPending
																																					if v111 != 0 {
																																						return
																																					} else {
																																						F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_37))
																																						v114 = m.ExcPending
																																						if v114 != 0 {
																																							return
																																						} else {
																																							F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_38))
																																							v117 = m.ExcPending
																																							if v117 != 0 {
																																								return
																																							} else {
																																								F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_39))
																																								v120 = m.ExcPending
																																								if v120 != 0 {
																																									return
																																								} else {
																																									F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_40))
																																									v123 = m.ExcPending
																																									if v123 != 0 {
																																										return
																																									} else {
																																										F_RegisterShmemCallbacks(m, int32(_a_F_RegisterBuiltinShmemCallbacks_41))
																																										v126 = m.ExcPending
																																										if v126 != 0 {
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
					}
				}
			}
		}
	}
}
func F_ReleaseAuxProcessResourcesCallback(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
	v5 = int32(1)
	v7 = base.B2i32(l0 == int32(0))
	F_ResourceOwnerReleaseInternal(m, v4, v5, v7, v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
		F_ResourceOwnerReleaseInternal(m, v12, int32(2), v7, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
			F_ResourceOwnerReleaseInternal(m, v18, int32(3), v7, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseAuxProcessResourcesCallback[0]))
				v25 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v25)
				return
			}
		}
	}
}
func F_ReleaseDeletionLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v38 int32
	_ = v38
	var v55 int32
	_ = v55
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(1259) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_UnlockRelationOid(m, v7, int32(8))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		v13 = int32(1)
		if v4 <= int32(3591) {
			if v4 <= int32(2670) {
				switch v4 - int32(1213) {
				case 0, 1, 19, 20, 47, 48, 49:
					v84 = v13
				case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
					v84 = int32(0)
				default:
					if base.Ui32(int32(2)) <= base.Ui32(v4-int32(2396)) {
						v84 = int32(0)
					} else {
						v84 = v13
					}
				}
			} else {
				v25 = v4 - int32(2671)
				if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v25))|base.B2i32(int32(1)<<(uint(v25)%32)&int32(226492515) == int32(0)) != 0 {
					if base.B2i32(base.Ui32(v4-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v4-int32(2846)) < base.Ui32(int32(2))) != 0 {
						v84 = v13
					} else {
						v84 = int32(0)
					}
				} else {
					v84 = v13
				}
			}
		} else {
			if v4 <= int32(_a_F_ReleaseDeletionLock_0) {
				v38 = v4 - int32(_a_F_ReleaseDeletionLock_1)
				if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v38))|base.B2i32(int32(1)<<(uint(v38)%32)&int32(963) == int32(0)) != 0 {
					if base.Ui32(v4-int32(3592)) < base.Ui32(int32(2)) {
						v84 = v13
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(v4-int32(4060)) {
							v84 = int32(0)
						} else {
							v84 = v13
						}
					}
				} else {
					v84 = v13
				}
			} else {
				switch v4 - int32(_a_F_ReleaseDeletionLock_2) {
				case 0, 1, 2, 3, 4, 59, 60:
					v84 = v13
				case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
					v84 = int32(0)
				default:
					if base.Ui32(v4-int32(_a_F_ReleaseDeletionLock_3)) < base.Ui32(int32(3)) {
						v84 = v13
					} else {
						v55 = v4 - int32(_a_F_ReleaseDeletionLock_4)
						if base.Ui32(int32(15)) < base.Ui32(v55) {
							v84 = int32(0)
						} else {
							if int32(1)<<(uint(v55)%32)&int32(_a_F_ReleaseDeletionLock_5) != 0 {
								v84 = v13
							} else {
								v84 = int32(0)
							}
						}
					}
				}
			}
		}
		v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v84 != 0 {
			F_UnlockSharedObject(m, v86, v85, int32(8))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return
			} else {
				return
			}
		} else {
			F_UnlockDatabaseObject(m, v86, v85, int32(8))
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ReleaseLockIfHeld(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseLockIfHeld[0]))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(0)
	goto L3
L2:
	;
	v14 = v13
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = v16
	goto L5
L4:
	;
	return
L5:
	;
	v28 = v19 - int32(1)
	if v28 < int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v33)+8))
	if v37 < v36 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v33 = v15 + v28<<(uint(int32(4))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v34 != v14 {
		v19 = v28
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v40 = v16 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v36 - v37
	if v14 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v56 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_LockRelease(m, l0, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L18
	}
L12:
	;
	F_ResourceOwnerForgetLock(m, v13, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v47 = v40
	goto L14
L14:
	;
	if v47 <= v28 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v47 = v46
	goto L14
L17:
	;
	v51 = v15 + v47<<(uint(int32(4))%32)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v54
	return
L18:
	;
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v65 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	if v65 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_errmsg_internal(m, int32(_a_F_ReleaseLockIfHeld_0), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ReleaseLockIfHeld_1), int32(2728), int32(_a_F_ReleaseLockIfHeld_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L4
}
func F_RemoveLogrotateSignalFiles(m *base.Module) {
	var v2 int32
	_ = v2
	v2 = F_unlink(m, int32(_a_F_RemoveLogrotateSignalFiles_0))
	return
}
func F_RemoveNonParentXlogFiles(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v293 int32
	_ = v293
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	v21 = int64(*(*int32)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[0])))
	v22 = base.I64_div_u_s(l0-int64(1), v21)
	v24 = base.I64_div_u_s(int64(4294967296), v21)
	v25 = base.I64_div_u_s(v22, v24)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+20)) = uint32(v25)
	v28 = v22 - v24*v25
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+24)) = uint32(v28)
	v30 = base.I64_div_u_s(l0, v21)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v30
	v33 = v15 + int32(48)
	v38 = F_pg_snprintf(m, v33, int32(64), int32(_a_F_RemoveNonParentXlogFiles_0), v15+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v42 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v42 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
	F_errmsg_internal(m, int32(_a_F_RemoveNonParentXlogFiles_1), v15)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v54 = F_AllocateDir(m, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_errfinish(m, int32(_a_F_RemoveNonParentXlogFiles_3), int32(4024), int32(_a_F_RemoveNonParentXlogFiles_4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v57 = F_ReadDir(m, v54, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v57 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = v15 + int32(48) | int32(8)
	v70 = v57
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_FreeDir(m, v54)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L66
	}
L14:
	;
	v78 = v70 + int32(19)
	v79 = F_strlen(m, v78)
	mBase = m.M
	if v79 != int32(24) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v278 = F_ReadDir(m, v54, int32(_a_F_RemoveNonParentXlogFiles_2))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L17:
	;
	v82 = int32(_a_F_RemoveNonParentXlogFiles_5)
	v86 = m.G0
	v88 = v86 - int32(32)
	v89 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+24)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = v89
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[1])))
	if v97 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v165 != int32(24) {
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v165 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RemoveNonParentXlogFiles[2])))
	if v101 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v105 = v78
	goto L25
L23:
	;
	goto L24
L24:
	;
	v115 = v82
	v116 = v97
	goto L28
L25:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	if v111 == v97 {
		v105 = v105 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v165 = v105 - v78
	goto L18
L27:
	;
	goto L26
L28:
	;
	v123 = v88 + int32(base.Ui32(v116)>>(uint(int32(3))%32))&int32(28)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v125 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v124 | v125<<(uint(v116)%32)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v129 != 0 {
		v115 = v115 + v125
		v116 = v129
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v132 == int32(0) {
		v155 = v78
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v165 = v155 - v78
	goto L18
L32:
	;
	v136 = v78
	v137 = v132
	goto L33
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(base.Ui32(v137)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v145)>>(uint(v137)%32))&int32(1) == int32(0) {
		v155 = v136
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v155 = v153
	goto L31
L35:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v153 = v136 + int32(1)
	if v151 != 0 {
		v136 = v153
		v137 = v151
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v169 = v15 + int32(48)
	goto L40
L38:
	;
	if int32(0) <= v207-v208 {
		goto L16
	} else {
		goto L51
	}
L40:
	;
	goto L41
L41:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v176 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v177 = v78
	v178 = v169
	v179 = int32(8)
	v180 = v176
	goto L46
L43:
	;
	v203 = v169
	v207 = int32(0)
	goto L44
L44:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	goto L38
L45:
	;
	v203 = v198
	v207 = v200
	goto L44
L46:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if base.B2i32(v180 != v182)|base.B2i32(v182 == int32(0)) != 0 {
		v198 = v178
		v200 = v180
		goto L45
	} else {
		goto L48
	}
L47:
	;
	v198 = v192
	v200 = int32(0)
	goto L45
L48:
	;
	v188 = v179 - int32(1)
	if v188 == int32(0) {
		v198 = v178
		v200 = v180
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v191 = int32(1)
	v192 = v178 + v191
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)))
	if v193 != 0 {
		v177 = v177 + v191
		v178 = v192
		v179 = v188
		v180 = v193
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v219 = v70 + int32(27)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if base.B2i32(v222 == int32(0))|base.B2i32(v222 != v225) != 0 {
		v243 = v222
		v244 = v225
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v243-v244 <= int32(0) {
		goto L16
	} else {
		goto L59
	}
L53:
	;
	goto L52
L54:
	;
	v228 = v219
	v229 = v64
	goto L55
L55:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+1)))
	if v233 == int32(0) {
		v243 = v233
		v244 = v232
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v243 = v233
	v244 = v232
	goto L53
L57:
	;
	v236 = int32(1)
	if v233 == v232 {
		v228 = v228 + v236
		v229 = v229 + v236
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v248 = m.G0
	v250 = v248 - int32(1136)
	m.G0 = v250
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v250)+4)) = int32(_a_F_RemoveNonParentXlogFiles_6)
	v256 = v250 + int32(112)
	v259 = F_pg_snprintf(m, v256, int32(1024), int32(_a_F_RemoveNonParentXlogFiles_7), v250)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v265 = F___fstatat(m, int32(-100), v256, v250+int32(16), int32(0))
	mBase = m.M
	goto L61
L61:
	;
	m.G0 = v250 + int32(1136)
	if v265 == int32(0) {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	F_RemoveXlogFile(m, v70, v30+int64(10), v15+int32(40), l1)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L16
L64:
	;
	if v278 != 0 {
		v70 = v278
		goto L14
	} else {
		goto L65
	}
L65:
	;
	goto L15
L66:
	;
	m.G0 = v15 + int32(112)
	return
}
func F_ReqShutdownXLOG(m *base.Module, l0 int32, l1 int32) {
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
	*(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_ReqShutdownXLOG_0), v8)
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
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_ReqShutdownXLOG_0), v15)
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
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[2]))
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[3]))
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[4]))
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
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ReqShutdownXLOG[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_ResolveCminCmaxDuringDecoding(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int64
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int64
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int64
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int64
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1264)
	m.G0 = v20
	if l0 == v7 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L13
	} else {
		goto L163
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L13
	} else {
		goto L159
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L13
	} else {
		goto L156
	}
L4:
	;
	m.G0 = v20 + int32(1264)
	return v694
L5:
	;
	v694 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v25
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+152)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v20)+144)) = v27
	v32 = v20 + int32(144)
	if l3 < v25 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+160)) = uint16(v63)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v65
	v67 = int32(0)
	v69 = F_hash_search(m, l0, v32, v67, v67)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(140)))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(136)))) = v61
	goto L8
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[0]))
	v54 = v41 + (l3^int32(-1))*int32(56)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[1]))
	v49 = int32(56)
	v54 = v48 + l3*v49 - v49
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	if v69 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v78 = int32(1)
	if v75 <= int32(3591) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v671 = v69
	goto L17
L17:
	;
	if l4 != 0 {
		goto L152
	} else {
		goto L153
	}
L18:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[2]))
	v153 = F_AllocateDir(m, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L43
	}
L19:
	;
	goto L18
L20:
	;
	v149 = int32(0)
	goto L19
L21:
	;
	if base.B2i32(base.Ui32(v75-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v75-int32(2846)) < base.Ui32(int32(2))) != 0 {
		v149 = v78
		goto L19
	} else {
		goto L42
	}
L22:
	;
	if v75 <= int32(2670) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v75 <= int32(_a_F_ResolveCminCmaxDuringDecoding_1) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	switch v75 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v149 = v78
		goto L19
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L20
	default:
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v90 = v75 - int32(2671)
	if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v90))|base.B2i32(int32(1)<<(uint(v90)%32)&int32(226492515) == int32(0)) != 0 {
		goto L21
	} else {
		goto L30
	}
L28:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v75-int32(2396)) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v149 = v78
	goto L19
L30:
	;
	v149 = v78
	goto L19
L31:
	;
	if base.Ui32(v75-int32(3592)) < base.Ui32(int32(2)) {
		v149 = v78
		goto L19
	} else {
		goto L40
	}
L32:
	;
	v103 = v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_2)
	if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v103))|base.B2i32(int32(1)<<(uint(v103)%32)&int32(963) == int32(0)) != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	switch v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_3) {
	case 0, 1, 2, 3, 4, 59, 60:
		v149 = v78
		goto L19
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L20
	default:
		goto L36
	}
L35:
	;
	v149 = v78
	goto L19
L36:
	;
	if base.Ui32(v75-int32(_a_F_ResolveCminCmaxDuringDecoding_4)) < base.Ui32(int32(3)) {
		v149 = v78
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v120 = v75 - int32(_a_F_ResolveCminCmaxDuringDecoding_5)
	if base.Ui32(int32(15)) < base.Ui32(v120) {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	if int32(1)<<(uint(v120)%32)&int32(_a_F_ResolveCminCmaxDuringDecoding_6) != 0 {
		v149 = v78
		goto L19
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v75-int32(4060)) {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	v149 = v78
	goto L19
L42:
	;
	goto L20
L43:
	;
	v156 = F_ReadDir(m, v153, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	if v156 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v149 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v393 = v7
	goto L47
L47:
	;
	F_FreeDir(m, v153)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L106
	}
L48:
	;
	v159 = int32(0)
	goto L50
L49:
	;
	v159 = v151
	goto L50
L50:
	;
	v166 = v156
	v173 = v7
	goto L51
L51:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+19)))
	if v181 != int32(46) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v393 = v377
	goto L47
L53:
	;
	v382 = F_ReadDir(m, v153, int32(_a_F_ResolveCminCmaxDuringDecoding_0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L104
	}
L54:
	;
	v194 = v166 + int32(19)
	v195 = int32(_a_F_ResolveCminCmaxDuringDecoding_7)
	goto L61
L55:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+20)))
	if v184 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L56
	}
L56:
	;
	if v184 != int32(46) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+21)))
	if v189 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	if v233-v234 != 0 {
		v377 = v173
		goto L53
	} else {
		goto L72
	}
L61:
	;
	goto L62
L62:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v202 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v203 = v194
	v204 = v195
	v205 = int32(4)
	v206 = v202
	goto L67
L64:
	;
	v229 = v195
	v233 = int32(0)
	goto L65
L65:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	goto L59
L66:
	;
	v229 = v224
	v233 = v226
	goto L65
L67:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if base.B2i32(v206 != v208)|base.B2i32(v208 == int32(0)) != 0 {
		v224 = v204
		v226 = v206
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v224 = v218
	v226 = int32(0)
	goto L66
L69:
	;
	v214 = v205 - int32(1)
	if v214 == int32(0) {
		v224 = v204
		v226 = v206
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v217 = int32(1)
	v218 = v204 + v217
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v219 != 0 {
		v203 = v203 + v217
		v204 = v218
		v205 = v214
		v206 = v219
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(132)))) = v20 + int32(176)
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(128)))) = v20 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+124)) = v20 + int32(168)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+120)) = v20 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v20 + int32(184)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v20 + int32(204)
	v263 = F_sscanf(m, v194, int32(_a_F_ResolveCminCmaxDuringDecoding_8), v20+int32(112))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	if v263 != int32(6) {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
	if v267 != v159 {
		v377 = v173
		goto L53
	} else {
		goto L75
	}
L75:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v20)+184))
	if v269 != v75 {
		v377 = v173
		goto L53
	} else {
		goto L76
	}
L76:
	;
	v271 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+172)))
	v272 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+168)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	v274 = F_TransactionIdDidCommit(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	if v274 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L78
	}
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v20)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v280
	v286 = F_bsearch(m, v20+int32(240), v279, v278, int32(4), int32(187))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	if v286 == int32(0) {
		v377 = v173
		goto L53
	} else {
		goto L80
	}
L80:
	;
	v291 = F_palloc(m, int32(1032))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v291))) = v271<<(uint(int64(32))%64) | v272
	v298 = v291 + int32(8)
	if (v194^v298)&int32(3) != 0 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v373 = F_lappend(m, v173, v291)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L13
	} else {
		goto L103
	}
L83:
	;
	goto L82
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v353))) = uint8(v352)
	if v352&int32(255) == int32(0) {
		goto L83
	} else {
		goto L99
	}
L85:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v351 = v194
	v352 = v304
	v353 = v298
	goto L84
L86:
	;
	goto L87
L87:
	;
	if v194&int32(3) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v308 = v194
	v310 = v298
	goto L91
L89:
	;
	v322 = v194
	v324 = v298
	goto L90
L90:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v329 = int32(-2139062144)
	if (int32(16843008)-v326|v326)&v329 != v329 {
		v351 = v322
		v352 = v326
		v353 = v324
		goto L84
	} else {
		goto L95
	}
L91:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	*(*uint8)(unsafe.Add(mBase, uint32(v310))) = uint8(v311)
	if v311 == int32(0) {
		goto L83
	} else {
		goto L93
	}
L92:
	;
	v322 = v318
	v324 = v316
	goto L90
L93:
	;
	v315 = int32(1)
	v316 = v310 + v315
	v318 = v308 + v315
	if v318&int32(3) != 0 {
		v308 = v318
		v310 = v316
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v334 = v322
	v335 = v326
	v336 = v324
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v336))) = v335
	v338 = int32(4)
	v339 = v336 + v338
	v341 = v334 + v338
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	v346 = int32(-2139062144)
	if (int32(16843008)-v343|v343)&v346 == v346 {
		v334 = v341
		v335 = v343
		v336 = v339
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v351 = v341
	v352 = v343
	v353 = v339
	goto L84
L98:
	;
	goto L97
L99:
	;
	v360 = v351
	v362 = v353
	goto L100
L100:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)) = uint8(v363)
	v365 = int32(1)
	if v363 != 0 {
		v360 = v360 + v365
		v362 = v362 + v365
		goto L100
	} else {
		goto L102
	}
L101:
	;
	goto L83
L102:
	;
	goto L101
L103:
	;
	v377 = v373
	goto L53
L104:
	;
	if v382 != 0 {
		v166 = v382
		v173 = v377
		goto L51
	} else {
		goto L105
	}
L105:
	;
	goto L52
L106:
	;
	F_list_sort(m, v393, int32(1089))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	if v393 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v659 = int32(0)
	v664 = F_hash_search(m, l0, v20+int32(144), v659, v659)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L13
	} else {
		goto L150
	}
L109:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v408 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v412 = v20 + int32(234)
	v414 = v20 + int32(216)
	v416 = v20 + int32(196)
	v418 = v20 + int32(228)
	v433 = v7
	goto L111
L111:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v436+v433<<(uint(int32(2))%32))))
	v443 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L13
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	if v443 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v440 + int32(8)
	F_errmsg_internal(m, int32(_a_F_ResolveCminCmaxDuringDecoding_9), v20+int32(80))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L13
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(_a_F_ResolveCminCmaxDuringDecoding_0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v440 + int32(8)
	v467 = v20 + int32(240)
	v471 = F_pg_sprintf(m, v467, int32(_a_F_ResolveCminCmaxDuringDecoding_10), v20-int32(-64))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L119
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_12), int32(_a_F_ResolveCminCmaxDuringDecoding_13))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v474 = F_OpenTransientFile(m, v467, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L13
	} else {
		goto L120
	}
L120:
	;
	if v474 < int32(0) {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	v478 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v478
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = v478
	v482 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v482
	v484 = int32(_a_F_ResolveCminCmaxDuringDecoding_14)
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = int32(167772207)
	v491 = F_read(m, v474, v20+int32(204), int32(36))
	mBase = m.M
	v493 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v493))) = v482
	if v482 <= v491 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v634 = F_CloseTransientFile(m, v474)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L13
	} else {
		goto L146
	}
L123:
	;
	v501 = v491
	goto L126
L124:
	;
	goto L125
L125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L13
	} else {
		goto L142
	}
L126:
	;
	if v501 != int32(36) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	if v501 == int32(0) {
		goto L122
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v541
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v20)+204))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v543
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v418)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+4)) = uint16(v545)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v547
	v550 = v20 + int32(184)
	v551 = int32(0)
	v553 = F_hash_search(m, l0, v550, v551, v551)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L13
	} else {
		goto L137
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_15), v20+int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L13
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_16), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v578 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v578
	*(*int64)(unsafe.Add(mBase, uint32(v20)+192)) = v578
	v582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v582
	v584 = int32(_a_F_ResolveCminCmaxDuringDecoding_14)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = int32(167772207)
	v591 = F_read(m, v474, v20+int32(204), int32(36))
	mBase = m.M
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_ResolveCminCmaxDuringDecoding[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v593))) = v582
	if v582 <= v591 {
		v501 = v591
		goto L126
	} else {
		goto L141
	}
L137:
	;
	if v553 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v414)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v557
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v414)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+184)) = v559
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v412)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v416)+4)) = uint16(v561)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = v563
	v568 = F_hash_search(m, l0, v550, int32(1), v20+int32(180))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+180)))
	if v570 != 0 {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v553)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+20)) = v571
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v553)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+24)) = v573
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v553)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+28)) = v575
	goto L136
L141:
	;
	goto L127
L142:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_18), v20+int32(16))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_19), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L13
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
	if v634 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_pfree(m, v440)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	v639 = v433 + int32(1)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v639 < v640 {
		v433 = v639
		goto L111
	} else {
		goto L149
	}
L149:
	;
	goto L112
L150:
	;
	if v664 == int32(0) {
		v694 = v659
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v671 = v664
	goto L17
L152:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v671)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v685
	goto L154
L153:
	;
	goto L154
L154:
	;
	v687 = int32(1)
	if l5 == int32(0) {
		v694 = v687
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v671)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v690
	v694 = v687
	goto L4
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v194
	F_errmsg_internal(m, int32(_a_F_ResolveCminCmaxDuringDecoding_20), v20+int32(96))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_21), int32(_a_F_ResolveCminCmaxDuringDecoding_13))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L13
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_22), v20)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_23), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L13
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L13
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v20 + int32(240)
	F_errmsg(m, int32(_a_F_ResolveCminCmaxDuringDecoding_24), v20+int32(48))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L13
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_ResolveCminCmaxDuringDecoding_11), int32(_a_F_ResolveCminCmaxDuringDecoding_25), int32(_a_F_ResolveCminCmaxDuringDecoding_17))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L13
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RestoreSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_RestoreSnapshot[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = int32(2)
	v19 = v17 << (uint(v18) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = v20 << (uint(v18) % 32)
	v26 = F_MemoryContextAlloc(m, v16, v19+v22+int32(72))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v26)+64)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v14
		v33 = int32(1)
		v34 = v13 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+29)) = uint8(v34)
		v37 = v12 & v33
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v20
		v40 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v40
		*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v40
		*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = v40
		v49 = l0 + int32(24)
		if v17 == v40 {
		} else {
			v53 = v26 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v53
			if v19 == int32(0) {
			} else {
				base.MemoryCopy(m, v53, v49, v19)
			}
		}
		if v20 <= int32(0) {
		} else {
			v62 = v17 << (uint(int32(2)) % 32)
			v65 = v26 + v62 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v65
			if v22 == int32(0) {
			} else {
				base.MemoryCopy(m, v65, v49+v62, v22)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v26)+44)) = int64(0)
		v75 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v26)+30)) = uint8(v75)
		return v26
	}
}
func F_RestoreUserContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != int32(-1) {
		F_AtEOXact_GUC(m, int32(0), v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[0])) = v10
			*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[1])) = v9
			return
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[0])) = v10
		*(*int32)(unsafe.Add(mBase, _c_F_RestoreUserContext[1])) = v9
		return
	}
}
func F_r_R1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return base.B2i32(v2 <= v3)
}
func F_r_Step_1c_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v8 {
		v204 = v2
		return v204
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v6-int32(1)))))
		v16 = v14 - int32(100)
		v17 = int32(0)
		if base.B2i32(v16 == v17)|base.B2i32(v16 == int32(16)) == v17 {
			v204 = v2
			return v204
		} else {
			v27 = F_find_among_b(m, l0, int32(_a_F_r_Step_1c_1_0), int32(2), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v27 == int32(0) {
					v204 = v2
					return v204
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v33
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					if v33 < v35 {
						v204 = v2
						return v204
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v38 = int32(2)
						v40 = int32(0)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v42-v43 < v38 {
							v53 = v40
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v49 = F_memcmp(m, v46+v42-v38, int32(_a_F_r_Step_1c_1_1), v38)
							mBase = m.M
							if v49 != 0 {
								v53 = v40
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42 - v38
								v53 = int32(1)
							}
						}
						if v53 != 0 {
							v204 = v2
							return v204
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v55 = v37 - v33
							v56 = v54 - v55
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v56 <= v66 {
								v106 = int32(-1)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v56-int32(1)))))
								if int32(252) < v81 {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - int32(1)
									v103 = int32(0)
								} else {
									v83 = v81 - int32(97)
									if v83 < int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - int32(1)
										v103 = int32(0)
									} else {
										v86 = int32(1)
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v83)>>(uint(int32(3))%32)))+uint32(_c_F_r_Step_1c_1[0]))))
										if int32(base.Ui32(v90)>>(uint(v83&int32(7))%32))&v86 != 0 {
											v103 = v86
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v56 - int32(1)
											v103 = int32(0)
										}
									}
								}
								v106 = v103
							}
							if v106 != 0 {
								v204 = v2
								return v204
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v108 = v107 - v55
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
								switch v27 - int32(1) {
								case 0:
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v108 <= v112 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
										v127 = int32(2)
										v129 = int32(0)
										v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v108-v132 < v127 {
											v142 = v129
										} else {
											v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v138 = F_memcmp(m, v135+v108-v127, int32(_a_F_r_Step_1c_1_2), v127)
											mBase = m.M
											if v138 != 0 {
												v142 = v129
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v127
												v142 = int32(1)
											}
										}
										if v142 == int32(0) {
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
											v157 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v157 {
												v204 = int32(1)
											} else {
												v204 = v157
											}
											return v204
										} else {
											v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v146 < v145 {
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
												v157 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v157 {
													v204 = int32(1)
												} else {
													v204 = v157
												}
												return v204
											} else {
												v150 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_1_3))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v150 {
														v204 = int32(1)
													} else {
														v204 = v150
													}
													return v204
												}
											}
										}
									} else {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114+v108-int32(1)))))
										if v118 != int32(110) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
											v127 = int32(2)
											v129 = int32(0)
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v108-v132 < v127 {
												v142 = v129
											} else {
												v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v138 = F_memcmp(m, v135+v108-v127, int32(_a_F_r_Step_1c_1_2), v127)
												mBase = m.M
												if v138 != 0 {
													v142 = v129
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v127
													v142 = int32(1)
												}
											}
											if v142 == int32(0) {
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
												v157 = F_slice_del(m, l0)
												mBase = m.M
												if int32(0) <= v157 {
													v204 = int32(1)
												} else {
													v204 = v157
												}
												return v204
											} else {
												v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v146 < v145 {
													v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
													v157 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v157 {
														v204 = int32(1)
													} else {
														v204 = v157
													}
													return v204
												} else {
													v150 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_1_3))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v150 {
															v204 = int32(1)
														} else {
															v204 = v150
														}
														return v204
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - int32(1)
											v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v124 < v108 {
												v204 = v2
												return v204
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
												v127 = int32(2)
												v129 = int32(0)
												v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v108-v132 < v127 {
													v142 = v129
												} else {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v138 = F_memcmp(m, v135+v108-v127, int32(_a_F_r_Step_1c_1_2), v127)
													mBase = m.M
													if v138 != 0 {
														v142 = v129
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v127
														v142 = int32(1)
													}
												}
												if v142 == int32(0) {
													v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
													v157 = F_slice_del(m, l0)
													mBase = m.M
													if int32(0) <= v157 {
														v204 = int32(1)
													} else {
														v204 = v157
													}
													return v204
												} else {
													v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													if v146 < v145 {
														v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 - v55
														v157 = F_slice_del(m, l0)
														mBase = m.M
														if int32(0) <= v157 {
															v204 = int32(1)
														} else {
															v204 = v157
														}
														return v204
													} else {
														v150 = F_slice_from_s(m, l0, int32(1), int32(_a_F_r_Step_1c_1_3))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															if int32(0) <= v150 {
																v204 = int32(1)
															} else {
																v204 = v150
															}
															return v204
														}
													}
												}
											}
										}
									}
								case 1:
									v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v108 <= v160 {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
										v175 = int32(2)
										v177 = int32(0)
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v108-v180 < v175 {
											v190 = v177
										} else {
											v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v186 = F_memcmp(m, v183+v108-v175, int32(_a_F_r_Step_1c_1_4), v175)
											mBase = m.M
											if v186 != 0 {
												v190 = v177
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v175
												v190 = int32(1)
											}
										}
										if v190 != 0 {
											v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v191 <= v192 {
												v204 = v2
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
												v197 = F_slice_del(m, l0)
												mBase = m.M
												if v197 < int32(0) {
													v204 = v197
												} else {
													v204 = int32(1)
												}
											}
										} else {
											v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
											v197 = F_slice_del(m, l0)
											mBase = m.M
											if v197 < int32(0) {
												v204 = v197
											} else {
												v204 = int32(1)
											}
										}
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v108-int32(1)))))
										if v166 != int32(104) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
											v175 = int32(2)
											v177 = int32(0)
											v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v108-v180 < v175 {
												v190 = v177
											} else {
												v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v186 = F_memcmp(m, v183+v108-v175, int32(_a_F_r_Step_1c_1_4), v175)
												mBase = m.M
												if v186 != 0 {
													v190 = v177
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v175
													v190 = int32(1)
												}
											}
											if v190 != 0 {
												v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v191 <= v192 {
													v204 = v2
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
													v197 = F_slice_del(m, l0)
													mBase = m.M
													if v197 < int32(0) {
														v204 = v197
													} else {
														v204 = int32(1)
													}
												}
											} else {
												v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
												v197 = F_slice_del(m, l0)
												mBase = m.M
												if v197 < int32(0) {
													v204 = v197
												} else {
													v204 = int32(1)
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - int32(1)
											v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v172 < v108 {
												v204 = v2
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
												v175 = int32(2)
												v177 = int32(0)
												v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v108-v180 < v175 {
													v190 = v177
												} else {
													v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v186 = F_memcmp(m, v183+v108-v175, int32(_a_F_r_Step_1c_1_4), v175)
													mBase = m.M
													if v186 != 0 {
														v190 = v177
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 - v175
														v190 = int32(1)
													}
												}
												if v190 != 0 {
													v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													if v191 <= v192 {
														v204 = v2
													} else {
														v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
														v197 = F_slice_del(m, l0)
														mBase = m.M
														if v197 < int32(0) {
															v204 = v197
														} else {
															v204 = int32(1)
														}
													}
												} else {
													v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v194 - v55
													v197 = F_slice_del(m, l0)
													mBase = m.M
													if v197 < int32(0) {
														v204 = v197
													} else {
														v204 = int32(1)
													}
												}
											}
										}
									}
									return v204
								default:
									v204 = int32(1)
									return v204
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_r_e_ending_1(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v2)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v9 {
		v89 = v2
		return v89
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v7-int32(1)))))
		if v15 != int32(101) {
			v89 = v2
			return v89
		} else {
			v19 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v7 <= v22 {
				v89 = v2
				return v89
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v32 <= v33 {
					v73 = int32(-1)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v32-int32(1)))))
					if int32(232) < v48 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 - int32(1)
						v70 = int32(0)
					} else {
						v50 = v48 - int32(97)
						if v50 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 - int32(1)
							v70 = int32(0)
						} else {
							v53 = int32(1)
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v50)>>(uint(int32(3))%32)))+uint32(_c_F_r_e_ending_1[0]))))
							if int32(base.Ui32(v57)>>(uint(v50&int32(7))%32))&v53 != 0 {
								v70 = v53
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 - int32(1)
								v70 = int32(0)
							}
						}
					}
					v73 = v70
				}
				if v73 != 0 {
					v89 = v2
					return v89
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74 + (v19 - v24)
					v78 = F_slice_del(m, l0)
					mBase = m.M
					if v78 < int32(0) {
						v89 = v78
						return v89
					} else {
						v81 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v81)
						v83 = F_r_undouble_1(m, l0)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v89 = v83
							return v89
						}
					}
				}
			}
		}
	}
}
func F_r_en_ending_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5 < v6 {
		v138 = v2
		return v138
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v16 <= v17 {
			v57 = int32(-1)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v16-int32(1)))))
			if int32(232) < v32 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 - int32(1)
				v54 = int32(0)
			} else {
				v34 = v32 - int32(97)
				if v34 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 - int32(1)
					v54 = int32(0)
				} else {
					v37 = int32(1)
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v34)>>(uint(int32(3))%32)))+uint32(_c_F_r_en_ending_1[0]))))
					if int32(base.Ui32(v41)>>(uint(v34&int32(7))%32))&v37 != 0 {
						v54 = v37
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v16 - int32(1)
						v54 = int32(0)
					}
				}
			}
			v57 = v54
		}
		if v57 != 0 {
			v138 = v2
			return v138
		} else {
			v58 = v5 - v8
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v60 = v58 + v59
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
			v62 = int32(3)
			v64 = int32(0)
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v60-v67 < v62 {
				v77 = v64
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v73 = F_memcmp(m, v70+v60-v62, int32(_a_F_r_en_ending_1_0), v62)
				mBase = m.M
				if v73 != 0 {
					v77 = v64
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60 - v62
					v77 = int32(1)
				}
			}
			if v77 != 0 {
				v138 = v2
				return v138
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v78 + v58
				v81 = F_slice_del(m, l0)
				mBase = m.M
				if v81 < int32(0) {
					v138 = v81
					return v138
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v86 = v84 - int32(1)
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v86 <= v87 {
						return int32(0)
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v86))))
						if v93&int32(224) != int32(96) {
							return int32(0)
						} else {
							v100 = int32(0)
							if int32(1)<<(uint(v93)%32)&int32(_a_F_r_en_ending_1_1) == v100 {
								v138 = v100
								return v138
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v111 = F_find_among_b(m, l0, int32(_a_F_r_en_ending_1_2), int32(3), int32(0))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									if v111 == int32(0) {
										v138 = v100
									} else {
										v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v119 = v117 + (v84 - v107)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v119
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119
										v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v119 <= v122 {
											v138 = v100
										} else {
											v124 = int32(1)
											v125 = v119 - v124
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v125
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
											v129 = F_slice_del(m, l0)
											mBase = m.M
											if int32(0) <= v129 {
												v135 = v124
											} else {
												v135 = v129 >> (uint(int32(31)) % 32) & v129
											}
											v138 = v135
										}
									}
									return v138
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_r_remove_suffix_2(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14366(m, l0, int32(_a_F_r_remove_suffix_2_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_r_shortv_2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 <= v19 {
		v128 = int32(-1)
		v135 = v128
	} else {
		v36 = int32(1)
		v37 = v4 - v36
		v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+v37))))
		v41 = v39 & int32(255)
		if base.B2i32(v37 == v19)|base.B2i32(int32(0) <= v39) != 0 {
			v99 = v41
			v103 = v36
		} else {
			v48 = v41 & int32(63)
			v50 = v4 - int32(2)
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v50))))
			v54 = v52 << (uint(int32(6)) % 32)
			if base.B2i32(v50 != v19)&base.B2i32(base.Ui32(v52) < base.Ui32(int32(192))) == int32(0) {
				v99 = v54&int32(1984) | v48
				v103 = int32(2)
			} else {
				v67 = v54&int32(4032) | v48
				v69 = v4 - int32(3)
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v69))))
				if base.B2i32(v69 != v19)&base.B2i32(base.Ui32(v71) < base.Ui32(int32(224))) == int32(0) {
					v99 = v71<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v67
					v103 = int32(3)
				} else {
					v89 = int32(4)
					v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v20-v89))))
					v99 = v71<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v91&int32(7)<<(uint(int32(18))%32) | v67
					v103 = v89
				}
			}
		}
		if int32(121) < v99 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - v103
			v128 = int32(0)
			v135 = v128
		} else {
			v105 = v99 - int32(89)
			if v105 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - v103
				v128 = int32(0)
				v135 = v128
			} else {
				v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v105)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[0]))))
				if int32(base.Ui32(v111)>>(uint(v105&int32(7))%32))&int32(1) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4 - v103
					v128 = int32(0)
					v135 = v128
				} else {
					v135 = v103
				}
			}
		}
	}
	if v135 != 0 {
		v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v398 = v5 - v4
		v399 = v397 - v398
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399
		v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v399 <= v414 {
			v523 = int32(-1)
			v530 = v523
		} else {
			v431 = int32(1)
			v432 = v399 - v431
			v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v432))))
			v436 = v434 & int32(255)
			if base.B2i32(v432 == v414)|base.B2i32(int32(0) <= v434) != 0 {
				v494 = v436
				v498 = v431
			} else {
				v443 = v436 & int32(63)
				v445 = v399 - int32(2)
				v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v445))))
				v449 = v447 << (uint(int32(6)) % 32)
				if base.B2i32(v445 != v414)&base.B2i32(base.Ui32(v447) < base.Ui32(int32(192))) == int32(0) {
					v494 = v449&int32(1984) | v443
					v498 = int32(2)
				} else {
					v462 = v449&int32(4032) | v443
					v464 = v399 - int32(3)
					v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v464))))
					if base.B2i32(v464 != v414)&base.B2i32(base.Ui32(v466) < base.Ui32(int32(224))) == int32(0) {
						v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v462
						v498 = int32(3)
					} else {
						v484 = int32(4)
						v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v415-v484))))
						v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v486&int32(7)<<(uint(int32(18))%32) | v462
						v498 = v484
					}
				}
			}
			if int32(121) < v494 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
				v523 = int32(0)
				v530 = v523
			} else {
				v500 = v494 - int32(97)
				if v500 < int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
					v523 = int32(0)
					v530 = v523
				} else {
					v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v500)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
					if int32(base.Ui32(v506)>>(uint(v500&int32(7))%32))&int32(1) == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
						v523 = int32(0)
						v530 = v523
					} else {
						v530 = v498
					}
				}
			}
		}
		if v530 != 0 {
			v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v664 = v663 - v398
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
			v666 = int32(4)
			v668 = int32(0)
			v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v664-v671 < v666 {
				v681 = v668
			} else {
				v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
				mBase = m.M
				if v677 != 0 {
					v681 = v668
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
					v681 = int32(1)
				}
			}
			if v681 != 0 {
				return int32(1)
			} else {
				return int32(0)
			}
		} else {
			v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v543 <= v544 {
				v652 = int32(-1)
				v659 = v652
			} else {
				v561 = int32(1)
				v562 = v543 - v561
				v564 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v562))))
				v566 = v564 & int32(255)
				if base.B2i32(v562 == v544)|base.B2i32(int32(0) <= v564) != 0 {
					v624 = v566
					v628 = v561
				} else {
					v573 = v566 & int32(63)
					v575 = v543 - int32(2)
					v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v575))))
					v579 = v577 << (uint(int32(6)) % 32)
					if base.B2i32(v575 != v544)&base.B2i32(base.Ui32(v577) < base.Ui32(int32(192))) == int32(0) {
						v624 = v579&int32(1984) | v573
						v628 = int32(2)
					} else {
						v592 = v579&int32(4032) | v573
						v594 = v543 - int32(3)
						v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v594))))
						if base.B2i32(v594 != v544)&base.B2i32(base.Ui32(v596) < base.Ui32(int32(224))) == int32(0) {
							v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v592
							v628 = int32(3)
						} else {
							v614 = int32(4)
							v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v545-v614))))
							v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v616&int32(7)<<(uint(int32(18))%32) | v592
							v628 = v614
						}
					}
				}
				if int32(121) < v624 {
					v659 = v628
				} else {
					v630 = v624 - int32(97)
					if v630 < int32(0) {
						v659 = v628
					} else {
						v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v630)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v636)>>(uint(v630&int32(7))%32))&int32(1) == int32(0) {
							v659 = v628
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - v628
							v652 = int32(0)
							v659 = v652
						}
					}
				}
			}
			if v659 != 0 {
				v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v664 = v663 - v398
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
				v666 = int32(4)
				v668 = int32(0)
				v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v664-v671 < v666 {
					v681 = v668
				} else {
					v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
					mBase = m.M
					if v677 != 0 {
						v681 = v668
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
						v681 = int32(1)
					}
				}
				if v681 != 0 {
					return int32(1)
				} else {
					return int32(0)
				}
			} else {
				v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v660 <= v661 {
					return int32(1)
				} else {
					v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v664 = v663 - v398
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
					v666 = int32(4)
					v668 = int32(0)
					v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v664-v671 < v666 {
						v681 = v668
					} else {
						v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
						mBase = m.M
						if v677 != 0 {
							v681 = v668
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
							v681 = int32(1)
						}
					}
					if v681 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v148 <= v149 {
			v257 = int32(-1)
			v264 = v257
		} else {
			v166 = int32(1)
			v167 = v148 - v166
			v169 = int32(*(*int8)(unsafe.Add(mBase, uint32(v150+v167))))
			v171 = v169 & int32(255)
			if base.B2i32(v167 == v149)|base.B2i32(int32(0) <= v169) != 0 {
				v229 = v171
				v233 = v166
			} else {
				v178 = v171 & int32(63)
				v180 = v148 - int32(2)
				v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v180))))
				v184 = v182 << (uint(int32(6)) % 32)
				if base.B2i32(v180 != v149)&base.B2i32(base.Ui32(v182) < base.Ui32(int32(192))) == int32(0) {
					v229 = v184&int32(1984) | v178
					v233 = int32(2)
				} else {
					v197 = v184&int32(4032) | v178
					v199 = v148 - int32(3)
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v199))))
					if base.B2i32(v199 != v149)&base.B2i32(base.Ui32(v201) < base.Ui32(int32(224))) == int32(0) {
						v229 = v201<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v197
						v233 = int32(3)
					} else {
						v219 = int32(4)
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148+v150-v219))))
						v229 = v201<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v221&int32(7)<<(uint(int32(18))%32) | v197
						v233 = v219
					}
				}
			}
			if int32(121) < v229 {
				v264 = v233
			} else {
				v235 = v229 - int32(97)
				if v235 < int32(0) {
					v264 = v233
				} else {
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v235)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
					if int32(base.Ui32(v241)>>(uint(v235&int32(7))%32))&int32(1) == int32(0) {
						v264 = v233
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v148 - v233
						v257 = int32(0)
						v264 = v257
					}
				}
			}
		}
		if v264 != 0 {
			v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v398 = v5 - v4
			v399 = v397 - v398
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399
			v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v399 <= v414 {
				v523 = int32(-1)
				v530 = v523
			} else {
				v431 = int32(1)
				v432 = v399 - v431
				v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v432))))
				v436 = v434 & int32(255)
				if base.B2i32(v432 == v414)|base.B2i32(int32(0) <= v434) != 0 {
					v494 = v436
					v498 = v431
				} else {
					v443 = v436 & int32(63)
					v445 = v399 - int32(2)
					v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v445))))
					v449 = v447 << (uint(int32(6)) % 32)
					if base.B2i32(v445 != v414)&base.B2i32(base.Ui32(v447) < base.Ui32(int32(192))) == int32(0) {
						v494 = v449&int32(1984) | v443
						v498 = int32(2)
					} else {
						v462 = v449&int32(4032) | v443
						v464 = v399 - int32(3)
						v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v464))))
						if base.B2i32(v464 != v414)&base.B2i32(base.Ui32(v466) < base.Ui32(int32(224))) == int32(0) {
							v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v462
							v498 = int32(3)
						} else {
							v484 = int32(4)
							v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v415-v484))))
							v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v486&int32(7)<<(uint(int32(18))%32) | v462
							v498 = v484
						}
					}
				}
				if int32(121) < v494 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
					v523 = int32(0)
					v530 = v523
				} else {
					v500 = v494 - int32(97)
					if v500 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
						v523 = int32(0)
						v530 = v523
					} else {
						v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v500)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v506)>>(uint(v500&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
							v523 = int32(0)
							v530 = v523
						} else {
							v530 = v498
						}
					}
				}
			}
			if v530 != 0 {
				v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v664 = v663 - v398
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
				v666 = int32(4)
				v668 = int32(0)
				v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v664-v671 < v666 {
					v681 = v668
				} else {
					v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
					mBase = m.M
					if v677 != 0 {
						v681 = v668
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
						v681 = int32(1)
					}
				}
				if v681 != 0 {
					return int32(1)
				} else {
					return int32(0)
				}
			} else {
				v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v543 <= v544 {
					v652 = int32(-1)
					v659 = v652
				} else {
					v561 = int32(1)
					v562 = v543 - v561
					v564 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v562))))
					v566 = v564 & int32(255)
					if base.B2i32(v562 == v544)|base.B2i32(int32(0) <= v564) != 0 {
						v624 = v566
						v628 = v561
					} else {
						v573 = v566 & int32(63)
						v575 = v543 - int32(2)
						v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v575))))
						v579 = v577 << (uint(int32(6)) % 32)
						if base.B2i32(v575 != v544)&base.B2i32(base.Ui32(v577) < base.Ui32(int32(192))) == int32(0) {
							v624 = v579&int32(1984) | v573
							v628 = int32(2)
						} else {
							v592 = v579&int32(4032) | v573
							v594 = v543 - int32(3)
							v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v594))))
							if base.B2i32(v594 != v544)&base.B2i32(base.Ui32(v596) < base.Ui32(int32(224))) == int32(0) {
								v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v592
								v628 = int32(3)
							} else {
								v614 = int32(4)
								v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v545-v614))))
								v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v616&int32(7)<<(uint(int32(18))%32) | v592
								v628 = v614
							}
						}
					}
					if int32(121) < v624 {
						v659 = v628
					} else {
						v630 = v624 - int32(97)
						if v630 < int32(0) {
							v659 = v628
						} else {
							v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v630)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
							if int32(base.Ui32(v636)>>(uint(v630&int32(7))%32))&int32(1) == int32(0) {
								v659 = v628
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - v628
								v652 = int32(0)
								v659 = v652
							}
						}
					}
				}
				if v659 != 0 {
					v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v664 = v663 - v398
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
					v666 = int32(4)
					v668 = int32(0)
					v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v664-v671 < v666 {
						v681 = v668
					} else {
						v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
						mBase = m.M
						if v677 != 0 {
							v681 = v668
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
							v681 = int32(1)
						}
					}
					if v681 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				} else {
					v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v660 <= v661 {
						return int32(1)
					} else {
						v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v664 = v663 - v398
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
						v666 = int32(4)
						v668 = int32(0)
						v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v664-v671 < v666 {
							v681 = v668
						} else {
							v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
							mBase = m.M
							if v677 != 0 {
								v681 = v668
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
								v681 = int32(1)
							}
						}
						if v681 != 0 {
							return int32(1)
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v277 <= v278 {
				v387 = int32(-1)
				v394 = v387
			} else {
				v295 = int32(1)
				v296 = v277 - v295
				v298 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279+v296))))
				v300 = v298 & int32(255)
				if base.B2i32(v296 == v278)|base.B2i32(int32(0) <= v298) != 0 {
					v358 = v300
					v362 = v295
				} else {
					v307 = v300 & int32(63)
					v309 = v277 - int32(2)
					v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v309))))
					v313 = v311 << (uint(int32(6)) % 32)
					if base.B2i32(v309 != v278)&base.B2i32(base.Ui32(v311) < base.Ui32(int32(192))) == int32(0) {
						v358 = v313&int32(1984) | v307
						v362 = int32(2)
					} else {
						v326 = v313&int32(4032) | v307
						v328 = v277 - int32(3)
						v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v328))))
						if base.B2i32(v328 != v278)&base.B2i32(base.Ui32(v330) < base.Ui32(int32(224))) == int32(0) {
							v358 = v330<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v326
							v362 = int32(3)
						} else {
							v348 = int32(4)
							v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277+v279-v348))))
							v358 = v330<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v350&int32(7)<<(uint(int32(18))%32) | v326
							v362 = v348
						}
					}
				}
				if int32(121) < v358 {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277 - v362
					v387 = int32(0)
					v394 = v387
				} else {
					v364 = v358 - int32(97)
					if v364 < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277 - v362
						v387 = int32(0)
						v394 = v387
					} else {
						v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v364)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
						if int32(base.Ui32(v370)>>(uint(v364&int32(7))%32))&int32(1) == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v277 - v362
							v387 = int32(0)
							v394 = v387
						} else {
							v394 = v362
						}
					}
				}
			}
			if v394 == int32(0) {
				return int32(1)
			} else {
				v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v398 = v5 - v4
				v399 = v397 - v398
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399
				v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v399 <= v414 {
					v523 = int32(-1)
					v530 = v523
				} else {
					v431 = int32(1)
					v432 = v399 - v431
					v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v415+v432))))
					v436 = v434 & int32(255)
					if base.B2i32(v432 == v414)|base.B2i32(int32(0) <= v434) != 0 {
						v494 = v436
						v498 = v431
					} else {
						v443 = v436 & int32(63)
						v445 = v399 - int32(2)
						v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v445))))
						v449 = v447 << (uint(int32(6)) % 32)
						if base.B2i32(v445 != v414)&base.B2i32(base.Ui32(v447) < base.Ui32(int32(192))) == int32(0) {
							v494 = v449&int32(1984) | v443
							v498 = int32(2)
						} else {
							v462 = v449&int32(4032) | v443
							v464 = v399 - int32(3)
							v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v464))))
							if base.B2i32(v464 != v414)&base.B2i32(base.Ui32(v466) < base.Ui32(int32(224))) == int32(0) {
								v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v462
								v498 = int32(3)
							} else {
								v484 = int32(4)
								v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399+v415-v484))))
								v494 = v466<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v486&int32(7)<<(uint(int32(18))%32) | v462
								v498 = v484
							}
						}
					}
					if int32(121) < v494 {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
						v523 = int32(0)
						v530 = v523
					} else {
						v500 = v494 - int32(97)
						if v500 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
							v523 = int32(0)
							v530 = v523
						} else {
							v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v500)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
							if int32(base.Ui32(v506)>>(uint(v500&int32(7))%32))&int32(1) == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v399 - v498
								v523 = int32(0)
								v530 = v523
							} else {
								v530 = v498
							}
						}
					}
				}
				if v530 != 0 {
					v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v664 = v663 - v398
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
					v666 = int32(4)
					v668 = int32(0)
					v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v664-v671 < v666 {
						v681 = v668
					} else {
						v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
						mBase = m.M
						if v677 != 0 {
							v681 = v668
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
							v681 = int32(1)
						}
					}
					if v681 != 0 {
						return int32(1)
					} else {
						return int32(0)
					}
				} else {
					v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v543 <= v544 {
						v652 = int32(-1)
						v659 = v652
					} else {
						v561 = int32(1)
						v562 = v543 - v561
						v564 = int32(*(*int8)(unsafe.Add(mBase, uint32(v545+v562))))
						v566 = v564 & int32(255)
						if base.B2i32(v562 == v544)|base.B2i32(int32(0) <= v564) != 0 {
							v624 = v566
							v628 = v561
						} else {
							v573 = v566 & int32(63)
							v575 = v543 - int32(2)
							v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v575))))
							v579 = v577 << (uint(int32(6)) % 32)
							if base.B2i32(v575 != v544)&base.B2i32(base.Ui32(v577) < base.Ui32(int32(192))) == int32(0) {
								v624 = v579&int32(1984) | v573
								v628 = int32(2)
							} else {
								v592 = v579&int32(4032) | v573
								v594 = v543 - int32(3)
								v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v594))))
								if base.B2i32(v594 != v544)&base.B2i32(base.Ui32(v596) < base.Ui32(int32(224))) == int32(0) {
									v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_0) | v592
									v628 = int32(3)
								} else {
									v614 = int32(4)
									v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543+v545-v614))))
									v624 = v596<<(uint(int32(12))%32)&int32(_a_F_r_shortv_2_1) | v616&int32(7)<<(uint(int32(18))%32) | v592
									v628 = v614
								}
							}
						}
						if int32(121) < v624 {
							v659 = v628
						} else {
							v630 = v624 - int32(97)
							if v630 < int32(0) {
								v659 = v628
							} else {
								v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v630)>>(uint(int32(3))%32)))+uint32(_c_F_r_shortv_2[1]))))
								if int32(base.Ui32(v636)>>(uint(v630&int32(7))%32))&int32(1) == int32(0) {
									v659 = v628
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - v628
									v652 = int32(0)
									v659 = v652
								}
							}
						}
					}
					if v659 != 0 {
						v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v664 = v663 - v398
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
						v666 = int32(4)
						v668 = int32(0)
						v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v664-v671 < v666 {
							v681 = v668
						} else {
							v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
							mBase = m.M
							if v677 != 0 {
								v681 = v668
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
								v681 = int32(1)
							}
						}
						if v681 != 0 {
							return int32(1)
						} else {
							return int32(0)
						}
					} else {
						v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v660 <= v661 {
							return int32(1)
						} else {
							v663 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v664 = v663 - v398
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664
							v666 = int32(4)
							v668 = int32(0)
							v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v664-v671 < v666 {
								v681 = v668
							} else {
								v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v677 = F_memcmp(m, v674+v664-v666, int32(_a_F_r_shortv_2_2), v666)
								mBase = m.M
								if v677 != 0 {
									v681 = v668
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 - v666
									v681 = int32(1)
								}
							}
							if v681 != 0 {
								return int32(1)
							} else {
								return int32(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_r_undouble_4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L4
L1:
	;
	return v131
L2:
	;
	if v58 < int32(0) {
		v131 = v2
		goto L1
	} else {
		goto L21
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v13 = v5
	v15 = int32(1)
	goto L9
L8:
	;
	v58 = v40
	goto L2
L9:
	;
	if v13 <= v6 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v58 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v20 = v13 - int32(1)
	v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v4+v20))))
	if base.B2i32(int32(0) <= v22)|base.B2i32(v20 <= v6) != 0 {
		v40 = v20
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = int32(1)
	if v44 < v15 {
		v13 = v40
		v15 = v15 - v44
		goto L9
	} else {
		goto L20
	}
L15:
	;
	v28 = v20
	goto L16
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v28))))
	if base.Ui32(int32(191)) < base.Ui32(v33) {
		v40 = v28
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v40 = v6
	goto L14
L18:
	;
	v37 = v28 - int32(1)
	if v6 < v37 {
		v28 = v37
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	goto L10
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v58
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L24
L22:
	;
	if v116 < int32(0) {
		v131 = v2
		goto L1
	} else {
		goto L41
	}
L24:
	;
	goto L25
L25:
	;
	goto L26
L26:
	;
	v71 = v58
	v73 = int32(1)
	goto L29
L28:
	;
	v116 = v98
	goto L22
L29:
	;
	if v71 <= v64 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v116 = int32(-1)
	goto L22
L32:
	;
	goto L33
L33:
	;
	v78 = v71 - int32(1)
	v80 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63+v78))))
	if base.B2i32(int32(0) <= v80)|base.B2i32(v78 <= v64) != 0 {
		v98 = v78
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v102 = int32(1)
	if v102 < v73 {
		v71 = v98
		v73 = v73 - v102
		goto L29
	} else {
		goto L40
	}
L35:
	;
	v86 = v78
	goto L36
L36:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v86))))
	if base.Ui32(int32(191)) < base.Ui32(v91) {
		v98 = v86
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v98 = v64
	goto L34
L38:
	;
	v95 = v86 - int32(1)
	if v64 < v95 {
		v86 = v95
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L30
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v122 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v122 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v128 = int32(1)
	goto L44
L43:
	;
	v128 = v122 >> (uint(int32(31)) % 32) & v122
	goto L44
L44:
	;
	v131 = v128
	goto L1
}
func F_radix_sort_recursive(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v53 int32
	_ = v53
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	v17 = m.G0
	v19 = v17 - int32(2336)
	m.G0 = v19
	base.MemoryFill(m, v19+int32(288), int32(0), int32(2048))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v28 == int32(202) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	v53 = l0
	v64 = int64(0)
	goto L6
L2:
	;
	v39 = v26 ^ int64(-9223372036854775807-1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if v28 != int32(199) {
		v39 = v26
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = v26&int64(4294967295) ^ int64(2147483648)
	goto L1
L6:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v69 == int32(202) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v108 = int32(0)
	v112 = v108
	v115 = v108
	v116 = v108
	goto L19
L8:
	;
	v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	v84 = v80 ^ (int64(0) - v82)
	v85 = int64(base.Ui64(v84) >> (uint(base.I64_extend_i32_u(l2<<(uint(int32(3))%32)^int32(56))) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+17)) = uint8(v85)
	v94 = v19 + int32(288) + base.I32_wrap_i64(v85)&int32(255)<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_radix_sort_recursive[0]))
	if v100 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v80 = v68 ^ int64(-9223372036854775807-1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v69 != int32(199) {
		v80 = v68
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v80 = v68&int64(4294967295) ^ int64(2147483648)
	goto L8
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v104 = v84 ^ (v39 ^ (int64(0) - v44)) | v64
	v106 = v53 + int32(24)
	if base.Ui32(v106) < base.Ui32(l0+l1*int32(24)) {
		v53 = v106
		v64 = v104
		goto L6
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	goto L7
L19:
	;
	v131 = v19 + int32(288) + v115<<(uint(int32(3))%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v132 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if int32(1) < v161 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v112
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(32)+v116))) = uint8(v115)
	v141 = v112 + v132
	v142 = v116 + int32(1)
	goto L23
L22:
	;
	v141 = v112
	v142 = v116
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v141
	v147 = v115 | int32(1)
	v150 = v19 + int32(288) + v147<<(uint(int32(3))%32)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v151 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v141
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(32)+v142))) = uint8(v147)
	v160 = v141 + v151
	v161 = v142 + int32(1)
	goto L26
L25:
	;
	v160 = v141
	v161 = v142
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v160
	v164 = v115 + int32(2)
	if v164 != int32(256) {
		v112 = v160
		v115 = v164
		v116 = v161
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	v171 = v161
	v176 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	if v161 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L31:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(32)+v176))))
	v194 = v19 + int32(288) + v191<<(uint(int32(3))%32)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if base.Ui32(v195) < base.Ui32(v196) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v198 = int32(24)
	v205 = l0 + v195*v198
	goto L36
L34:
	;
	v275 = v195
	goto L35
L35:
	;
	v277 = v171 - base.B2i32(v275 == v196)
	v279 = v176 + int32(1)
	if v279 != v161 {
		v171 = v277
		v176 = v279
		goto L31
	} else {
		goto L43
	}
L36:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+17)))
	v225 = v19 + int32(288) + v222<<(uint(int32(3))%32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v226 + int32(1)
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v205)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v230
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v205)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v205)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v234
	v238 = l0 + v226*int32(24)
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v238)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+16)) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v238)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v205)+8)) = v241
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v238)))
	*(*int64)(unsafe.Add(mBase, uint32(v205))) = v243
	v245 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v245
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v238))) = v247
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v238)+8)) = v249
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_radix_sort_recursive[0]))
	if v252 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v275 = v258
	goto L35
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L16
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v256 = v205 + int32(24)
	if base.Ui32(v256) < base.Ui32(l0+v196*v198) {
		v205 = v256
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	goto L37
L43:
	;
	v283 = int32(0)
	if base.B2i32(v277 < int32(2)) == v283 {
		v171 = v161
		v176 = v283
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	m.G0 = v19 + int32(2336)
	return
L46:
	;
	if v104 == int64(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v161 <= int32(0) {
		goto L45
	} else {
		goto L52
	}
L48:
	;
	v318 = v19 + int32(32)
	v324 = int32(0)
	v327 = v318
	v329 = l0
	goto L53
L49:
	;
	v311 = int32(8)
	goto L51
L50:
	;
	v311 = int32(base.Ui32(base.I32_wrap_i64(base.I64_clz(v104))) >> (uint(int32(3)) % 32))
	goto L51
L51:
	;
	v316 = v311
	goto L48
L52:
	;
	v316 = l2 + int32(1)
	goto L48
L53:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(288)+v341<<(uint(int32(3))%32))+4))
	v348 = v345 - v324
	if base.Ui32(v348) < base.Ui32(int32(2)) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L45
L55:
	;
	v366 = v327 + int32(1)
	if base.Ui32(v366) < base.Ui32(v318+v161) {
		v324 = v345
		v327 = v366
		v329 = l0 + v345*int32(24)
		goto L53
	} else {
		goto L67
	}
L56:
	;
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v316)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if base.Ui32(v348) <= base.Ui32(int32(39)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	if v360 != 0 {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_qsort_tuple(m, v329, v348, v355, l3)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L16
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_radix_sort_recursive(m, v329, v348, v316, l3)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
	} else {
		goto L64
	}
L63:
	;
	goto L55
L64:
	;
	goto L55
L65:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	F_qsort_tuple(m, v329, v348, v361, l3)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	goto L55
L67:
	;
	goto L54
}
func F_rainbow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
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
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	if l2 != int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(24)
	v19 = v13 + v14*v15 + v15
	if base.Ui32(v19) <= base.Ui32(v13) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_rainbow[0]))
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L5:
	;
	v27 = v13
	v28 = int32(0)
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L1
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	if v33&int32(1)|v33&int32(2) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v122 = v27 + int32(24)
	if base.Ui32(v122) < base.Ui32(v19) {
		v27 = v122
		v28 = v28 + int32(1)
		goto L6
	} else {
		goto L39
	}
L10:
	;
	v39 = int32(_a_F_rainbow_0)
	v40 = v28 & v39
	if v40 == l2&v39 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)))
	if v44 == v40 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_rainbow[0]))
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v50 <= v51 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	F_createarc(m, l0, int32(112), base.I32_extend16_s(v28), l3, l4)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L16
	} else {
		goto L38
	}
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v53 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v74 == int32(0) {
		goto L18
	} else {
		goto L30
	}
L22:
	;
	v61 = v53
	goto L23
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v66 != l4 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L18
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v73 != 0 {
		v61 = v73
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
	if v68 != v40 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v70 == int32(112) {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L24
L30:
	;
	v82 = v74
	goto L31
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if v87 != l3 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L18
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v94 != 0 {
		v82 = v94
		goto L31
	} else {
		goto L37
	}
L34:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)))
	if v89 != v40 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v91 == int32(112) {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	goto L32
L38:
	;
	goto L9
L39:
	;
	goto L7
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v128 <= v129 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	F_createarc(m, l0, int32(112), int32(-2), l3, l4)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L64
	}
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v131 == int32(0) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v153 == int32(0) {
		goto L44
	} else {
		goto L56
	}
L48:
	;
	v139 = v131
	goto L49
L49:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	if v144 != l4 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L44
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
	if v152 != 0 {
		v139 = v152
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
	if v146 != int32(_a_F_rainbow_1) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v149 == int32(112) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	goto L50
L56:
	;
	v161 = v153
	goto L57
L57:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	if v166 != l3 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L44
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v174 != 0 {
		v161 = v174
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+4)))
	if v168 != int32(_a_F_rainbow_1) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v171 == int32(112) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L58
L64:
	;
	goto L1
}
func F_recordMultipleDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v239 int32
	_ = v239
	v5 = int32(0)
	if l2 <= v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_recordMultipleDependencies[0]))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v28 = int32(2340)
	if base.Ui32(v28) <= base.Ui32(l2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = v28
	goto L8
L7:
	;
	v31 = l2
	goto L8
L8:
	;
	v32 = F_palloc_mul(m, int32(4), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v35 = l1
	v40 = v5
	v42 = v5
	v44 = v5
	goto L11
L10:
	;
	if v188 != 0 {
		goto L43
	} else {
		goto L44
	}
L11:
	;
	v50 = v35
	v54 = int32(0)
	v57 = v42
	v59 = v44
	goto L13
L12:
	;
	if v167 <= int32(0) {
		v188 = v40
		v189 = v168
		goto L10
	} else {
		goto L37
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	goto L16
L14:
	;
	goto L12
L15:
	;
	v174 = v59 + int32(1)
	if v174 != l2 {
		v50 = v50 + int32(12)
		v54 = v167
		v57 = v168
		v59 = v174
		goto L13
	} else {
		goto L36
	}
L16:
	;
	if base.B2i32(base.B2i32(v63 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_recordMultipleDependencies_0)) < base.Ui32(v64)) == int32(0))&((base.B2i32(v63 != int32(2615))|base.B2i32(v64 != int32(2200)))&base.B2i32(v63 != int32(1262))) != 0 {
		v167 = v54
		v168 = v57
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	F_dependencyLockAndCheckObject(m, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v31 <= v57 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	m.T0[v104].(func(*base.Module, int32))(m, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L24
	}
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v32+v54<<(uint(int32(2))%32))))
	v101 = v90
	v102 = v57
	goto L19
L21:
	;
	goto L22
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v96 = F_MakeSingleTupleTableSlot(m, v94, int32(_a_F_recordMultipleDependencies_1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+v54<<(uint(int32(2))%32)))) = v96
	v101 = v96
	v102 = v57 + int32(1)
	goto L19
L24:
	;
	v109 = v32 + v54<<(uint(int32(2))%32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v50))))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+24)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v50)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+32)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v50)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v119)+40)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v123)+48)) = base.I64_extend8_s(base.I64_extend_i32_u(l3))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v127 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0))))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v131 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v130)+8)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v135 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v139 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	base.MemoryFill(m, v140, int32(0), v139)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+4)))
	v146 = v144 & int32(_a_F_recordMultipleDependencies_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+4)) = uint16(v146)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+6)) = uint16(v149)
	goto L28
L28:
	;
	v152 = v54 + int32(1)
	if v152 != v31 {
		v167 = v152
		v168 = v102
		goto L15
	} else {
		goto L29
	}
L29:
	;
	if v40 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v156 = F_CatalogOpenIndexes(m, v25)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	v158 = v40
	goto L32
L32:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v25, v32, v31, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v158 = v156
	goto L32
L34:
	;
	v164 = v59 + int32(1)
	if v164 != l2 {
		v35 = v50 + int32(12)
		v40 = v158
		v42 = v102
		v44 = v164
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v188 = v158
	v189 = v102
	goto L10
L36:
	;
	goto L14
L37:
	;
	if v40 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v180 = F_CatalogOpenIndexes(m, v25)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	v182 = v40
	goto L40
L40:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, v25, v32, v167, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	v182 = v180
	goto L40
L42:
	;
	v188 = v182
	v189 = v168
	goto L10
L43:
	;
	F_CatalogCloseIndexes(m, v188)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_relation_close(m, v25, int32(3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	if int32(0) < v189 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v202 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_pfree(m, v32)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L55
	}
L51:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v32+v202<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v222 = v202 + int32(1)
	if v222 != v189 {
		v202 = v222
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L1
}
func F_record_image_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int64
	_ = v278
	var v282 int64
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v596 int32
	_ = v596
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	v27 = m.G0
	v29 = v27 - int32(80)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v37 = F_pg_detoast_datum(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v41 = F_lookup_rowtype_tupdesc(m, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v46 = F_lookup_rowtype_tupdesc(m, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v32
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+68)) = uint16(v51)
	v55 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v55
	v57 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = int32(base.Ui32(v49) >> (uint(v57) % 32))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+48)) = uint16(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = int32(base.Ui32(v60) >> (uint(v57) % 32))
	if v48 < v43 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v72 = v43
	goto L8
L7:
	;
	v72 = v48
	goto L8
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	if v74 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v39 != v97 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v85 = F_MemoryContextAlloc(m, v80, v72<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v77 < v72 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v96 = v74
	v97 = v79
	goto L9
L13:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v85
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v90)+4)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v90)+12)) = v91
	v96 = v90
	v97 = int32(0)
	goto L9
L14:
	;
	v153 = F_palloc(m, v43<<(uint(int32(3))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v106 = v96 + int32(20)
	v110 = v72 << (uint(int32(2)) % 32)
	if v106&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v110)) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v99 != v40 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	if v101 != v44 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	if v103 == v45 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v39
	goto L14
L21:
	;
	if v110 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v110 == int32(0) {
		goto L20
	} else {
		goto L29
	}
L24:
	;
	v120 = v96 + v110 + int32(20)
	v122 = v96 + int32(24)
	if base.Ui32(v122) < base.Ui32(v120) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v120
	goto L27
L26:
	;
	v124 = v122
	goto L27
L27:
	;
	v131 = (v124-v96-int32(21))&int32(-4) + int32(4)
	if v131 == int32(0) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	base.MemoryFill(m, v106, int32(0), v131)
	goto L20
L29:
	;
	base.MemoryFill(m, v106, int32(0), v110)
	goto L20
L30:
	;
	v155 = F_palloc(m, v43)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v29+int32(60), v41, v153, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v163 = F_palloc(m, v48<<(uint(int32(3))%32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v165 = F_palloc(m, v48)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v29+int32(40), v46, v163, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v169 = int32(0)
	v173 = base.B2i32(v169 < v43)
	if v169 < v43 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L168
	}
L37:
	;
	F_pfree(m, v153)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L148
	}
L38:
	;
	if base.B2i32(v563 != v43)|base.B2i32(v562 != v48) != 0 {
		goto L36
	} else {
		goto L147
	}
L39:
	;
	v183 = v180
	v184 = v169
	v185 = base.B2i32(v169 < v48)
	v187 = v173
	v189 = v181
	goto L44
L40:
	;
	v174 = int32(0)
	v180 = v174
	v181 = v174
	goto L39
L41:
	;
	goto L42
L42:
	;
	v176 = int32(0)
	if v48 <= v176 {
		v562 = v176
		v563 = v169
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v180 = v176
	v181 = v176
	goto L39
L44:
	;
	if v187&int32(1) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v562 = v544
	v563 = v545
	goto L38
L46:
	;
	v558 = base.B2i32(v544 < v48)
	v559 = base.B2i32(v545 < v43)
	if v558|v559 != 0 {
		v183 = v544
		v184 = v545
		v185 = v558
		v187 = v559
		v189 = v549
		goto L44
	} else {
		goto L146
	}
L47:
	;
	if v185&int32(1) == int32(0) {
		v562 = v183
		v563 = v184
		goto L38
	} else {
		goto L50
	}
L48:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v212<<(uint(int32(3))%32)+v184*int32(100))+119)))
	if v219 != int32(1) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v544 = v183
	v545 = v184 + int32(1)
	v549 = v189
	goto L46
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v234 = v46 + v228<<(uint(int32(3))%32) + v183*int32(100)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+119)))
	if v235 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v544 = v183 + int32(1)
	v545 = v184
	v549 = v189
	goto L46
L52:
	;
	goto L53
L53:
	;
	if v187&int32(1) == int32(0) {
		v562 = v183
		v563 = v184
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v250 = v41 + v244<<(uint(int32(3))%32) + v184*int32(100)
	v252 = v250 + int32(28)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+96))
	v255 = v234 + int32(96)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v253 == v256 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v538 = int32(1)
	v544 = v183 + v538
	v545 = v184 + v538
	v549 = v189 + v538
	goto L46
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L143
	}
L57:
	;
	v258 = int32(1)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v165))))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v155))))
	if v262 == v258 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L137
	}
L60:
	;
	if v260&int32(1) == int32(0) {
		v596 = v258
		goto L37
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v269 = int32(-1)
	if v260&int32(1) != 0 {
		v596 = v269
		goto L37
	} else {
		goto L64
	}
L63:
	;
	goto L55
L64:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+82)))
	if v272 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v275 = int32(3)
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v153+v184<<(uint(v275)%32))))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v163+v183<<(uint(v275)%32))))
	if v278 == v282 {
		goto L55
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+72)))
	if int32(0) < v288 {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	if base.Ui64(v278) < base.Ui64(v282) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v287 = int32(-1)
	goto L71
L70:
	;
	v287 = int32(1)
	goto L71
L71:
	;
	v596 = v287
	goto L37
L72:
	;
	if v472 < int32(0) {
		v596 = v269
		goto L37
	} else {
		goto L135
	}
L73:
	;
	v291 = int32(3)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v153+v184<<(uint(v291)%32))))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v163+v183<<(uint(v291)%32))))
	if base.Ui32(int32(4)) <= base.Ui32(v288) {
		goto L79
	} else {
		goto L80
	}
L74:
	;
	goto L75
L75:
	;
	if v288 != int32(-1) {
		goto L56
	} else {
		goto L94
	}
L76:
	;
	v472 = v360
	goto L72
L77:
	;
	v360 = int32(0)
	goto L76
L78:
	;
	v334 = v329
	v335 = v330
	v336 = v331
	goto L88
L79:
	;
	if (v294|v298)&int32(3) != 0 {
		v329 = v294
		v330 = v298
		v331 = v288
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v322 = v294
	v323 = v298
	v324 = v288
	goto L81
L81:
	;
	if v324 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v306 = v294
	v307 = v298
	v308 = v288
	goto L83
L83:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v311 != v312 {
		v329 = v306
		v330 = v307
		v331 = v308
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v322 = v317
	v323 = v315
	v324 = v319
	goto L81
L85:
	;
	v314 = int32(4)
	v315 = v307 + v314
	v317 = v306 + v314
	v319 = v308 - v314
	if base.Ui32(int32(3)) < base.Ui32(v319) {
		v306 = v317
		v307 = v315
		v308 = v319
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v329 = v322
	v330 = v323
	v331 = v324
	goto L78
L88:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v339 == v340 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v360 = v339 - v340
	goto L76
L90:
	;
	v342 = int32(1)
	v347 = v336 - v342
	if v347 != 0 {
		v334 = v334 + v342
		v335 = v335 + v342
		v336 = v347
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	v365 = v153 + v184<<(uint(int32(3))%32)
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v365)))
	v367 = F_toast_raw_datum_size(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v371 = v163 + v183<<(uint(int32(3))%32)
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
	v373 = F_toast_raw_datum_size(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v376 = base.B2i32(base.Ui32(v367) < base.Ui32(v373))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v379 = F_pg_detoast_datum_packed(m, v378)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	v382 = F_pg_detoast_datum_packed(m, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v384 = int32(1)
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v386&v384 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v389 = v384
	goto L101
L100:
	;
	v389 = int32(4)
	goto L101
L101:
	;
	v390 = v379 + v389
	v391 = int32(1)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	if v393&v391 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v396 = v391
	goto L104
L103:
	;
	v396 = int32(4)
	goto L104
L104:
	;
	v397 = v382 + v396
	if base.Ui32(v367) < base.Ui32(v373) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v398 = v367
	goto L107
L106:
	;
	v398 = v373
	goto L107
L107:
	;
	v399 = int32(4)
	v400 = v398 - v399
	if base.Ui32(v399) <= base.Ui32(v400) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v463 != v379 {
		goto L126
	} else {
		goto L127
	}
L109:
	;
	v462 = int32(0)
	goto L108
L110:
	;
	v436 = v431
	v437 = v432
	v438 = v433
	goto L120
L111:
	;
	if (v390|v397)&int32(3) != 0 {
		v431 = v390
		v432 = v397
		v433 = v400
		goto L110
	} else {
		goto L114
	}
L112:
	;
	v424 = v390
	v425 = v397
	v426 = v400
	goto L113
L113:
	;
	if v426 == int32(0) {
		goto L109
	} else {
		goto L119
	}
L114:
	;
	v408 = v390
	v409 = v397
	v410 = v400
	goto L115
L115:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	if v413 != v414 {
		v431 = v408
		v432 = v409
		v433 = v410
		goto L110
	} else {
		goto L117
	}
L116:
	;
	v424 = v419
	v425 = v417
	v426 = v421
	goto L113
L117:
	;
	v416 = int32(4)
	v417 = v409 + v416
	v419 = v408 + v416
	v421 = v410 - v416
	if base.Ui32(int32(3)) < base.Ui32(v421) {
		v408 = v419
		v409 = v417
		v410 = v421
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v431 = v424
	v432 = v425
	v433 = v426
	goto L110
L120:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v441 == v442 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v462 = v441 - v442
	goto L108
L122:
	;
	v444 = int32(1)
	v449 = v438 - v444
	if v449 != 0 {
		v436 = v436 + v444
		v437 = v437 + v444
		v438 = v449
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L109
L126:
	;
	F_pfree(m, v379)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if v462 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v467 = v462
	goto L132
L131:
	;
	v467 = base.B2i32(base.Ui32(v373) < base.Ui32(v367)) - v376
	goto L132
L132:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	if v382 == v468 {
		v472 = v467
		goto L72
	} else {
		goto L133
	}
L133:
	;
	F_pfree(m, v382)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v472 = v467
	goto L72
L135:
	;
	if v472 == int32(0) {
		goto L55
	} else {
		goto L136
	}
L136:
	;
	v596 = int32(1)
	goto L37
L137:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v252)+68))
	v493 = F_format_type_be(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v496 = F_format_type_be(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v189 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v493
	F_errmsg(m, int32(_a_F_record_image_cmp_0), v29+int32(16))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1474), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v517 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v517
	F_errmsg_internal(m, int32(_a_F_record_image_cmp_3), v29)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1538), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
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
	goto L45
L147:
	;
	v596 = int32(0)
	goto L37
L148:
	;
	F_pfree(m, v155)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_pfree(m, v163)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_pfree(m, v165)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if int32(0) <= v625 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_DecrTupleDescRefCount(m, v41)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if int32(0) <= v630 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L154
L156:
	;
	F_DecrTupleDescRefCount(m, v46)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v635 != v32 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	F_pfree(m, v32)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v639 != v37 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	F_pfree(m, v37)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	m.G0 = v29 + int32(80)
	return v596
L167:
	;
	goto L166
L168:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_record_image_cmp_4), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_record_image_cmp_1), int32(1568), int32(_a_F_record_image_cmp_2))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_lt(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(int32(base.Ui32(v2) >> (uint(int32(31)) % 32)))
	}
}
func F_record_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
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
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
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
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(80)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_stack_depth(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = int32(0)
	if base.B2i32(v23 == int32(2249))&base.B2i32(v24 < v32) == v32 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v507 = F_heap_form_tuple(m, v37, v109, v112)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L96
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L92
	}
L5:
	;
	if v115 == int32(0) {
		goto L3
	} else {
		goto L91
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L77
	}
L9:
	;
	v37 = F_lookup_rowtype_tupdesc(m, v23, v24)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L73
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v41 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v109 = F_palloc_mul(m, int32(8), v39)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L14:
	;
	if v62 == v23 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v52 = F_MemoryContextAlloc(m, v47, v39*int32(44)+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v44 != v39 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v61 = v41
	v62 = v46
	goto L14
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
	v61 = v57
	v62 = int32(0)
	goto L14
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 == v24 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v69 = v39 * int32(44)
	v71 = v69 + int32(12)
	if v61&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v71)) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v23
	goto L13
L24:
	;
	if v71 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v71 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v83 = v61 + v69 + int32(12)
	v85 = v61 + int32(4)
	if base.Ui32(v85) < base.Ui32(v83) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v87 = v83
	goto L30
L29:
	;
	v87 = v85
	goto L30
L30:
	;
	v92 = (v61^int32(-1)+v87)&int32(-4) + int32(4)
	if v92 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.MemoryFill(m, v61, int32(0), v92)
	goto L23
L32:
	;
	base.MemoryFill(m, v61, int32(0), v71)
	goto L23
L33:
	;
	v112 = F_palloc_mul(m, int32(1), v39)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v115 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v39 <= int32(0) {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v119 = int32(3)
	v120 = v39 & v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v124 = v37 + v121<<(uint(v119)%32)
	v125 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v226 != v115 {
		v455 = v226
		goto L4
	} else {
		goto L48
	}
L38:
	;
	v133 = v125
	v134 = v125
	v147 = v2
	goto L41
L39:
	;
	v177 = v125
	v178 = v125
	goto L40
L40:
	;
	v195 = v177
	v196 = v178
	v207 = v2
	goto L45
L41:
	;
	v151 = v124 + v133*int32(100)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+119)))
	v153 = int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+219)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+319)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+419)))
	v167 = v134 + (v152 ^ v153) + (v156 ^ v153) + (v160 ^ v153) + (v164 ^ v153)
	v168 = int32(4)
	v169 = v133 + v168
	v171 = v147 + v168
	if v171 != v39&int32(2147483644) {
		v133 = v169
		v134 = v167
		v147 = v171
		goto L41
	} else {
		goto L43
	}
L42:
	;
	if v120 == int32(0) {
		v226 = v167
		goto L37
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v177 = v169
	v178 = v167
	goto L40
L45:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v195*int32(100))+119)))
	v215 = int32(1)
	v217 = v196 + (v214 ^ v215)
	v221 = v207 + v215
	if v221 != v120 {
		v195 = v195 + v215
		v196 = v217
		v207 = v221
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v226 = v217
	goto L37
L47:
	;
	goto L46
L48:
	;
	v247 = int32(0)
	goto L49
L49:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v269 = v37 + v263<<(uint(int32(3))%32) + v247*int32(100)
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+119)))
	if v270 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L3
L51:
	;
	v364 = v247 + int32(1)
	if v364 != v39 {
		v247 = v364
		goto L49
	} else {
		goto L72
	}
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v109+v247<<(uint(int32(3))%32)))) = int64(0)
	v279 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v112))) = uint8(v279)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v269)+96))
	v283 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v286 = int32(_a_F_record_recv_0)
	if base.B2i32(base.B2i32(v281 == v283)|base.B2i32(base.Ui32(v286) < base.Ui32(v281)) == int32(0))&base.B2i32(base.Ui32(v283) <= base.Ui32(v286)) != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v295 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v295 < int32(-1) {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v299-v300 < v295 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	v305 = v61 + int32(12) + v247*int32(44)
	if v295 == int32(-1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v247+v112))) = uint8(v321)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if v281 != v326 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v321 = int32(1)
	v323 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v295 + v300
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v300 + v312
	v321 = int32(0)
	v323 = v21 - int32(-64)
	goto L60
L64:
	;
	F_getTypeBinaryInputInfo(m, v281, v305+int32(4), v305+int32(8))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v269+int32(28))+76))
	v351 = F_ReceiveFunctionCall(m, v305+int32(16), v323, v347, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+20))
	F_fmgr_info_cxt(m, v334, v305+int32(16), v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v281
	goto L66
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v109+v247<<(uint(int32(3))%32)))) = v351
	if v323 == int32(0) {
		goto L51
	} else {
		goto L70
	}
L70:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	if v356 != v295 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L51
L72:
	;
	goto L50
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(_a_F_record_recv_1), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(509), int32(_a_F_record_recv_3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v391 = F_format_type_extended(m, v283, int32(-1), int32(2))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v395 = F_format_type_extended(m, v281, int32(-1), int32(2))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v247 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v283
	F_errmsg(m, int32(_a_F_record_recv_4), v21)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(606), int32(_a_F_record_recv_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_record_recv_5), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(613), int32(_a_F_record_recv_3))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v247 + int32(1)
	F_errmsg(m, int32(_a_F_record_recv_6), v21+int32(32))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(661), int32(_a_F_record_recv_3))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v455 = int32(0)
	goto L4
L92:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v115
	F_errmsg(m, int32(_a_F_record_recv_7), v21+int32(48))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_record_recv_2), int32(559), int32(_a_F_record_recv_3))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	v510 = F_palloc(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507)))
	if v512 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v507)+16))
	base.MemoryCopy(m, v510, v513, v512)
	goto L100
L99:
	;
	goto L100
L100:
	;
	F_pfree(m, v507)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v109)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v112)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if int32(0) <= v521 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	F_DecrTupleDescRefCount(m, v37)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v526 = F_HeapTupleHeaderGetDatum(m, v510)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	m.G0 = v21 + int32(80)
	return v526
}
func F_recurse_pushdown_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0
	goto L2
L1:
	;
	m.G0 = v8 + int32(16)
	return v58
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v15 != int32(142) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L4:
	;
	goto L3
L5:
	;
	if v15 != int32(63) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v35 == int32(3) {
		v58 = v34
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)-int32(4))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v30 = F_subquery_is_pushdown_safe(m, v29, l1, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v58 = v30
	goto L1
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v39 = F_recurse_pushdown_safe(m, v38, l1, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v39 == int32(0) {
		v58 = v34
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v10 = v43
	goto L2
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
	F_errmsg_internal(m, int32(_a_F_recurse_pushdown_safe_0), v8)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_recurse_pushdown_safe_1), int32(_a_F_recurse_pushdown_safe_2), int32(_a_F_recurse_pushdown_safe_3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_recvfrom(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	v5 = int32(0)
	v7 = m.Env.X__syscall_recvfrom(m, l0, l1, int32(_a_F_recvfrom_0), int32(64), v5, v5)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v7) {
		*(*int32)(unsafe.Add(mBase, _c_F_recvfrom[0])) = int32(0) - v7
		v15 = int32(-1)
	} else {
		v15 = v7
	}
	return v15
}
func F_redirect_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	F_appendStringInfo(m, l0, int32(_a_F_redirect_elem_desc_0), v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_regclassout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v44 int32
	_ = v44
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = base.I32_wrap_i64(v11)
	if v12 == int32(0) {
		v16 = F_pstrdup(m, int32(_a_F_regclassout_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v57 = v16
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v57)
		}
	} else {
		v23 = F_SearchSysCache1(m, int32(57), v11&int64(4294967295))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v29 = v27 + int32(4)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_regclassout[0]))
				if v31 == int32(0) {
					v34 = F_pstrdup(m, v29)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v23)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							v57 = v34
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v57)
						}
					}
				} else {
					v38 = F_RelationIsVisible(m, v12)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v38 != 0 {
							v44 = int32(0)
							v45 = F_quote_qualified_identifier(m, v44, v29)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_ReleaseCatCache(m, v23)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v57 = v45
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_u(v57)
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
							v42 = F_get_namespace_name(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								v44 = v42
								v45 = F_quote_qualified_identifier(m, v44, v29)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int64(0)
								} else {
									F_ReleaseCatCache(m, v23)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int64(0)
									} else {
										v57 = v45
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v57)
									}
								}
							}
						}
					}
				}
			} else {
				v50 = F_palloc(m, int32(64))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
					v55 = F_pg_snprintf(m, v50, int32(64), int32(_a_F_regclassout_1), v9)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v57 = v50
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v57)
					}
				}
			}
		}
	}
}
func F_regdatabasein(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
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
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int64
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L29
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v9 + int32(16)
	return v177
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_regdatabasein[0]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v177 = v6
	goto L2
L8:
	;
	v23 = int32(_a_F_regdatabasein_0)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regdatabasein[1])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L28
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_regdatabasein[2])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v96 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v96 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		v96 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v96 = v94
	goto L22
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(588), v12, int32(-1), v11, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int64(0)
L30:
	;
	v117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+8)))
	v177 = v117
	goto L2
L31:
	;
	v122 = F_stringToQualifiedNameList(m, v12, v11)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v177 = v6
	goto L2
L34:
	;
	goto L35
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v128 != int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v131 = F_errsave_start(m, v11)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v151 = F_get_database_oid(m, v149, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L29
	} else {
		goto L44
	}
L39:
	;
	if v131 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_regdatabasein_1), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L42
	}
L42:
	;
	F_errsave_finish(m, v11, int32(_a_F_regdatabasein_2), int32(1806), int32(_a_F_regdatabasein_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v177 = v6
	goto L2
L44:
	;
	if v151 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = F_errsave_start(m, v11)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v177 = base.I64_extend_i32_u(v151)
	goto L2
L48:
	;
	if v155 == int32(0) {
		v177 = v6
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v164
	F_errmsg(m, int32(_a_F_regdatabasein_4), v9)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v11, int32(_a_F_regdatabasein_2), int32(1814), int32(_a_F_regdatabasein_3))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	v177 = v6
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(_a_F_regdatabasein_5), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L29
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_regdatabasein_2), int32(1796), int32(_a_F_regdatabasein_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L29
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regex_selectivity_sub(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 float64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 float64
	_ = v138
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v145 int32
	_ = v145
	var v149 float64
	_ = v149
	var v152 float64
	_ = v152
	var v155 float64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 float64
	_ = v167
	var v174 float64
	_ = v174
	var v177 float64
	_ = v177
	v5 = int32(0)
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
	return float64(0)
L2:
	;
	v14 = float64(1)
	if l1 <= int32(0) {
		v167 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v174 = float64(1)
	if base.F64_gt(v167, v174) != 0 {
		goto L68
	} else {
		goto L69
	}
L4:
	;
	v19 = v14
	v21 = v5
	v22 = v5
	v24 = v5
	goto L5
L5:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v21))))
	if v27 == int32(40) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v167 = v155
	goto L3
L7:
	;
	v163 = v157 + int32(1)
	if v163 < l1 {
		v19 = v155
		v21 = v163
		v22 = v158
		v24 = v160
		goto L5
	} else {
		goto L67
	}
L8:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v35 = int32(0)
	if base.B2i32(v27 != int32(41))|base.B2i32(v22 <= v35) == v35 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v30 = v24
	goto L13
L12:
	;
	v30 = v21
	goto L13
L13:
	;
	v155 = v19
	v157 = v21
	v158 = v22 + int32(1)
	v160 = v30
	goto L7
L14:
	;
	v41 = v22 - int32(1)
	if v41 != 0 {
		v155 = v19
		v157 = v21
		v158 = v41
		v160 = v24
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.B2i32(v27 != int32(124))|v22 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v43 = v24 + int32(1)
	v46 = F_regex_selectivity_sub(m, l0+v43, v21-v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v155 = base.F64_mul(v19, v46)
	v157 = v21
	v158 = int32(0)
	v160 = v24
	goto L7
L19:
	;
	v56 = v21 + int32(1)
	v59 = F_regex_selectivity_sub(m, l0+v56, l1-v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v27 - int32(42) {
	case 0, 1, 21:
		goto L25
	case 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48:
		goto L23
	case 4:
		goto L26
	case 49:
		goto L27
	case 50:
		goto L24
	default:
		goto L28
	}
L22:
	;
	v167 = base.F64_add(v19, v59)
	goto L3
L23:
	;
	if v22 != 0 {
		goto L64
	} else {
		goto L65
	}
L24:
	;
	v145 = v21 + int32(1)
	if l1 <= v145 {
		v167 = v19
		goto L3
	} else {
		goto L60
	}
L25:
	;
	if v22 != 0 {
		goto L57
	} else {
		goto L58
	}
L26:
	;
	if v22 != 0 {
		goto L54
	} else {
		goto L55
	}
L27:
	;
	v97 = v21 + int32(1)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v97))))
	v101 = base.B2i32(v99 == int32(94))
	if v99 == int32(94) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	if v27 != int32(123) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if l1 <= v21 {
		v87 = v21
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v22 != 0 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v71 = v21
	goto L32
L32:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
	if v77 == int32(125) {
		v87 = v71
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v87 = l1
	goto L30
L34:
	;
	v81 = v71 + int32(1)
	if v81 != l1 {
		v71 = v81
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v93 = v19
	goto L38
L37:
	;
	v93 = base.F64_add(v19, v19)
	goto L38
L38:
	;
	v155 = v93
	v157 = v87
	v158 = v22
	v160 = v24
	goto L7
L39:
	;
	v102 = float64(0.75)
	goto L41
L40:
	;
	v102 = float64(0.25)
	goto L41
L41:
	;
	if v99 == int32(94) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v22 != 0 {
		goto L51
	} else {
		goto L52
	}
L43:
	;
	v105 = v21 + int32(2)
	goto L45
L44:
	;
	v105 = v97
	goto L45
L45:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v105))))
	v110 = v105 + base.B2i32(v107 == int32(93))
	if l1 <= v110 {
		v132 = v110
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v116 = v110
	goto L47
L47:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v116))))
	if v122 == int32(93) {
		v132 = v116
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v132 = l1
	goto L42
L49:
	;
	v126 = v116 + int32(1)
	if v126 < l1 {
		v116 = v126
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v138 = v19
	goto L53
L52:
	;
	v138 = base.F64_mul(v19, v102)
	goto L53
L53:
	;
	v155 = v138
	v157 = v132
	v158 = v22
	v160 = v24
	goto L7
L54:
	;
	v141 = v19
	goto L56
L55:
	;
	v141 = base.F64_mul(v19, float64(0.9))
	goto L56
L56:
	;
	v155 = v141
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L57:
	;
	v143 = v19
	goto L59
L58:
	;
	v143 = base.F64_add(v19, v19)
	goto L59
L59:
	;
	v155 = v143
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L60:
	;
	if v22 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v149 = v19
	goto L63
L62:
	;
	v149 = base.F64_mul(v19, float64(0.2))
	goto L63
L63:
	;
	v155 = v149
	v157 = v145
	v158 = v22
	v160 = v24
	goto L7
L64:
	;
	v152 = v19
	goto L66
L65:
	;
	v152 = base.F64_mul(v19, float64(0.2))
	goto L66
L66:
	;
	v155 = v152
	v157 = v21
	v158 = v22
	v160 = v24
	goto L7
L67:
	;
	goto L6
L68:
	;
	v177 = v174
	goto L70
L69:
	;
	v177 = v167
	goto L70
L70:
	;
	return v177
}
func F_regexnesel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14371(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_register_dirty_segment(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v206 int64
	_ = v206
	var v211 int32
	_ = v211
	v2 = l1
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v14
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+10)) = uint16(v2)
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v17
	v23 = F_RegisterSyncRequest(m, v8+int32(8), v4, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		if v23 == int32(0) {
			v29 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 != 0 {
					F_errmsg_internal(m, int32(_a_F_register_dirty_segment_0), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1533), int32(_a_F_register_dirty_segment_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[0])))
							v44 = m.G0
							v46 = v44 - int32(16)
							m.G0 = v46
							if v41 != 0 {
								F___clock_gettime(m, int32(1), v46)
								mBase = m.M
								v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
								v55 = v50 + v51*int64(1000000000)
							} else {
								v55 = int64(0)
							}
							m.G0 = v46 + int32(16)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v61 = F_FileSync(m, v59, int32(167772184))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								if int32(0) <= v61 {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v158 = int32(0)
										v164 = int32(200)
										v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
										*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
										if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
											v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
										} else {
										}
									} else {
									}
									v201 = int32(200)
									v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
									v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v211 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[9])))
									if v68 != 0 {
										v69 = int32(21)
									} else {
										v69 = int32(24)
									}
									v71 = F_errstart(m, v69, int32(0))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										if v71 == int32(0) {
											v95 = int32(1)
											v97 = int64(0)
											v101 = m.G0
											v103 = v101 - int32(16)
											m.G0 = v103
											if v55 != v97 {
												F___clock_gettime(m, int32(1), v103)
												mBase = m.M
												v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
												v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
												v114 = v109 + (v110*int64(1000000000) - v55)
												v158 = int32(0)
												v164 = int32(200)
												v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
												v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
												if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
													v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
													*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
													v185 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
													*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
												} else {
												}
											} else {
											}
											v201 = int32(200)
											v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
											v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
											F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
											mBase = m.M
											v211 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
											m.G0 = v103 + int32(16)
											m.G0 = v8 + int32(32)
											return
										} else {
											F_errcode_for_file_access(m)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
												v79 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[10]))
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
												F_errmsg(m, int32(_a_F_register_dirty_segment_4), v8)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1541), int32(_a_F_register_dirty_segment_2))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														v95 = int32(1)
														v97 = int64(0)
														v101 = m.G0
														v103 = v101 - int32(16)
														m.G0 = v103
														if v55 != v97 {
															F___clock_gettime(m, int32(1), v103)
															mBase = m.M
															v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
															v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
															v114 = v109 + (v110*int64(1000000000) - v55)
															v158 = int32(0)
															v164 = int32(200)
															v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
															*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
															v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
															if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
																v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
																*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
																v185 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
																*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
															} else {
															}
														} else {
														}
														v201 = int32(200)
														v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
														v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
														F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
														mBase = m.M
														v211 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
														m.G0 = v103 + int32(16)
														m.G0 = v8 + int32(32)
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
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[0])))
					v44 = m.G0
					v46 = v44 - int32(16)
					m.G0 = v46
					if v41 != 0 {
						F___clock_gettime(m, int32(1), v46)
						mBase = m.M
						v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v46)+8)))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
						v55 = v50 + v51*int64(1000000000)
					} else {
						v55 = int64(0)
					}
					m.G0 = v46 + int32(16)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v61 = F_FileSync(m, v59, int32(167772184))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						if int32(0) <= v61 {
							v95 = int32(1)
							v97 = int64(0)
							v101 = m.G0
							v103 = v101 - int32(16)
							m.G0 = v103
							if v55 != v97 {
								F___clock_gettime(m, int32(1), v103)
								mBase = m.M
								v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
								v114 = v109 + (v110*int64(1000000000) - v55)
								v158 = int32(0)
								v164 = int32(200)
								v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
								*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
								v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
								if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
									v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
								} else {
								}
							} else {
							}
							v201 = int32(200)
							v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
							*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
							v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
							*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
							F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
							mBase = m.M
							v211 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
							*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
							m.G0 = v103 + int32(16)
							m.G0 = v8 + int32(32)
							return
						} else {
							v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[9])))
							if v68 != 0 {
								v69 = int32(21)
							} else {
								v69 = int32(24)
							}
							v71 = F_errstart(m, v69, int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								if v71 == int32(0) {
									v95 = int32(1)
									v97 = int64(0)
									v101 = m.G0
									v103 = v101 - int32(16)
									m.G0 = v103
									if v55 != v97 {
										F___clock_gettime(m, int32(1), v103)
										mBase = m.M
										v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
										v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
										v114 = v109 + (v110*int64(1000000000) - v55)
										v158 = int32(0)
										v164 = int32(200)
										v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
										*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
										v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
										if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
											v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
											*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
											v185 = int32(1)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
											*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
										} else {
										}
									} else {
									}
									v201 = int32(200)
									v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
									v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
									*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
									F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
									mBase = m.M
									v211 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
									*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
									m.G0 = v103 + int32(16)
									m.G0 = v8 + int32(32)
									return
								} else {
									F_errcode_for_file_access(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v79 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[10]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v77*int32(48))+32))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
										F_errmsg(m, int32(_a_F_register_dirty_segment_4), v8)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_register_dirty_segment_1), int32(1541), int32(_a_F_register_dirty_segment_2))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v95 = int32(1)
												v97 = int64(0)
												v101 = m.G0
												v103 = v101 - int32(16)
												m.G0 = v103
												if v55 != v97 {
													F___clock_gettime(m, int32(1), v103)
													mBase = m.M
													v109 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
													v110 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
													v114 = v109 + (v110*int64(1000000000) - v55)
													v158 = int32(0)
													v164 = int32(200)
													v165 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1]))
													*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[1])) = v165 + v114
													v169 = *(*int32)(unsafe.Add(mBase, _c_F_register_dirty_segment[2]))
													if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v169))|base.B2i32(int32(1)<<(uint(v169)%32)&int32(_a_F_register_dirty_segment_3) == v158) == v158 {
														v181 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3]))
														*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[3])) = v181 + v114
														v185 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v185)
														*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[5])) = uint8(v185)
													} else {
													}
												} else {
												}
												v201 = int32(200)
												v202 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[6])) = v202 + base.I64_extend_i32_u(v95)
												v206 = *(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7]))
												*(*int64)(unsafe.Add(mBase, _c_F_register_dirty_segment[7])) = v206 + v97
												F_pgstat_count_backend_io_op(m, int32(0), int32(3), v95, v95, v97)
												mBase = m.M
												v211 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[4])) = uint8(v211)
												*(*uint8)(unsafe.Add(mBase, _c_F_register_dirty_segment[8])) = uint8(v211)
												m.G0 = v103 + int32(16)
												m.G0 = v8 + int32(32)
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
		} else {
			m.G0 = v8 + int32(32)
			return
		}
	}
}
func F_regoperout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
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
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = base.I32_wrap_i64(v12)
	if v13 == int32(0) {
		v17 = F_pstrdup(m, int32(_a_F_regoperout_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v92 = v17
			m.G0 = v10 + int32(48)
			return base.I64_extend_i32_u(v92)
		}
	} else {
		v24 = F_SearchSysCache1(m, int32(40), v12&int64(4294967295))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			if v24 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
				v28 = v26 + v27
				v30 = v28 + int32(4)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_regoperout[0]))
				if v32 == int32(0) {
					v35 = F_pstrdup(m, v30)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v24)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v92 = v35
							m.G0 = v10 + int32(48)
							return base.I64_extend_i32_u(v92)
						}
					}
				} else {
					v39 = F_makeString(m, v30)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v39
						v46 = F_list_make1_impl(m, int32(1), v10+int32(36))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							v48 = int32(0)
							v52 = F_OpernameGetCandidates(m, v46, v48, v48, v10+int32(44))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								if v52 == int32(0) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
									v64 = F_get_namespace_name(m, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										v66 = F_quote_identifier(m, v64)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int64(0)
										} else {
											v68 = F_strlen(m, v66)
											mBase = m.M
											v69 = F_strlen(m, v30)
											mBase = m.M
											v73 = F_palloc(m, v68+v69+int32(2))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v66
												v80 = F_pg_sprintf(m, v73, int32(_a_F_regoperout_1), v10+int32(16))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int64(0)
													} else {
														v92 = v73
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v92)
													}
												}
											}
										}
									}
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									if v56 != 0 {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
										v64 = F_get_namespace_name(m, v63)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int64(0)
										} else {
											v66 = F_quote_identifier(m, v64)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int64(0)
											} else {
												v68 = F_strlen(m, v66)
												mBase = m.M
												v69 = F_strlen(m, v30)
												mBase = m.M
												v73 = F_palloc(m, v68+v69+int32(2))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v66
													v80 = F_pg_sprintf(m, v73, int32(_a_F_regoperout_1), v10+int32(16))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int64(0)
													} else {
														F_ReleaseCatCache(m, v24)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int64(0)
														} else {
															v92 = v73
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v92)
														}
													}
												}
											}
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
										if v57 != v13 {
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
											v64 = F_get_namespace_name(m, v63)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int64(0)
											} else {
												v66 = F_quote_identifier(m, v64)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int64(0)
												} else {
													v68 = F_strlen(m, v66)
													mBase = m.M
													v69 = F_strlen(m, v30)
													mBase = m.M
													v73 = F_palloc(m, v68+v69+int32(2))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v66
														v80 = F_pg_sprintf(m, v73, int32(_a_F_regoperout_1), v10+int32(16))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int64(0)
														} else {
															F_ReleaseCatCache(m, v24)
															mBase = m.M
															v83 = m.ExcPending
															if v83 != 0 {
																return int64(0)
															} else {
																v92 = v73
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v92)
															}
														}
													}
												}
											}
										} else {
											v59 = F_pstrdup(m, v30)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int64(0)
											} else {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int64(0)
												} else {
													v92 = v59
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v92)
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
				v85 = F_palloc(m, int32(64))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
					v90 = F_pg_snprintf(m, v85, int32(64), int32(_a_F_regoperout_2), v10)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						v92 = v85
						m.G0 = v10 + int32(48)
						return base.I64_extend_i32_u(v92)
					}
				}
			}
		}
	}
}
func F_regprocout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = base.I32_wrap_i64(v12)
	if v13 == int32(0) {
		v17 = F_pstrdup(m, int32(_a_F_regprocout_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v80 = v17
			m.G0 = v10 + int32(16)
			return base.I64_extend_i32_u(v80)
		}
	} else {
		v24 = F_SearchSysCache1(m, int32(47), v12&int64(4294967295))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			if v24 != 0 {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
				v28 = v26 + v27
				v30 = v28 + int32(4)
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_regprocout[0]))
				if v32 == int32(0) {
					v35 = F_pstrdup(m, v30)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v24)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v80 = v35
							m.G0 = v10 + int32(16)
							return base.I64_extend_i32_u(v80)
						}
					}
				} else {
					v39 = F_makeString(m, v30)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v39
						v46 = F_list_make1_impl(m, int32(1), v10+int32(4))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int64(0)
						} else {
							v49 = int32(0)
							v56 = F_FuncnameGetCandidates(m, v46, int32(-1), v49, v49, v49, v49, v49, v10+int32(12))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								if v56 == int32(0) {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
									v65 = F_get_namespace_name(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
									} else {
										v67 = v65
										v68 = F_quote_qualified_identifier(m, v67, v30)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int64(0)
										} else {
											F_ReleaseCatCache(m, v24)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int64(0)
											} else {
												v80 = v68
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v80)
											}
										}
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									if v60 != 0 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
										v65 = F_get_namespace_name(m, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int64(0)
										} else {
											v67 = v65
											v68 = F_quote_qualified_identifier(m, v67, v30)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int64(0)
											} else {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int64(0)
												} else {
													v80 = v68
													m.G0 = v10 + int32(16)
													return base.I64_extend_i32_u(v80)
												}
											}
										}
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
										if v62 == v13 {
											v67 = int32(0)
											v68 = F_quote_qualified_identifier(m, v67, v30)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int64(0)
											} else {
												F_ReleaseCatCache(m, v24)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int64(0)
												} else {
													v80 = v68
													m.G0 = v10 + int32(16)
													return base.I64_extend_i32_u(v80)
												}
											}
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
											v65 = F_get_namespace_name(m, v64)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int64(0)
											} else {
												v67 = v65
												v68 = F_quote_qualified_identifier(m, v67, v30)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int64(0)
												} else {
													F_ReleaseCatCache(m, v24)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int64(0)
													} else {
														v80 = v68
														m.G0 = v10 + int32(16)
														return base.I64_extend_i32_u(v80)
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
				v73 = F_palloc(m, int32(64))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
					v78 = F_pg_snprintf(m, v73, int32(64), int32(_a_F_regprocout_1), v10)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						v80 = v73
						m.G0 = v10 + int32(16)
						return base.I64_extend_i32_u(v80)
					}
				}
			}
		}
	}
}
func F_regprocrecv(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_oidrecv(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_regtypeout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = base.I32_wrap_i64(v10)
	if v11 == int32(0) {
		v15 = F_pstrdup(m, int32(_a_F_regtypeout_0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v49 = v15
			m.G0 = v8 + int32(16)
			return base.I64_extend_i32_u(v49)
		}
	} else {
		v22 = F_SearchSysCache1(m, int32(82), v10&int64(4294967295))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			if v22 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_regtypeout[0]))
				if v25 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
					v33 = F_pstrdup(m, v28+v29+int32(4))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v22)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v49 = v33
							m.G0 = v8 + int32(16)
							return base.I64_extend_i32_u(v49)
						}
					}
				} else {
					v37 = F_format_type_be(m, v11)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						F_ReleaseCatCache(m, v22)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v49 = v37
							m.G0 = v8 + int32(16)
							return base.I64_extend_i32_u(v49)
						}
					}
				}
			} else {
				v42 = F_palloc(m, int32(64))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
					v47 = F_pg_snprintf(m, v42, int32(64), int32(_a_F_regtypeout_1), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = v42
						m.G0 = v8 + int32(16)
						return base.I64_extend_i32_u(v49)
					}
				}
			}
		}
	}
}
func F_relmap_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	if base.Ui32(l0) < base.Ui32(int32(16)) {
		v6 = int32(_a_F_relmap_identify_0)
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_relmap_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(544)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v12 = v10 & int32(240)
	if v12 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
		if v16 != int32(524) {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v73
				F_errmsg_internal(m, int32(_a_F_relmap_redo_0), v7)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_relmap_redo_1), int32(1112), int32(_a_F_relmap_redo_2))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v20 = v7 + int32(20)
			base.MemoryCopy(m, v20, v15+int32(12), int32(524))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v27 = F_GetDatabasePath(m, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_relmap_redo[0]))
				v34 = F_LWLockAcquire(m, v30+int32(3200), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(0)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					F_write_relmap_file(m, v20, v36, int32(1), v36, v39, v40, v27)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_relmap_redo[0]))
						F_LWLockRelease(m, v44+int32(3200))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								m.G0 = v7 + int32(544)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v12
			F_errmsg_internal(m, int32(_a_F_relmap_redo_3), v7+int32(16))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_relmap_redo_1), int32(1141), int32(_a_F_relmap_redo_2))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
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
func F_remove_dbtablespaces(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v17 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+188))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	m.T0[v156].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L44
	}
L2:
	;
	return
L3:
	;
	v19 = int32(0)
	v21 = F_table_beginscan_catalog(m, v17, v19, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = F_heap_getnext(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v23 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v23
	v30 = int32(0)
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)))
	if v40 != int32(1664) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v80 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v43 = F_GetDatabasePath(m, l0, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L13
	}
L10:
	;
	v80 = v30
	goto L11
L11:
	;
	v81 = F_heap_getnext(m, v21)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L26
	}
L12:
	;
	F_pfree(m, v43)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L25
	}
L13:
	;
	v49 = F___fstatat(m, int32(-100), v43, v13+int32(16), int32(256))
	mBase = m.M
	goto L14
L14:
	;
	if v49 < int32(0) {
		v76 = v30
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v52&int32(_a_F_remove_dbtablespaces_0) != int32(_a_F_remove_dbtablespaces_1) {
		v76 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v57 = F_rmtree(m, v43)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v74 = F_lappend_oid(m, v30, v40)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L24
	}
L18:
	;
	if v57 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v61 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	if v61 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v43
	F_errmsg(m, int32(_a_F_remove_dbtablespaces_2), v13)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_remove_dbtablespaces_3), int32(3058), int32(_a_F_remove_dbtablespaces_4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v76 = v74
	goto L12
L25:
	;
	v80 = v76
	goto L11
L26:
	;
	if v81 != 0 {
		v28 = v81
		v30 = v80
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v87 = v85 << (uint(int32(2)) % 32)
	v88 = F_palloc(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	goto L1
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if int32(0) < v90 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	F_XLogBeginInsert(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v105 = v95 << (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107+v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v88+v105))) = v109
	v112 = v95 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v112 < v113 {
		v95 = v112
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L34
L37:
	;
	goto L36
L38:
	;
	F_XLogRegisterData(m, v13+int32(16), int32(8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	F_XLogRegisterData(m, v88, v87)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v138 = F_XLogInsert(m, int32(4), int32(33))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_list_free(m, v80)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_pfree(m, v88)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L1
L44:
	;
	F_relation_close(m, v17, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	m.G0 = v13 + int32(112)
	return
}
func F_remove_nulling_relids(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14227(m, l0, l1, l2, int32(1131))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_rename(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = int32(-100)
	v5 = m.Env.X__syscall_renameat(m, v3, l0, v3, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, _c_F_rename[0])) = int32(0) - v5
		v13 = int32(-1)
	} else {
		v13 = v5
	}
	return v13
}
func F_rename_constraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L60
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L56
	}
L3:
	;
	v37 = F_SearchSysCache1(m, int32(19), base.I64_extend_i32_u(v35))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L12
	}
L4:
	;
	v22 = F_get_domain_constraint_oid(m, l2, l3, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = F_relation_open(m, l1, int32(8))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v34 = int32(0)
	v35 = v22
	goto L3
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	F_renameatt_check(m, l1, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v32 = F_get_relation_constraint_oid(m, l1, l3, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = v25
	v35 = v32
	goto L3
L12:
	;
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
	v41 = v39 + v40
	if l1 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L53
	}
L16:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v41)+88))
	if v140 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+72)))
	switch v44 - int32(99) {
	case 0, 11:
		goto L18
	default:
		goto L16
	}
L18:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+106)))
	if v47 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if l5 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+104)))
	if l6 < v123 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v51 = F_find_all_inheritors(m, l1, int32(8), v16+int32(-4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if l6 != 0 {
		goto L20
	} else {
		goto L37
	}
L24:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v57 = int32(0)
	goto L25
L25:
	;
	v70 = int32(0)
	if v51 == v70 {
		v80 = v70
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v53 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v74 <= v57 {
		v80 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v80 = v76 + v57<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if base.B2i32(v80 == int32(0))|base.B2i32(v85 <= v57) != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v88 == int32(0) {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if l1 != v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v95 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88+v57<<(uint(int32(2))%32))))
	F_rename_constraint_internal(m, v16+int32(-16), v91, v95, l3, l4, v95, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v57 = v57 + int32(1)
	goto L25
L36:
	;
	goto L35
L37:
	;
	v106 = F_find_inheritance_children(m, l1, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	if v106 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	goto L20
L40:
	;
	goto L16
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	F_ReleaseCatCache(m, v37)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L47
	}
L42:
	;
	F_RenameConstraintById(m, v35, l4)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L46
	}
L43:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+72)))
	v145 = v143 - int32(112)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v145))|base.B2i32(int32(1)<<(uint(v145)%32)&int32(289) == int32(0)) != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_RenameRelationInternal(m, v140, l4, int32(0), int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L41
L47:
	;
	if v34 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_CacheInvalidateRelcache(m, v34)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	m.G0 = v18 - int32(-64)
	return
L51:
	;
	F_relation_close(m, v34, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v35
	F_errmsg_internal(m, int32(_a_F_rename_constraint_internal_0), v18)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(_a_F_rename_constraint_internal_2), int32(_a_F_rename_constraint_internal_3))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l3
	F_errmsg(m, int32(_a_F_rename_constraint_internal_4), v16+int32(-32))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(_a_F_rename_constraint_internal_5), int32(_a_F_rename_constraint_internal_3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l3
	F_errmsg(m, int32(_a_F_rename_constraint_internal_6), v16+int32(-48))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_rename_constraint_internal_1), int32(_a_F_rename_constraint_internal_7), int32(_a_F_rename_constraint_internal_3))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_reorderqueue_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v7 = F_pairingheap_remove_first(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = int32(0)
	v18 = v12
	goto L6
L4:
	;
	goto L5
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v17))))
	if v22 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v36 = v17 + int32(1)
	if v36 < v34 {
		v17 = v36
		v18 = v34
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))))
	if v25 != 0 {
		v34 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v17<<(uint(int32(3))%32))))
	F_pfree(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v34 = v33
	goto L8
L12:
	;
	goto L7
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	F_pfree(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_pfree(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return v11
}
func F_repalloc0(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.Ui32(l1) <= base.Ui32(l2) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13&int32(15)*int32(36))+uint32(_c_F_repalloc0[0])))
		v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, l0, l2, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = l2 - l1
			if v23 != 0 {
				base.MemoryFill(m, v19+l1, int32(0), v23)
			} else {
			}
			m.G0 = v7 + int32(16)
			return v19
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
			F_errmsg_internal(m, int32(_a_F_repalloc0_0), v7)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_repalloc0_1), int32(1721), int32(_a_F_repalloc0_2))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
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
func F_repalloc_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_c_F_repalloc_extended[0])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(2))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_repalloc_huge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6&int32(15)*int32(36))+uint32(_c_F_repalloc_huge[0])))
	v12 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_reparameterize_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v162 int32
	_ = v162
	var v163 float64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int64
	_ = v290
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
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 float64
	_ = v343
	var v347 float64
	_ = v347
	var v355 float64
	_ = v355
	var v361 float64
	_ = v361
	var v367 float64
	_ = v367
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 float64
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v417 float64
	_ = v417
	var v426 float64
	_ = v426
	var v430 float64
	_ = v430
	var v431 int64
	_ = v431
	var v436 int32
	_ = v436
	var v438 float64
	_ = v438
	var v440 float64
	_ = v440
	var v443 float64
	_ = v443
	var v446 float64
	_ = v446
	var v452 int32
	_ = v452
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = v18
	goto L3
L2:
	;
	v19 = v5
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = int32(0)
	if v19 == v21 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v15 + int32(32)
	return v452
L5:
	;
	if v74 == int32(0) {
		v452 = v5
		goto L4
	} else {
		goto L19
	}
L6:
	;
	v74 = int32(1)
	goto L5
L7:
	;
	goto L8
L8:
	;
	if l2 == int32(0) {
		v67 = v21
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v74 = v67
	goto L5
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v31 < v30 {
		v67 = v21
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(1)
	if v30 <= v33 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = v33
	goto L14
L13:
	;
	v36 = v30
	goto L14
L14:
	;
	v37 = int32(8)
	v42 = int32(0)
	goto L15
L15:
	;
	v49 = v42 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v19+v37+v49)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2+v37+v49)))
	v56 = v51 & (v53 ^ int32(-1))
	v58 = base.B2i32(v56 == int32(0))
	if v56 != 0 {
		v67 = v58
		goto L9
	} else {
		goto L17
	}
L16:
	;
	v67 = v58
	goto L9
L17:
	;
	v60 = v42 + int32(1)
	if v60 != v36 {
		v42 = v60
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v77 - int32(335) {
	case 0:
		goto L23
	default:
		v452 = v5
		goto L4
	case 3:
		goto L22
	case 8:
		goto L28
	case 9:
		goto L27
	case 10, 11:
		goto L26
	case 13:
		goto L25
	case 16:
		goto L24
	case 29:
		goto L21
	case 30:
		goto L20
	}
L20:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v376 = F_reparameterize_path(m, l0, v375, l2, l3)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L29
	} else {
		goto L82
	}
L21:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v298 = F_reparameterize_path(m, l0, v297, l2, l3)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L29
	} else {
		goto L72
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v221 != 0 {
		goto L52
	} else {
		goto L53
	}
L23:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v193 != int32(282) {
		v452 = v5
		goto L4
	} else {
		goto L48
	}
L24:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v162)+56))
	v165 = F_palloc0(m, int32(80))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L29
	} else {
		goto L42
	}
L25:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v139 = F_palloc0(m, int32(80))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L29
	} else {
		goto L39
	}
L26:
	;
	v125 = F_palloc0(m, int32(112))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L29
	} else {
		goto L36
	}
L27:
	;
	v104 = F_palloc0(m, int32(72))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L29
	} else {
		goto L33
	}
L28:
	;
	v81 = F_palloc0(m, int32(72))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = int64(1473173782810)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v88
	v90 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+20)) = uint8(v92)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v90
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+64)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = v92
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+21)) = uint8(v95)
	F_cost_seqscan(m, v81, l0, v20, v90)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v452 = v81
	goto L4
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = int64(1477468750106)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v109
	v111 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+20)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+16)) = v111
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+64)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v104)+24)) = v113
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+21)) = uint8(v116)
	F_cost_samplescan(m, v104, l0, v20, v111)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	v452 = v104
	goto L4
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = int32(283)
	base.MemoryCopy(m, v125, l1, int32(112))
	v131 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v131
	F_cost_index(m, v125, l0, l3, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	v452 = v125
	goto L4
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = int64(1494648619293)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = v144
	v146 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+20)) = uint8(v148)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v146
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+72)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v139)+64)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v148
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+21)) = uint8(v151)
	F_cost_bitmap_heap_scan(m, v139, l0, v20, v146, v137, l3)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L29
	} else {
		goto L41
	}
L41:
	;
	v452 = v139
	goto L4
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(1507533521186)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v170
	v172 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+20)) = uint8(v174)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v172
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v177 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+21)))
	v182 = v180
	goto L46
L45:
	;
	v182 = int32(0)
	goto L46
L46:
	;
	v184 = v182 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+21)) = uint8(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v162)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+72)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v165)+64)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = v186
	F_cost_subqueryscan(m, v165, l0, v20, v172, base.F64_eq(v161, v163))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L29
	} else {
		goto L47
	}
L47:
	;
	v452 = v165
	goto L4
L48:
	;
	v197 = F_palloc0(m, int32(72))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = int64(1438814044442)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+12)) = v202
	v204 = F_get_baserel_parampathinfo(m, l0, v20, l2)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v206)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v204
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+64)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v197)+24)) = v206
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+21)) = uint8(v209)
	F_cost_resultscan(m, v197, l0, v20, v204)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L29
	} else {
		goto L51
	}
L51:
	;
	v452 = v197
	goto L4
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if int32(0) < v222 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v281 = v5
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v281
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v290
	v295 = F_create_append_path(m, l0, v20, v15+int32(8), v287, l2, v286, v285, float64(-1))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L29
	} else {
		goto L71
	}
L55:
	;
	v229 = v5
	v234 = v5
	v235 = v5
	goto L58
L56:
	;
	v268 = v5
	v269 = v5
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v269
	v281 = v268
	goto L54
L58:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v229<<(uint(int32(2))%32))))
	v242 = F_reparameterize_path(m, l0, v241, l2, l3)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L29
	} else {
		goto L60
	}
L59:
	;
	v268 = v253
	v269 = v254
	goto L57
L60:
	;
	if v242 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v452 = int32(0)
	goto L4
L62:
	;
	goto L63
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v229 < v247 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v256 = v229 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v256 < v257 {
		v229 = v256
		v234 = v253
		v235 = v254
		goto L58
	} else {
		goto L70
	}
L65:
	;
	v249 = F_lappend(m, v234, v242)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L29
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v251 = F_lappend(m, v235, v242)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L29
	} else {
		goto L69
	}
L68:
	;
	v253 = v249
	v254 = v235
	goto L64
L69:
	;
	v253 = v234
	v254 = v251
	goto L64
L70:
	;
	goto L59
L71:
	;
	v452 = v295
	goto L4
L72:
	;
	if v298 == int32(0) {
		v452 = v5
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+40))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v305 = F_palloc0(m, int32(80))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L29
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v305))) = int64(1563368096040)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+12)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v298)+16))
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305)+20)) = uint8(v313)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+16)) = v312
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v316 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+21)))
	v321 = v319
	goto L77
L76:
	;
	v321 = int32(0)
	goto L77
L77:
	;
	v323 = v321 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v305)+21)) = uint8(v323)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v298)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+24)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v298)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v305)+72)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v305)+64)) = v327
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v298)+40))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v298)+48))
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v298)+56))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v298)+32))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+32))
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_reparameterize_path[0]))
	*(*float64)(unsafe.Add(mBase, uint32(v305)+32)) = v334
	v343 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[1]))
	v347 = base.F64_add(base.F64_mul(base.F64_add(v343, v343), v334), base.F64_sub(v333, v332))
	v355 = base.F64_mul(v334, base.F64_convert_i32_u((v336+int32(7))&int32(-8)+int32(24)))
	if base.F64_gt(v355, base.F64_convert_i32_u(v340<<(uint(int32(10))%32))) != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v452 = v305
	goto L4
L79:
	;
	v361 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[2]))
	v367 = base.F64_add(base.F64_mul(v361, base.F64_ceil(base.F64_mul(v355, float64(0.0001220703125)))), v347)
	goto L81
L80:
	;
	v367 = v347
	goto L81
L81:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v305)+48)) = v332
	*(*float64)(unsafe.Add(mBase, uint32(v305)+56)) = base.F64_add(v332, v367)
	*(*int32)(unsafe.Add(mBase, uint32(v305)+40)) = v331 + (base.B2i32(v303 <= v302) ^ int32(1))
	goto L78
L82:
	;
	if v376 == int32(0) {
		v452 = v5
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v380 = *(*float64)(unsafe.Add(mBase, uint32(l1)+96))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+85)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+84)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v386 = F_palloc0(m, int32(120))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L29
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v386))) = int64(1567663063337)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v376)+16))
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+20)) = uint8(v394)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+16)) = v393
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+26)))
	if v397 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+21)))
	v402 = v400
	goto L87
L86:
	;
	v402 = int32(0)
	goto L87
L87:
	;
	v404 = v402 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+21)) = uint8(v404)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v376)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v376)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+88)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+85)) = uint8(v381)
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+84)) = uint8(v382)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+80)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v386)+76)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v386)+72)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v386)+64)) = v408
	v417 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v380)&int64(9223372036854775807)))|base.F64_gt(v380, v417) != 0 {
		v430 = v417
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v431 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v386)+104)) = v431
	*(*float64)(unsafe.Add(mBase, uint32(v386)+96)) = v430
	*(*int64)(unsafe.Add(mBase, uint32(v386)+112)) = v431
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v376)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+40)) = v436
	v438 = *(*float64)(unsafe.Add(mBase, uint32(v376)+48))
	v440 = *(*float64)(unsafe.Add(mBase, _c_F_reparameterize_path[3]))
	*(*float64)(unsafe.Add(mBase, uint32(v386)+48)) = base.F64_add(v438, v440)
	v443 = *(*float64)(unsafe.Add(mBase, uint32(v376)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v386)+56)) = base.F64_add(v440, v443)
	v446 = *(*float64)(unsafe.Add(mBase, uint32(v376)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v386)+32)) = v446
	v452 = v386
	goto L4
L89:
	;
	goto L88
L90:
	;
	v426 = float64(1)
	if base.F64_le(v380, v426) != 0 {
		v430 = v426
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v430 = base.F64_nearest(v380)
	goto L89
}
func F_replace_vars_in_jointree(m *base.Module, l0 int32, l1 int32) {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(63) {
	case 0:
		goto L7
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L4
	}
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v134 = F_replace_rte_variables(m, v130, v16, int32(0), int32(899), l1, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L36
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L33
	}
L5:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_replace_vars_in_jointree(m, v96, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L27
	}
L6:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v60 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v15 == v16 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+v15<<(uint(int32(2))%32)-int32(4))))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+124)))
	if v28 != int32(1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	switch v31 {
	case 0:
		goto L3
	case 1:
		goto L13
	default:
		goto L1
	case 3:
		goto L12
	case 4:
		goto L11
	case 5:
		goto L10
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v57 = F_replace_rte_variables(m, v53, v16, int32(0), int32(899), l1, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L18
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v50 = F_replace_rte_variables(m, v46, v16, int32(0), int32(899), l1, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L14
	} else {
		goto L17
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v43 = F_replace_rte_variables(m, v39, v16, int32(0), int32(899), l1, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L16
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	v36 = F_replace_rte_variables(m, v32, v16, int32(1), int32(899), l1, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v36
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v43
	goto L1
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v50
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v57
	goto L1
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v92 = F_replace_rte_variables(m, v87, v88, int32(0), int32(899), l1, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L26
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = v3
	goto L22
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v70<<(uint(int32(2))%32))))
	F_replace_vars_in_jointree(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v79 = v70 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v79 < v80 {
		v70 = v79
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92
	goto L1
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_replace_vars_in_jointree(m, v99, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v102 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(2)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v112 = F_replace_rte_variables(m, v107, v108, int32(0), int32(899), l1, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v95
	goto L1
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v120
	F_errmsg_internal(m, int32(_a_F_replace_vars_in_jointree_0), v8)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_replace_vars_in_jointree_1), int32(2762), int32(_a_F_replace_vars_in_jointree_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v134
	goto L1
}
func F_report_corruption(m *base.Module, l0 int32, l1 int32) {
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+124)))
	v16 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+84)))
	v17 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = base.I64_extend16_s(base.I64_extend_i32_u(v15))
	v26 = int32(base.Ui32(v15) >> (uint(int32(15)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v26)
	v28 = F_cstring_to_text(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = base.I64_extend_i32_u(v28)
		F_pfree(m, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v38 = F_heap_form_tuple(m, v14, v11+int32(16), v11+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_tuplestore_puttuple(m, v13, v38)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					v42 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v42)
					m.G0 = v11 + int32(48)
					return
				}
			}
		}
	}
}
func F_report_invalid_encoding_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l2 < l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = l2
	goto L3
L2:
	;
	v13 = l3
	goto L3
L3:
	;
	if int32(0) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = int32(8)
	if v16 <= v13 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L19
	}
L7:
	;
	v19 = v16
	goto L9
L8:
	;
	v19 = v13
	goto L9
L9:
	;
	v27 = int32(0)
	v28 = v10 + int32(32)
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v27))))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v33
	v38 = F_pg_sprintf(m, v28, int32(_a_F_report_invalid_encoding_int_0), v10+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L6
L12:
	;
	return
L13:
	;
	v40 = v38 + v28
	if v27 < v19-int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = F_pg_sprintf(m, v40, int32(_a_F_report_invalid_encoding_int_1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	v47 = v40
	goto L16
L16:
	;
	v49 = v27 + int32(1)
	if v49 != v19 {
		v27 = v49
		v28 = v47
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v47 = v44 + v40
	goto L16
L18:
	;
	goto L11
L19:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_report_invalid_encoding_int[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v10 + int32(32)
	F_errmsg(m, int32(_a_F_report_invalid_encoding_int_2), v10)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_report_invalid_encoding_int_3), int32(1854), int32(_a_F_report_invalid_encoding_int_4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_report_triggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int64
	_ = v133
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v16 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(80)
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 <= int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = v16
	v34 = v4
	goto L5
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v39 = v36 + v34*int32(368)
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+360))
	if v40 != int64(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L1
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v44 = int32(0)
	F_ExplainOpenGroup(m, int32(_a_F_report_triggers_0), v44, int32(1), l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v155 = v28
	goto L9
L9:
	;
	v161 = v34 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v161 < v162 {
		v28 = v155
		v34 = v161
		goto L5
	} else {
		goto L54
	}
L10:
	;
	return
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v54 = v43 + v34*int32(60)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	if v55 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v56 = F_get_constraint_name(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v58 = v44
	goto L14
L14:
	;
	v60 = v51 + int32(4)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v61 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v58 = v56
	goto L14
L16:
	;
	if v58 != 0 {
		goto L49
	} else {
		goto L50
	}
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v66 = int32(0)
	if v65|base.B2i32(v58 == v66) == v66 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L19
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_1), v119, l2)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L38
	}
L20:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v58
	F_appendStringInfo(m, v83, int32(_a_F_report_triggers_2), v14+int32(48))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L28
	}
L22:
	;
	F_appendStringInfoString(m, v64, int32(_a_F_report_triggers_0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v74
	F_appendStringInfo(m, v64, int32(_a_F_report_triggers_3), v14-int32(-64))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	goto L21
L26:
	;
	if v58 == int32(0) {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	goto L20
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v60
	F_appendStringInfo(m, v90, int32(_a_F_report_triggers_4), v14+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v98 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v39)+184))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v39)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v102
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = base.F64_div(base.F64_convert_i64_s(v101), float64(1e+06))
	F_appendStringInfo(m, v97, int32(_a_F_report_triggers_5), v14)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v39)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v111
	F_appendStringInfo(m, v97, int32(_a_F_report_triggers_6), v14+int32(16))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	goto L16
L37:
	;
	goto L16
L38:
	;
	if v58 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_7), v58, l2)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_ExplainPropertyText(m, int32(_a_F_report_triggers_8), v60, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v128 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v39)+184))
	F_ExplainPropertyFloat(m, int32(_a_F_report_triggers_9), int32(_a_F_report_triggers_10), base.F64_div(base.F64_convert_i64_s(v133), float64(1e+06)), int32(3), l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v39)+360))
	F_ExplainPropertyInteger(m, int32(_a_F_report_triggers_11), int32(0), v142, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L16
L49:
	;
	F_pfree(m, v58)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_ExplainCloseGroup(m, int32(_a_F_report_triggers_0), int32(1), l2)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v155 = v154
	goto L9
L54:
	;
	goto L6
}
func F_reserveAllocatedDesc(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0]))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[1]))
	if v8 < v6 {
		v56 = int32(1)
		return v56
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2]))
		if v11 == int32(0) {
			v16 = F_emscripten_builtin_malloc(m, int32(192))
			mBase = m.M
			if v16 != 0 {
				v46 = int32(16)
				v47 = v16
				*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0])) = v46
				*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2])) = v47
				v56 = int32(1)
				return v56
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_reserveAllocatedDesc_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_reserveAllocatedDesc_1), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_reserveAllocatedDesc_2), int32(2581), int32(_a_F_reserveAllocatedDesc_3))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
			v35 = int32(0)
			v37 = *(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[3]))
			v39 = base.I32_div_s(v37, int32(3))
			if v39 <= v6 {
				v56 = v35
			} else {
				v43 = F_emscripten_builtin_realloc(m, v11, v39*int32(12))
				mBase = m.M
				if v43 == int32(0) {
					v56 = v35
				} else {
					v46 = v39
					v47 = v43
					*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[0])) = v46
					*(*int32)(unsafe.Add(mBase, _c_F_reserveAllocatedDesc[2])) = v47
					v56 = int32(1)
				}
			}
			return v56
		}
	}
}
func F_resolve_anymultirange_from_others(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = F_getBaseType(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = F_get_range_multirange(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				if v11 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v40 = F_format_type_be(m, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v40
								F_errmsg(m, int32(_a_F_resolve_anymultirange_from_others_0), v6)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_resolve_anymultirange_from_others_1), int32(728), int32(_a_F_resolve_anymultirange_from_others_2))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
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
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v11
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_resolve_anymultirange_from_others_3), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_resolve_anymultirange_from_others_1), int32(732), int32(_a_F_resolve_anymultirange_from_others_2))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
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
func F_restore(m *base.Module, l0 int32, l1 float32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
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
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
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
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v776 int32
	_ = v776
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = base.F64_promote_f32(l1)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_restore[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+56)) = uint8(v18)
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_restore[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v21
	v24 = *(*int64)(unsafe.Add(mBase, _c_F_restore[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v24
	v27 = *(*int64)(unsafe.Add(mBase, _c_F_restore[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v27
	v29 = int32(6)
	if base.Ui32(v29) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = v29
	goto L3
L2:
	;
	v33 = l2
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v36 = v29
	goto L6
L5:
	;
	v36 = v33
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v36 - int32(1)
	v43 = F_pg_sprintf(m, l0, int32(_a_F_restore_0), v11+int32(-48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v47 = int32(101)
	v48 = F___strchrnul(m, l0, v47)
	mBase = m.M
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v50 == v47 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v776 = F_strlen(m, l0)
	mBase = m.M
	m.G0 = v13 - int32(-64)
	return v776
L10:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L14
	}
L11:
	;
	v54 = v48
	goto L13
L12:
	;
	v54 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	v62 = v54 + int32(1)
	goto L16
L15:
	;
	if v106 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	v67 = v62 + int32(1)
	v68 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v69 = F___isspace(m, v68)
	mBase = m.M
	if v69 != 0 {
		v62 = v67
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v70 = int32(1)
	switch v68&int32(255) - int32(43) {
	case 0:
		v76 = v70
		goto L20
	default:
		v78 = v68
		v79 = v62
		v80 = v70
		goto L19
	case 2:
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	v81 = int32(0)
	v83 = v78 - int32(48)
	if base.Ui32(v83) <= base.Ui32(int32(9)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v67))))
	v78 = v77
	v79 = v67
	v80 = v76
	goto L19
L21:
	;
	v76 = int32(0)
	goto L20
L22:
	;
	v86 = v81
	v87 = v83
	v88 = v79
	goto L25
L23:
	;
	v100 = v81
	goto L24
L24:
	;
	if v80 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v90 = int32(10)
	v92 = v86*v90 - v87
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v88)+1)))
	v97 = v93 - int32(48)
	if base.Ui32(v97) < base.Ui32(v90) {
		v86 = v92
		v87 = v97
		v88 = v88 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v100 = v92
	goto L24
L27:
	;
	goto L26
L28:
	;
	v106 = int32(0) - v100
	goto L30
L29:
	;
	v106 = v100
	goto L30
L30:
	;
	goto L15
L31:
	;
	v109 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v109)
	goto L9
L32:
	;
	goto L33
L33:
	;
	v112 = v106 >> (uint(int32(31)) % 32)
	if base.Ui32(int32(4)) < base.Ui32(v106^v112-v112) {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v117 = int32(10)
	v120 = l0 + base.F32_lt(l1, float32(0))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v121 != int32(101) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = v117
	v128 = v121
	v130 = v120
	v131 = int32(0)
	goto L38
L36:
	;
	v154 = v117
	goto L37
L37:
	;
	if int32(0) < v106 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(-32)+v126))) = uint8(v128)
	v141 = base.B2i32(v128&int32(255) == int32(46))
	if v128&int32(255) == int32(46) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v142 != 0 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v142 = v126
	goto L42
L41:
	;
	v142 = v131
	goto L42
L42:
	;
	v144 = int32(1)
	v145 = v126 - v141 + v144
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v146 != int32(101) {
		v126 = v145
		v128 = v146
		v130 = v130 + v144
		v131 = v142
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v151 = v142
	goto L46
L45:
	;
	v151 = v145
	goto L46
L46:
	;
	v154 = v151
	goto L37
L47:
	;
	v164 = v154 + v106
	v166 = v164 - int32(10)
	if v36 <= v166 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v600 = v11 + int32(-32)
	v602 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v600+v36)+10)) = uint8(v602)
	v605 = v154 + v600 + v106
	v608 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v605-int32(1)))) = uint8(v608)
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L159
	} else {
		goto L160
	}
L50:
	;
	v169 = v11 + int32(-32)
	v170 = v169 + v36
	v171 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+10)) = uint8(v171)
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if int32(22) < v164 {
		goto L102
	} else {
		goto L103
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v13)+51))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v13)+43))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+44)) = v177
	v179 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+43)) = uint8(v179)
	v181 = v170
	goto L55
L54:
	;
	v181 = v169
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v166 - int32(1)
	v188 = F_pg_sprintf(m, v181+int32(11), int32(_a_F_restore_1), v13)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v192 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+41)) = uint8(v192)
	v197 = v11 + int32(-32) | int32(9)
	if (v197^l0)&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	goto L59
L59:
	;
	v275 = v11 + int32(-32) | int32(10)
	if (v275^l0)&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L60:
	;
	goto L9
L61:
	;
	goto L60
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v251)
	if v251&int32(255) == int32(0) {
		goto L61
	} else {
		goto L77
	}
L63:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v250 = v197
	v251 = v203
	v252 = l0
	goto L62
L64:
	;
	goto L65
L65:
	;
	if v197&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v207 = v197
	v209 = l0
	goto L69
L67:
	;
	v221 = v197
	v223 = l0
	goto L68
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v228 = int32(-2139062144)
	if (int32(16843008)-v225|v225)&v228 != v228 {
		v250 = v221
		v251 = v225
		v252 = v223
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v210)
	if v210 == int32(0) {
		goto L61
	} else {
		goto L71
	}
L70:
	;
	v221 = v217
	v223 = v215
	goto L68
L71:
	;
	v214 = int32(1)
	v215 = v209 + v214
	v217 = v207 + v214
	if v217&int32(3) != 0 {
		v207 = v217
		v209 = v215
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v233 = v221
	v234 = v225
	v235 = v223
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v234
	v237 = int32(4)
	v238 = v235 + v237
	v240 = v233 + v237
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v245 = int32(-2139062144)
	if (int32(16843008)-v242|v242)&v245 == v245 {
		v233 = v240
		v234 = v242
		v235 = v238
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v250 = v240
	v251 = v242
	v252 = v238
	goto L62
L76:
	;
	goto L75
L77:
	;
	v259 = v250
	v261 = v252
	goto L78
L78:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)) = uint8(v262)
	v264 = int32(1)
	if v262 != 0 {
		v259 = v259 + v264
		v261 = v261 + v264
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L61
L80:
	;
	goto L79
L81:
	;
	goto L9
L82:
	;
	goto L81
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v329)
	if v329&int32(255) == int32(0) {
		goto L82
	} else {
		goto L98
	}
L84:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	v328 = v275
	v329 = v281
	v330 = l0
	goto L83
L85:
	;
	goto L86
L86:
	;
	if v275&int32(3) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v285 = v275
	v287 = l0
	goto L90
L88:
	;
	v299 = v275
	v301 = l0
	goto L89
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v306 = int32(-2139062144)
	if (int32(16843008)-v303|v303)&v306 != v306 {
		v328 = v299
		v329 = v303
		v330 = v301
		goto L83
	} else {
		goto L94
	}
L90:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	*(*uint8)(unsafe.Add(mBase, uint32(v287))) = uint8(v288)
	if v288 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L91:
	;
	v299 = v295
	v301 = v293
	goto L89
L92:
	;
	v292 = int32(1)
	v293 = v287 + v292
	v295 = v285 + v292
	if v295&int32(3) != 0 {
		v285 = v295
		v287 = v293
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v311 = v299
	v312 = v303
	v313 = v301
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v312
	v315 = int32(4)
	v316 = v313 + v315
	v318 = v311 + v315
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v323 = int32(-2139062144)
	if (int32(16843008)-v320|v320)&v323 == v323 {
		v311 = v318
		v312 = v320
		v313 = v316
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v328 = v318
	v329 = v320
	v330 = v316
	goto L83
L97:
	;
	goto L96
L98:
	;
	v337 = v328
	v339 = v330
	goto L99
L99:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v339)+1)) = uint8(v340)
	v342 = int32(1)
	if v340 != 0 {
		v337 = v337 + v342
		v339 = v339 + v342
		goto L99
	} else {
		goto L101
	}
L100:
	;
	goto L82
L101:
	;
	goto L100
L102:
	;
	v434 = v11 + int32(-32)
	v436 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v434+v36)+11)) = uint8(v436)
	v439 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v434+v164))) = uint8(v439)
	if base.F32_lt(l1, float32(0)) != 0 {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v352 = int32(23)
	v354 = v352 - v164
	v355 = int32(3)
	v356 = v354 & v355
	if base.Ui32(v355) <= base.Ui32(v164-int32(20)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v366 = v352
	v368 = int32(0)
	goto L107
L105:
	;
	v392 = v352
	goto L106
L106:
	;
	v403 = v392
	v405 = int32(0)
	goto L111
L107:
	;
	v376 = v11 + int32(-32) + v366
	v379 = int32(4)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v376-v379)))
	*(*int32)(unsafe.Add(mBase, uint32(v376-int32(3)))) = v381
	v384 = v366 - v379
	v386 = v368 + v379
	if v386 != v354&int32(-4) {
		v366 = v384
		v368 = v386
		goto L107
	} else {
		goto L109
	}
L108:
	;
	if v356 == int32(0) {
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v392 = v384
	goto L106
L111:
	;
	v413 = v11 + int32(-32) + v403
	v414 = int32(1)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413-v414))))
	*(*uint8)(unsafe.Add(mBase, uint32(v413))) = uint8(v416)
	v421 = v405 + v414
	if v421 != v356 {
		v403 = v403 - v414
		v405 = v421
		goto L111
	} else {
		goto L113
	}
L112:
	;
	goto L102
L113:
	;
	goto L112
L114:
	;
	v443 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+41)) = uint8(v443)
	v446 = v434 | int32(9)
	if (v446^l0)&int32(3) != 0 {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	goto L116
L116:
	;
	v524 = v11 + int32(-32) | int32(10)
	if (v524^l0)&int32(3) != 0 {
		goto L141
	} else {
		goto L142
	}
L117:
	;
	goto L9
L118:
	;
	goto L117
L119:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v501))) = uint8(v500)
	if v500&int32(255) == int32(0) {
		goto L118
	} else {
		goto L134
	}
L120:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v499 = v446
	v500 = v452
	v501 = l0
	goto L119
L121:
	;
	goto L122
L122:
	;
	if v446&int32(3) != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v456 = v446
	v458 = l0
	goto L126
L124:
	;
	v470 = v446
	v472 = l0
	goto L125
L125:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v477 = int32(-2139062144)
	if (int32(16843008)-v474|v474)&v477 != v477 {
		v499 = v470
		v500 = v474
		v501 = v472
		goto L119
	} else {
		goto L130
	}
L126:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v459)
	if v459 == int32(0) {
		goto L118
	} else {
		goto L128
	}
L127:
	;
	v470 = v466
	v472 = v464
	goto L125
L128:
	;
	v463 = int32(1)
	v464 = v458 + v463
	v466 = v456 + v463
	if v466&int32(3) != 0 {
		v456 = v466
		v458 = v464
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v482 = v470
	v483 = v474
	v484 = v472
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v484))) = v483
	v486 = int32(4)
	v487 = v484 + v486
	v489 = v482 + v486
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	v494 = int32(-2139062144)
	if (int32(16843008)-v491|v491)&v494 == v494 {
		v482 = v489
		v483 = v491
		v484 = v487
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v499 = v489
	v500 = v491
	v501 = v487
	goto L119
L133:
	;
	goto L132
L134:
	;
	v508 = v499
	v510 = v501
	goto L135
L135:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)) = uint8(v511)
	v513 = int32(1)
	if v511 != 0 {
		v508 = v508 + v513
		v510 = v510 + v513
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L118
L137:
	;
	goto L136
L138:
	;
	goto L9
L139:
	;
	goto L138
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v579))) = uint8(v578)
	if v578&int32(255) == int32(0) {
		goto L139
	} else {
		goto L155
	}
L141:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524))))
	v577 = v524
	v578 = v530
	v579 = l0
	goto L140
L142:
	;
	goto L143
L143:
	;
	if v524&int32(3) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v534 = v524
	v536 = l0
	goto L147
L145:
	;
	v548 = v524
	v550 = l0
	goto L146
L146:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v555 = int32(-2139062144)
	if (int32(16843008)-v552|v552)&v555 != v555 {
		v577 = v548
		v578 = v552
		v579 = v550
		goto L140
	} else {
		goto L151
	}
L147:
	;
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v537)
	if v537 == int32(0) {
		goto L139
	} else {
		goto L149
	}
L148:
	;
	v548 = v544
	v550 = v542
	goto L146
L149:
	;
	v541 = int32(1)
	v542 = v536 + v541
	v544 = v534 + v541
	if v544&int32(3) != 0 {
		v534 = v544
		v536 = v542
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v560 = v548
	v561 = v552
	v562 = v550
	goto L152
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v561
	v564 = int32(4)
	v565 = v562 + v564
	v567 = v560 + v564
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v572 = int32(-2139062144)
	if (int32(16843008)-v569|v569)&v572 == v572 {
		v560 = v567
		v561 = v569
		v562 = v565
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v577 = v567
	v578 = v569
	v579 = v565
	goto L140
L154:
	;
	goto L153
L155:
	;
	v586 = v577
	v588 = v579
	goto L156
L156:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v588)+1)) = uint8(v589)
	v591 = int32(1)
	if v589 != 0 {
		v586 = v586 + v591
		v588 = v588 + v591
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L139
L158:
	;
	goto L157
L159:
	;
	v612 = int32(3)
	v613 = v605 - v612
	v614 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v613))) = uint8(v614)
	if (v613^l0)&v612 != 0 {
		goto L165
	} else {
		goto L166
	}
L160:
	;
	goto L161
L161:
	;
	v691 = v605 - int32(2)
	if (v691^l0)&int32(3) != 0 {
		goto L186
	} else {
		goto L187
	}
L162:
	;
	goto L9
L163:
	;
	goto L162
L164:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v669)
	if v669&int32(255) == int32(0) {
		goto L163
	} else {
		goto L179
	}
L165:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v668 = v613
	v669 = v621
	v670 = l0
	goto L164
L166:
	;
	goto L167
L167:
	;
	if v613&int32(3) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v625 = v613
	v627 = l0
	goto L171
L169:
	;
	v639 = v613
	v641 = l0
	goto L170
L170:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v646 = int32(-2139062144)
	if (int32(16843008)-v643|v643)&v646 != v646 {
		v668 = v639
		v669 = v643
		v670 = v641
		goto L164
	} else {
		goto L175
	}
L171:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v628)
	if v628 == int32(0) {
		goto L163
	} else {
		goto L173
	}
L172:
	;
	v639 = v635
	v641 = v633
	goto L170
L173:
	;
	v632 = int32(1)
	v633 = v627 + v632
	v635 = v625 + v632
	if v635&int32(3) != 0 {
		v625 = v635
		v627 = v633
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v651 = v639
	v652 = v643
	v653 = v641
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v653))) = v652
	v655 = int32(4)
	v656 = v653 + v655
	v658 = v651 + v655
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v651)+4))
	v663 = int32(-2139062144)
	if (int32(16843008)-v660|v660)&v663 == v663 {
		v651 = v658
		v652 = v660
		v653 = v656
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v668 = v658
	v669 = v660
	v670 = v656
	goto L164
L178:
	;
	goto L177
L179:
	;
	v677 = v668
	v679 = v670
	goto L180
L180:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v679)+1)) = uint8(v680)
	v682 = int32(1)
	if v680 != 0 {
		v677 = v677 + v682
		v679 = v679 + v682
		goto L180
	} else {
		goto L182
	}
L181:
	;
	goto L163
L182:
	;
	goto L181
L183:
	;
	goto L9
L184:
	;
	goto L183
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v746))) = uint8(v745)
	if v745&int32(255) == int32(0) {
		goto L184
	} else {
		goto L200
	}
L186:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v691))))
	v744 = v691
	v745 = v697
	v746 = l0
	goto L185
L187:
	;
	goto L188
L188:
	;
	if v691&int32(3) != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v701 = v691
	v703 = l0
	goto L192
L190:
	;
	v715 = v691
	v717 = l0
	goto L191
L191:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v722 = int32(-2139062144)
	if (int32(16843008)-v719|v719)&v722 != v722 {
		v744 = v715
		v745 = v719
		v746 = v717
		goto L185
	} else {
		goto L196
	}
L192:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v704)
	if v704 == int32(0) {
		goto L184
	} else {
		goto L194
	}
L193:
	;
	v715 = v711
	v717 = v709
	goto L191
L194:
	;
	v708 = int32(1)
	v709 = v703 + v708
	v711 = v701 + v708
	if v711&int32(3) != 0 {
		v701 = v711
		v703 = v709
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v727 = v715
	v728 = v719
	v729 = v717
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v728
	v731 = int32(4)
	v732 = v729 + v731
	v734 = v727 + v731
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	v739 = int32(-2139062144)
	if (int32(16843008)-v736|v736)&v739 == v739 {
		v727 = v734
		v728 = v736
		v729 = v732
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v744 = v734
	v745 = v736
	v746 = v732
	goto L185
L199:
	;
	goto L198
L200:
	;
	v753 = v744
	v755 = v746
	goto L201
L201:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v755)+1)) = uint8(v756)
	v758 = int32(1)
	if v756 != 0 {
		v753 = v753 + v758
		v755 = v755 + v758
		goto L201
	} else {
		goto L203
	}
L202:
	;
	goto L184
L203:
	;
	goto L202
}
func F_restrict_and_check_grant(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
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
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
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
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v213 int64
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	v15 = m.G0
	v17 = v15 - int32(192)
	m.G0 = v17
	switch l6 - int32(6) {
	case 0:
		v59 = int64(167503724583)
		goto L1
	default:
		goto L3
	case 3:
		goto L11
	case 8:
		goto L5
	case 10, 11, 15, 44:
		goto L9
	case 13:
		goto L10
	case 16:
		goto L8
	case 21:
		goto L4
	case 31:
		goto L7
	case 32:
		goto L12
	case 36:
		goto L2
	case 37:
		goto L6
	}
L1:
	;
	if l1 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v59 = int64(70914205040767)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L17
	}
L4:
	;
	v59 = int64(52776558145536)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v59 = int64(2199023256064)
	goto L1
L7:
	;
	v59 = int64(3298534884096)
	goto L1
L8:
	;
	v59 = int64(25769803782)
	goto L1
L9:
	;
	v59 = int64(1099511628032)
	goto L1
L10:
	;
	v59 = int64(549755814016)
	goto L1
L11:
	;
	v59 = int64(15393162792448)
	goto L1
L12:
	;
	v59 = int64(1125281431814)
	goto L1
L13:
	;
	return int64(0)
L14:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_14), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(285), int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l6
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_12), v17)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(295), int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L13
	} else {
		goto L142
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L138
	}
L22:
	;
	v213 = l3 & int64(base.Ui64(l1)>>(uint(int64(32))%64))
	if l0 != 0 {
		goto L86
	} else {
		goto L87
	}
L23:
	;
	switch l6 - int32(6) {
	case 0:
		goto L25
	default:
		goto L26
	case 3:
		goto L38
	case 8:
		goto L28
	case 10:
		goto L30
	case 11:
		goto L29
	case 13:
		goto L37
	case 15:
		goto L36
	case 16:
		goto L35
	case 21:
		goto L34
	case 31:
		goto L33
	case 32, 36:
		goto L39
	case 34:
		goto L32
	case 37:
		goto L31
	case 44:
		goto L27
	}
L24:
	;
	if v194 != int64(0) {
		goto L22
	} else {
		goto L78
	}
L25:
	;
	v186 = F_pg_class_aclmask_ext(m, l4, l5, v59, int32(1), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L76
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L13
	} else {
		goto L73
	}
L27:
	;
	v167 = F_object_aclmask_ext(m, int32(1247), l4, l5, v59, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
	} else {
		goto L72
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L69
	}
L29:
	;
	v150 = F_object_aclmask_ext(m, int32(1417), l4, l5, v59, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L68
	}
L30:
	;
	v146 = F_object_aclmask_ext(m, int32(2328), l4, l5, v59, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L67
	}
L31:
	;
	v142 = F_object_aclmask_ext(m, int32(1213), l4, l5, v59, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L66
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L63
	}
L33:
	;
	v125 = F_object_aclmask_ext(m, int32(2615), l4, l5, v59, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L62
	}
L34:
	;
	v83 = F_superuser_arg(m, l5)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L45
	}
L35:
	;
	v81 = F_pg_largeobject_aclmask_snapshot(m, l4, l5, v59, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L44
	}
L36:
	;
	v78 = F_object_aclmask_ext(m, int32(2612), l4, l5, v59, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L43
	}
L37:
	;
	v74 = F_object_aclmask_ext(m, int32(1255), l4, l5, v59, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L13
	} else {
		goto L42
	}
L38:
	;
	v70 = F_object_aclmask_ext(m, int32(1262), l4, l5, v59, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L41
	}
L39:
	;
	v66 = F_pg_class_aclmask_ext(m, l4, l5, v59, int32(1), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v194 = v66
	goto L24
L41:
	;
	v194 = v70
	goto L24
L42:
	;
	v194 = v74
	goto L24
L43:
	;
	v194 = v78
	goto L24
L44:
	;
	v194 = v81
	goto L24
L45:
	;
	if v83 != 0 {
		v194 = v59
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v87 = F_SearchSysCache1(m, int32(44), base.I64_extend_i32_u(l4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	if v87 == int32(0) {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v95 = F_SysCacheGetAttr(m, int32(44), v87, int32(3), v17+int32(191))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+191)))
	if v97 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v111 = F_aclmask(m, v108, l5, int32(10), v59, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L56
	}
L51:
	;
	v102 = F_acldefault(m, int32(27), int32(10))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v104 = base.I32_wrap_i64(v95)
	v105 = F_pg_detoast_datum(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L55
	}
L54:
	;
	v107 = int32(0)
	v108 = v102
	goto L50
L55:
	;
	v107 = v104
	v108 = v105
	goto L50
L56:
	;
	v113 = int32(0)
	if base.B2i32(v108 == v113)|base.B2i32(v108 == v107) == v113 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v108)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_ReleaseCatCache(m, v87)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v194 = v111
	goto L24
L62:
	;
	v194 = v125
	goto L24
L63:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_15), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3002), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v194 = v142
	goto L24
L67:
	;
	v194 = v146
	goto L24
L68:
	;
	v194 = v150
	goto L24
L69:
	;
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_14), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3012), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v194 = v167
	goto L24
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = l6
	F_errmsg_internal(m, int32(_a_F_restrict_and_check_grant_12), v17+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3019), int32(_a_F_restrict_and_check_grant_13))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v189 = F_pg_attribute_aclmask_ext(m, l4, l8, l5, v59, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v194 = v186 | v189
	goto L24
L78:
	;
	if l9 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v202 = base.B2i32(l6 == int32(6))
	goto L81
L80:
	;
	v202 = int32(0)
	goto L81
L81:
	;
	if v202 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	F_aclcheck_error(m, int32(1), l6, l7)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	goto L22
L84:
	;
	m.G0 = v17 + int32(192)
	return v213
L85:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), v368, int32(_a_F_restrict_and_check_grant_2))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L13
	} else {
		goto L137
	}
L86:
	;
	if v213 == int64(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v213 == int64(0) {
		goto L113
	} else {
		goto L114
	}
L89:
	;
	v219 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L13
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if l2|base.B2i32(v213 == l3) != 0 {
		goto L84
	} else {
		goto L102
	}
L92:
	;
	v221 = int32(0)
	if base.B2i32(l9 == v221)|base.B2i32(l6 != int32(6)) == v221 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v219 == int32(0) {
		goto L84
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v219 == int32(0) {
		goto L84
	} else {
		goto L99
	}
L96:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_0), v17+int32(48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v368 = int32(334)
	goto L85
L99:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_3), v17-int32(-64))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v368 = int32(339)
	goto L85
L102:
	;
	v257 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	v259 = int32(0)
	if base.B2i32(l9 == v259)|base.B2i32(l6 != int32(6)) == v259 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v257 == int32(0) {
		goto L84
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v257 == int32(0) {
		goto L84
	} else {
		goto L110
	}
L107:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_4), v17+int32(80))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	v368 = int32(347)
	goto L85
L110:
	;
	F_errcode(m, int32(117440576))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_5), v17+int32(96))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v368 = int32(352)
	goto L85
L113:
	;
	v295 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L13
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if l2|base.B2i32(v213 == l3) != 0 {
		goto L84
	} else {
		goto L126
	}
L116:
	;
	v297 = int32(0)
	if base.B2i32(l9 == v297)|base.B2i32(l6 != int32(6)) == v297 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if v295 == int32(0) {
		goto L84
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	if v295 == int32(0) {
		goto L84
	} else {
		goto L123
	}
L120:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+116)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_6), v17+int32(112))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L13
	} else {
		goto L122
	}
L122:
	;
	v368 = int32(363)
	goto L85
L123:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_7), v17+int32(128))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L13
	} else {
		goto L125
	}
L125:
	;
	v368 = int32(368)
	goto L85
L126:
	;
	v333 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	v335 = int32(0)
	if base.B2i32(l9 == v335)|base.B2i32(l6 != int32(6)) == v335 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v333 == int32(0) {
		goto L84
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v333 == int32(0) {
		goto L84
	} else {
		goto L134
	}
L131:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+148)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_8), v17+int32(144))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L13
	} else {
		goto L133
	}
L133:
	;
	v368 = int32(376)
	goto L85
L134:
	;
	F_errcode(m, int32(100663360))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L13
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = l7
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_9), v17+int32(160))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L13
	} else {
		goto L136
	}
L136:
	;
	v368 = int32(381)
	goto L85
L137:
	;
	goto L84
L138:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = l4
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_16), v17+int32(176))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(3486), int32(_a_F_restrict_and_check_grant_17))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L13
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
	F_errcode(m, int32(16797828))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l9
	F_errmsg(m, int32(_a_F_restrict_and_check_grant_10), v17+int32(32))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L13
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_restrict_and_check_grant_1), int32(2946), int32(_a_F_restrict_and_check_grant_11))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_restriction_is_always_true(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v6 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v8 == int32(52) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v11 != int32(1) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	goto L11
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v14 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v17 = F_expr_is_nonnullable(m, l0, v15, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v17
L11:
	;
	if base.B2i32(v22 != int32(0)) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v28 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v36 = int32(0)
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v36<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v44 != int32(320) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L1
L17:
	;
	v54 = v36 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v54 < v55 {
		v36 = v54
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v47 = F_restriction_is_always_true(m, l0, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	if v47 == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	return int32(1)
L21:
	;
	goto L16
}
func F_rmdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = m.Env.X__syscall_rmdir(m, l0)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v2) {
		*(*int32)(unsafe.Add(mBase, _c_F_rmdir[0])) = int32(0) - v2
		v10 = int32(-1)
	} else {
		v10 = v2
	}
	return v10
}
func F_romanian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v464 int32
	_ = v464
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v589 int32
	_ = v589
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v648 int32
	_ = v648
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v767 int32
	_ = v767
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v886 int32
	_ = v886
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v930 int32
	_ = v930
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1007 int32
	_ = v1007
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1045 int32
	_ = v1045
	var v1052 int32
	_ = v1052
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1125 int32
	_ = v1125
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1159 int32
	_ = v1159
	var v1170 int32
	_ = v1170
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1248 int32
	_ = v1248
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1274 int32
	_ = v1274
	var v1285 int32
	_ = v1285
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1366 int32
	_ = v1366
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1403 int32
	_ = v1403
	var v1410 int32
	_ = v1410
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1434 int32
	_ = v1434
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1485 int32
	_ = v1485
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1511 int32
	_ = v1511
	var v1518 int32
	_ = v1518
	var v1529 int32
	_ = v1529
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1593 int32
	_ = v1593
	var v1606 int32
	_ = v1606
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1632 int32
	_ = v1632
	var v1644 int32
	_ = v1644
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1740 int32
	_ = v1740
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1778 int32
	_ = v1778
	var v1791 int32
	_ = v1791
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1817 int32
	_ = v1817
	var v1824 int32
	_ = v1824
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1862 int32
	_ = v1862
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1913 int32
	_ = v1913
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1933 int32
	_ = v1933
	var v1939 int32
	_ = v1939
	var v1947 int32
	_ = v1947
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1986 int32
	_ = v1986
	var v1993 int32
	_ = v1993
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2037 int32
	_ = v2037
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2063 int32
	_ = v2063
	var v2070 int32
	_ = v2070
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2108 int32
	_ = v2108
	var v2115 int32
	_ = v2115
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2130 int32
	_ = v2130
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2146 int32
	_ = v2146
	var v2159 int32
	_ = v2159
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2185 int32
	_ = v2185
	var v2193 int32
	_ = v2193
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2245 int32
	_ = v2245
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2321 int32
	_ = v2321
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2333 int32
	_ = v2333
	var v2334 int32
	_ = v2334
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2384 int32
	_ = v2384
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2414 int32
	_ = v2414
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2447 int32
	_ = v2447
	var v2449 int32
	_ = v2449
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2478 int32
	_ = v2478
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2497 int32
	_ = v2497
	var v2499 int32
	_ = v2499
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2512 int32
	_ = v2512
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2557 int32
	_ = v2557
	var v2561 int32
	_ = v2561
	var v2563 int32
	_ = v2563
	var v2569 int32
	_ = v2569
	var v2586 int32
	_ = v2586
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2604 int32
	_ = v2604
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	var v2618 int32
	_ = v2618
	var v2623 int32
	_ = v2623
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2666 int32
	_ = v2666
	var v2669 int32
	_ = v2669
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2702 int32
	_ = v2702
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2726 int32
	_ = v2726
	var v2728 int32
	_ = v2728
	var v2733 int32
	_ = v2733
	var v2735 int32
	_ = v2735
	var v2742 int32
	_ = v2742
	var v2745 int32
	_ = v2745
	var v2749 int32
	_ = v2749
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2771 int32
	_ = v2771
	var v2777 int32
	_ = v2777
	var v2786 int32
	_ = v2786
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v7
	goto L2
L1:
	;
	return v2786
L2:
	;
	v15 = v10 + int32(1)
	goto L4
L3:
	;
	v117 = v7
	goto L43
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v24 <= v15 {
		v39 = v23
		v40 = v24
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L4
L8:
	;
	v109 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L41
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	v10 = v93
	goto L2
L10:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v96
	switch v33 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L38
	default:
		goto L7
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	goto L19
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	switch v27 - int32(159) {
	case 0, 4:
		goto L13
	default:
		v39 = v23
		v40 = v24
		goto L11
	}
L13:
	;
	v33 = F_find_among(m, l0, int32(_a_F_romanian_UTF_8_stem_1), int32(2), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v33 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = v38
	v40 = v37
	goto L11
L17:
	;
	if int32(0) <= v93 {
		goto L9
	} else {
		goto L37
	}
L19:
	;
	goto L20
L20:
	;
	goto L21
L21:
	;
	v48 = v10
	v50 = int32(1)
	goto L24
L23:
	;
	v93 = v78
	goto L17
L24:
	;
	if v40 <= v48 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v93 = int32(-1)
	goto L17
L27:
	;
	goto L28
L28:
	;
	v55 = v48 + int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v48))))
	if base.Ui32(v57) < base.Ui32(int32(192)) {
		v78 = v55
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = int32(1)
	if v79 < v50 {
		v48 = v78
		v50 = v50 - v79
		goto L24
	} else {
		goto L36
	}
L30:
	;
	if v40 <= v55 {
		v78 = v55
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v64 = v55
	goto L32
L32:
	;
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39+v64))))
	if int32(-65) < v67 {
		v78 = v64
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v78 = v40
	goto L29
L34:
	;
	v71 = v64 + int32(1)
	if v71 != v40 {
		v64 = v71
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L25
L37:
	;
	goto L6
L38:
	;
	v102 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	if int32(0) <= v102 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v2786 = v102
	goto L1
L41:
	;
	if v109 < int32(0) {
		v2786 = v109
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L7
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L48
L44:
	;
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2644
	v2646 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2644
	v2652 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_3), int32(5), v2646)
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L14
	} else {
		goto L634
	}
L45:
	;
	goto L44
L46:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v239 != 0 {
		v513 = v240
		goto L71
	} else {
		goto L72
	}
L47:
	;
	v239 = v232
	goto L46
L48:
	;
	if v134 <= v117 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v232 = int32(0)
	goto L47
L50:
	;
	v239 = int32(-1)
	goto L46
L51:
	;
	goto L52
L52:
	;
	v150 = int32(1)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v135))))
	if base.Ui32(v152) < base.Ui32(int32(192)) {
		v209 = v152
		v210 = v150
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if int32(259) < v209 {
		v232 = v210
		goto L47
	} else {
		goto L66
	}
L54:
	;
	v156 = v117 + int32(1)
	if v156 == v134 {
		v209 = v152
		v210 = v150
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v135))))
	v161 = v159 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v152) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v135))))
	v177 = v175 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v152) {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v165 = v117 + int32(2)
	if v165 != v134 {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v209 = v152<<(uint(int32(6))%32)&int32(1984) | v161
	v210 = int32(2)
	goto L53
L60:
	;
	goto L59
L61:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v181))))
	v209 = v194&int32(63) | (v152<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v161<<(uint(int32(12))%32) | v177<<(uint(int32(6))%32))
	v210 = int32(4)
	goto L53
L62:
	;
	v181 = v117 + int32(3)
	if v181 != v134 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v209 = v152<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v161<<(uint(int32(6))%32) | v177
	v210 = int32(3)
	goto L53
L65:
	;
	goto L64
L66:
	;
	v214 = v209 - int32(97)
	if v214 < int32(0) {
		v232 = v210
		goto L47
	} else {
		goto L67
	}
L67:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v214)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v220)>>(uint(v214&int32(7))%32))&int32(1) == int32(0) {
		v232 = v210
		goto L47
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210 + v117
	goto L69
L69:
	;
	goto L49
L70:
	;
	v2636 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_6))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L14
	} else {
		goto L631
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L136
L72:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241
	if v240 == v241 {
		v380 = v240
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v241
	if v380 == v241 {
		goto L105
	} else {
		goto L106
	}
L74:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v241))))
	if v246 != int32(117) {
		v380 = v240
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v250 = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L78
L76:
	;
	if v370 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L77:
	;
	v370 = v363
	goto L76
L78:
	;
	if v265 <= v250 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v363 = int32(0)
	goto L77
L80:
	;
	v370 = int32(-1)
	goto L76
L81:
	;
	goto L82
L82:
	;
	v281 = int32(1)
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v266))))
	if base.Ui32(v283) < base.Ui32(int32(192)) {
		v340 = v283
		v341 = v281
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if int32(259) < v340 {
		v363 = v341
		goto L77
	} else {
		goto L96
	}
L84:
	;
	v287 = v241 + int32(2)
	if v287 == v265 {
		v340 = v283
		v341 = v281
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287+v266))))
	v292 = v290 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v283) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v266))))
	v308 = v306 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v283) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v296 = v241 + int32(3)
	if v296 != v265 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v340 = v283<<(uint(int32(6))%32)&int32(1984) | v292
	v341 = int32(2)
	goto L83
L90:
	;
	goto L89
L91:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266+v312))))
	v340 = v325&int32(63) | (v283<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v292<<(uint(int32(12))%32) | v308<<(uint(int32(6))%32))
	v341 = int32(4)
	goto L83
L92:
	;
	v312 = v241 + int32(4)
	if v312 != v265 {
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v340 = v283<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v292<<(uint(int32(6))%32) | v308
	v341 = int32(3)
	goto L83
L95:
	;
	goto L94
L96:
	;
	v345 = v340 - int32(97)
	if v345 < int32(0) {
		v363 = v341
		goto L77
	} else {
		goto L97
	}
L97:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v345)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v351)>>(uint(v345&int32(7))%32))&int32(1) == int32(0) {
		v363 = v341
		goto L77
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v341 + v250
	goto L99
L99:
	;
	goto L79
L100:
	;
	v375 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_7))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L14
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v380 = v379
	goto L73
L103:
	;
	if v375 < int32(0) {
		v2786 = v375
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L43
L105:
	;
	v513 = v241
	goto L71
L106:
	;
	goto L107
L107:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383+v241))))
	if v385 != int32(105) {
		v513 = v380
		goto L71
	} else {
		goto L108
	}
L108:
	;
	v389 = v241 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v389
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L111
L109:
	;
	if v509 == int32(0) {
		goto L70
	} else {
		goto L133
	}
L110:
	;
	v509 = v502
	goto L109
L111:
	;
	if v404 <= v389 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v502 = int32(0)
	goto L110
L113:
	;
	v509 = int32(-1)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v420 = int32(1)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389+v405))))
	if base.Ui32(v422) < base.Ui32(int32(192)) {
		v479 = v422
		v480 = v420
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if int32(259) < v479 {
		v502 = v480
		goto L110
	} else {
		goto L129
	}
L117:
	;
	v426 = v241 + int32(2)
	if v426 == v404 {
		v479 = v422
		v480 = v420
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426+v405))))
	v431 = v429 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v422) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v405))))
	v447 = v445 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v422) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v435 = v241 + int32(3)
	if v435 != v404 {
		goto L119
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v479 = v422<<(uint(int32(6))%32)&int32(1984) | v431
	v480 = int32(2)
	goto L116
L123:
	;
	goto L122
L124:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405+v451))))
	v479 = v464&int32(63) | (v422<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v431<<(uint(int32(12))%32) | v447<<(uint(int32(6))%32))
	v480 = int32(4)
	goto L116
L125:
	;
	v451 = v241 + int32(4)
	if v451 != v404 {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v479 = v422<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v431<<(uint(int32(6))%32) | v447
	v480 = int32(3)
	goto L116
L128:
	;
	goto L127
L129:
	;
	v484 = v479 - int32(97)
	if v484 < int32(0) {
		v502 = v480
		goto L110
	} else {
		goto L130
	}
L130:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v484)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v490)>>(uint(v484&int32(7))%32))&int32(1) == int32(0) {
		v502 = v480
		goto L110
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480 + v389
	goto L132
L132:
	;
	goto L112
L133:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v513 = v512
	goto L71
L134:
	;
	if int32(0) <= v568 {
		v117 = v568
		goto L43
	} else {
		goto L154
	}
L136:
	;
	goto L137
L137:
	;
	goto L138
L138:
	;
	v523 = v117
	v525 = int32(1)
	goto L141
L140:
	;
	v568 = v553
	goto L134
L141:
	;
	if v513 <= v523 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	goto L140
L143:
	;
	v568 = int32(-1)
	goto L134
L144:
	;
	goto L145
L145:
	;
	v530 = v523 + int32(1)
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516+v523))))
	if base.Ui32(v532) < base.Ui32(int32(192)) {
		v553 = v530
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v554 = int32(1)
	if v554 < v525 {
		v523 = v553
		v525 = v525 - v554
		goto L141
	} else {
		goto L153
	}
L147:
	;
	if v513 <= v530 {
		v553 = v530
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v539 = v530
	goto L149
L149:
	;
	v542 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516+v539))))
	if int32(-65) < v542 {
		v553 = v539
		goto L146
	} else {
		goto L151
	}
L150:
	;
	v553 = v513
	goto L146
L151:
	;
	v546 = v539 + int32(1)
	if v546 != v513 {
		v539 = v546
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L142
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v572
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L161
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1740 = v7
	goto L417
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1714
	goto L155
L157:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1714 = v1711 + v1710
	goto L156
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L290
L159:
	;
	if v693 != 0 {
		goto L158
	} else {
		goto L183
	}
L160:
	;
	v693 = v686
	goto L159
L161:
	;
	if v572 <= v7 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v686 = int32(0)
	goto L160
L163:
	;
	v693 = int32(-1)
	goto L159
L164:
	;
	goto L165
L165:
	;
	v604 = int32(1)
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v589))))
	if base.Ui32(v606) < base.Ui32(int32(192)) {
		v663 = v606
		v664 = v604
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if int32(259) < v663 {
		v686 = v664
		goto L160
	} else {
		goto L179
	}
L167:
	;
	v610 = v7 + int32(1)
	if v610 == v572 {
		v663 = v606
		v664 = v604
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610+v589))))
	v615 = v613 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v606) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619+v589))))
	v631 = v629 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v606) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	v619 = v7 + int32(2)
	if v619 != v572 {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v663 = v606<<(uint(int32(6))%32)&int32(1984) | v615
	v664 = int32(2)
	goto L166
L173:
	;
	goto L172
L174:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589+v635))))
	v663 = v648&int32(63) | (v606<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v615<<(uint(int32(12))%32) | v631<<(uint(int32(6))%32))
	v664 = int32(4)
	goto L166
L175:
	;
	v635 = v7 + int32(3)
	if v635 != v572 {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v663 = v606<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v615<<(uint(int32(6))%32) | v631
	v664 = int32(3)
	goto L166
L178:
	;
	goto L177
L179:
	;
	v668 = v663 - int32(97)
	if v668 < int32(0) {
		v686 = v664
		goto L160
	} else {
		goto L180
	}
L180:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v668)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v674)>>(uint(v668&int32(7))%32))&int32(1) == int32(0) {
		v686 = v664
		goto L160
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v664 + v7
	goto L182
L182:
	;
	goto L162
L183:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L186
L184:
	;
	if v811 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L185:
	;
	v811 = v804
	goto L184
L186:
	;
	if v707 <= v694 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v804 = int32(0)
	goto L185
L188:
	;
	v811 = int32(-1)
	goto L184
L189:
	;
	goto L190
L190:
	;
	v723 = int32(1)
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v708))))
	if base.Ui32(v725) < base.Ui32(int32(192)) {
		v782 = v725
		v783 = v723
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if int32(259) < v782 {
		goto L204
	} else {
		goto L205
	}
L192:
	;
	v729 = v694 + int32(1)
	if v729 == v707 {
		v782 = v725
		v783 = v723
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729+v708))))
	v734 = v732 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v725) {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738+v708))))
	v750 = v748 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v725) {
		goto L200
	} else {
		goto L201
	}
L195:
	;
	v738 = v694 + int32(2)
	if v738 != v707 {
		goto L194
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v782 = v725<<(uint(int32(6))%32)&int32(1984) | v734
	v783 = int32(2)
	goto L191
L198:
	;
	goto L197
L199:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708+v754))))
	v782 = v767&int32(63) | (v725<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v734<<(uint(int32(12))%32) | v750<<(uint(int32(6))%32))
	v783 = int32(4)
	goto L191
L200:
	;
	v754 = v694 + int32(3)
	if v754 != v707 {
		goto L199
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v782 = v725<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v734<<(uint(int32(6))%32) | v750
	v783 = int32(3)
	goto L191
L203:
	;
	goto L202
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783 + v694
	goto L208
L205:
	;
	v787 = v782 - int32(97)
	if v787 < int32(0) {
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v787)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v793)>>(uint(v787&int32(7))%32))&int32(1) != 0 {
		v804 = v783
		goto L185
	} else {
		goto L207
	}
L207:
	;
	goto L204
L208:
	;
	goto L187
L209:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v835 = v825
	goto L214
L210:
	;
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v694
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L240
L212:
	;
	if int32(0) <= v930 {
		v1710 = v930
		goto L157
	} else {
		goto L237
	}
L213:
	;
	v930 = v902
	goto L212
L214:
	;
	if v826 <= v835 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v930 = int32(-1)
	goto L212
L217:
	;
	goto L218
L218:
	;
	v842 = int32(1)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835+v827))))
	if base.Ui32(v844) < base.Ui32(int32(192)) {
		v901 = v844
		v902 = v842
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if int32(259) < v901 {
		goto L232
	} else {
		goto L233
	}
L220:
	;
	v848 = v835 + int32(1)
	if v848 == v826 {
		v901 = v844
		v902 = v842
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v848+v827))))
	v853 = v851 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v844) {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v857+v827))))
	v869 = v867 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v844) {
		goto L228
	} else {
		goto L229
	}
L223:
	;
	v857 = v835 + int32(2)
	if v857 != v826 {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v901 = v844<<(uint(int32(6))%32)&int32(1984) | v853
	v902 = int32(2)
	goto L219
L226:
	;
	goto L225
L227:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v827+v873))))
	v901 = v886&int32(63) | (v844<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v853<<(uint(int32(12))%32) | v869<<(uint(int32(6))%32))
	v902 = int32(4)
	goto L219
L228:
	;
	v873 = v835 + int32(3)
	if v873 != v826 {
		goto L227
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v901 = v844<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v853<<(uint(int32(6))%32) | v869
	v902 = int32(3)
	goto L219
L231:
	;
	goto L230
L232:
	;
	v919 = v902 + v835
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v919
	v835 = v919
	goto L214
L233:
	;
	v906 = v901 - int32(97)
	if v906 < int32(0) {
		goto L232
	} else {
		goto L234
	}
L234:
	;
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v906)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v912)>>(uint(v906&int32(7))%32))&int32(1) != 0 {
		goto L213
	} else {
		goto L235
	}
L235:
	;
	goto L232
L237:
	;
	goto L211
L238:
	;
	if v1052 != 0 {
		goto L158
	} else {
		goto L262
	}
L239:
	;
	v1052 = v1045
	goto L238
L240:
	;
	if v947 <= v694 {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1045 = int32(0)
	goto L239
L242:
	;
	v1052 = int32(-1)
	goto L238
L243:
	;
	goto L244
L244:
	;
	v963 = int32(1)
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v948))))
	if base.Ui32(v965) < base.Ui32(int32(192)) {
		v1022 = v965
		v1023 = v963
		goto L245
	} else {
		goto L246
	}
L245:
	;
	if int32(259) < v1022 {
		v1045 = v1023
		goto L239
	} else {
		goto L258
	}
L246:
	;
	v969 = v694 + int32(1)
	if v969 == v947 {
		v1022 = v965
		v1023 = v963
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969+v948))))
	v974 = v972 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v965) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v978+v948))))
	v990 = v988 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v965) {
		goto L254
	} else {
		goto L255
	}
L249:
	;
	v978 = v694 + int32(2)
	if v978 != v947 {
		goto L248
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1022 = v965<<(uint(int32(6))%32)&int32(1984) | v974
	v1023 = int32(2)
	goto L245
L252:
	;
	goto L251
L253:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948+v994))))
	v1022 = v1007&int32(63) | (v965<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v974<<(uint(int32(12))%32) | v990<<(uint(int32(6))%32))
	v1023 = int32(4)
	goto L245
L254:
	;
	v994 = v694 + int32(3)
	if v994 != v947 {
		goto L253
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1022 = v965<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v974<<(uint(int32(6))%32) | v990
	v1023 = int32(3)
	goto L245
L257:
	;
	goto L256
L258:
	;
	v1027 = v1022 - int32(97)
	if v1027 < int32(0) {
		v1045 = v1023
		goto L239
	} else {
		goto L259
	}
L259:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1027)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1033)>>(uint(v1027&int32(7))%32))&int32(1) == int32(0) {
		v1045 = v1023
		goto L239
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1023 + v694
	goto L261
L261:
	;
	goto L241
L262:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1074 = v1064
	goto L265
L263:
	;
	if int32(0) <= v1170 {
		v1710 = v1170
		goto L157
	} else {
		goto L287
	}
L264:
	;
	v1170 = v1141
	goto L263
L265:
	;
	if v1065 <= v1074 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v1170 = int32(-1)
	goto L263
L268:
	;
	goto L269
L269:
	;
	v1081 = int32(1)
	v1083 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1074+v1066))))
	if base.Ui32(v1083) < base.Ui32(int32(192)) {
		v1140 = v1083
		v1141 = v1081
		goto L270
	} else {
		goto L271
	}
L270:
	;
	if int32(259) < v1140 {
		goto L264
	} else {
		goto L283
	}
L271:
	;
	v1087 = v1074 + int32(1)
	if v1087 == v1065 {
		v1140 = v1083
		v1141 = v1081
		goto L270
	} else {
		goto L272
	}
L272:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+v1066))))
	v1092 = v1090 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1083) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096+v1066))))
	v1108 = v1106 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1083) {
		goto L279
	} else {
		goto L280
	}
L274:
	;
	v1096 = v1074 + int32(2)
	if v1096 != v1065 {
		goto L273
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1140 = v1083<<(uint(int32(6))%32)&int32(1984) | v1092
	v1141 = int32(2)
	goto L270
L277:
	;
	goto L276
L278:
	;
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066+v1112))))
	v1140 = v1125&int32(63) | (v1083<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1092<<(uint(int32(12))%32) | v1108<<(uint(int32(6))%32))
	v1141 = int32(4)
	goto L270
L279:
	;
	v1112 = v1074 + int32(3)
	if v1112 != v1065 {
		goto L278
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1140 = v1083<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1092<<(uint(int32(6))%32) | v1108
	v1141 = int32(3)
	goto L270
L282:
	;
	goto L281
L283:
	;
	v1145 = v1140 - int32(97)
	if v1145 < int32(0) {
		goto L264
	} else {
		goto L284
	}
L284:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1145)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1151)>>(uint(v1145&int32(7))%32))&int32(1) == int32(0) {
		goto L264
	} else {
		goto L285
	}
L285:
	;
	v1159 = v1141 + v1074
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1159
	v1074 = v1159
	goto L265
L287:
	;
	goto L158
L288:
	;
	if v1292 != 0 {
		goto L155
	} else {
		goto L313
	}
L289:
	;
	v1292 = v1285
	goto L288
L290:
	;
	if v1188 <= v7 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1285 = int32(0)
	goto L289
L292:
	;
	v1292 = int32(-1)
	goto L288
L293:
	;
	goto L294
L294:
	;
	v1204 = int32(1)
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v1189))))
	if base.Ui32(v1206) < base.Ui32(int32(192)) {
		v1263 = v1206
		v1264 = v1204
		goto L295
	} else {
		goto L296
	}
L295:
	;
	if int32(259) < v1263 {
		goto L308
	} else {
		goto L309
	}
L296:
	;
	v1210 = v7 + int32(1)
	if v1210 == v1188 {
		v1263 = v1206
		v1264 = v1204
		goto L295
	} else {
		goto L297
	}
L297:
	;
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1210+v1189))))
	v1215 = v1213 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1206) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219+v1189))))
	v1231 = v1229 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1206) {
		goto L304
	} else {
		goto L305
	}
L299:
	;
	v1219 = v7 + int32(2)
	if v1219 != v1188 {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1263 = v1206<<(uint(int32(6))%32)&int32(1984) | v1215
	v1264 = int32(2)
	goto L295
L302:
	;
	goto L301
L303:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189+v1235))))
	v1263 = v1248&int32(63) | (v1206<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1215<<(uint(int32(12))%32) | v1231<<(uint(int32(6))%32))
	v1264 = int32(4)
	goto L295
L304:
	;
	v1235 = v7 + int32(3)
	if v1235 != v1188 {
		goto L303
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1263 = v1206<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1215<<(uint(int32(6))%32) | v1231
	v1264 = int32(3)
	goto L295
L307:
	;
	goto L306
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1264 + v7
	goto L312
L309:
	;
	v1268 = v1263 - int32(97)
	if v1268 < int32(0) {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1268)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1274)>>(uint(v1268&int32(7))%32))&int32(1) != 0 {
		v1285 = v1264
		goto L289
	} else {
		goto L311
	}
L311:
	;
	goto L308
L312:
	;
	goto L291
L313:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L316
L314:
	;
	if v1410 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L315:
	;
	v1410 = v1403
	goto L314
L316:
	;
	if v1306 <= v1293 {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	v1403 = int32(0)
	goto L315
L318:
	;
	v1410 = int32(-1)
	goto L314
L319:
	;
	goto L320
L320:
	;
	v1322 = int32(1)
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293+v1307))))
	if base.Ui32(v1324) < base.Ui32(int32(192)) {
		v1381 = v1324
		v1382 = v1322
		goto L321
	} else {
		goto L322
	}
L321:
	;
	if int32(259) < v1381 {
		goto L334
	} else {
		goto L335
	}
L322:
	;
	v1328 = v1293 + int32(1)
	if v1328 == v1306 {
		v1381 = v1324
		v1382 = v1322
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328+v1307))))
	v1333 = v1331 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1324) {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1337+v1307))))
	v1349 = v1347 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1324) {
		goto L330
	} else {
		goto L331
	}
L325:
	;
	v1337 = v1293 + int32(2)
	if v1337 != v1306 {
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1381 = v1324<<(uint(int32(6))%32)&int32(1984) | v1333
	v1382 = int32(2)
	goto L321
L328:
	;
	goto L327
L329:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307+v1353))))
	v1381 = v1366&int32(63) | (v1324<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1333<<(uint(int32(12))%32) | v1349<<(uint(int32(6))%32))
	v1382 = int32(4)
	goto L321
L330:
	;
	v1353 = v1293 + int32(3)
	if v1353 != v1306 {
		goto L329
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1381 = v1324<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1333<<(uint(int32(6))%32) | v1349
	v1382 = int32(3)
	goto L321
L333:
	;
	goto L332
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1382 + v1293
	goto L338
L335:
	;
	v1386 = v1381 - int32(97)
	if v1386 < int32(0) {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1386)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1392)>>(uint(v1386&int32(7))%32))&int32(1) != 0 {
		v1403 = v1382
		goto L315
	} else {
		goto L337
	}
L337:
	;
	goto L334
L338:
	;
	goto L317
L339:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1434 = v1424
	goto L344
L340:
	;
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1293
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L370
L342:
	;
	if int32(0) <= v1529 {
		v1710 = v1529
		goto L157
	} else {
		goto L367
	}
L343:
	;
	v1529 = v1501
	goto L342
L344:
	;
	if v1425 <= v1434 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1529 = int32(-1)
	goto L342
L347:
	;
	goto L348
L348:
	;
	v1441 = int32(1)
	v1443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434+v1426))))
	if base.Ui32(v1443) < base.Ui32(int32(192)) {
		v1500 = v1443
		v1501 = v1441
		goto L349
	} else {
		goto L350
	}
L349:
	;
	if int32(259) < v1500 {
		goto L362
	} else {
		goto L363
	}
L350:
	;
	v1447 = v1434 + int32(1)
	if v1447 == v1425 {
		v1500 = v1443
		v1501 = v1441
		goto L349
	} else {
		goto L351
	}
L351:
	;
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1447+v1426))))
	v1452 = v1450 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1443) {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1456+v1426))))
	v1468 = v1466 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1443) {
		goto L358
	} else {
		goto L359
	}
L353:
	;
	v1456 = v1434 + int32(2)
	if v1456 != v1425 {
		goto L352
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1500 = v1443<<(uint(int32(6))%32)&int32(1984) | v1452
	v1501 = int32(2)
	goto L349
L356:
	;
	goto L355
L357:
	;
	v1485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426+v1472))))
	v1500 = v1485&int32(63) | (v1443<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1452<<(uint(int32(12))%32) | v1468<<(uint(int32(6))%32))
	v1501 = int32(4)
	goto L349
L358:
	;
	v1472 = v1434 + int32(3)
	if v1472 != v1425 {
		goto L357
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1500 = v1443<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1452<<(uint(int32(6))%32) | v1468
	v1501 = int32(3)
	goto L349
L361:
	;
	goto L360
L362:
	;
	v1518 = v1501 + v1434
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1518
	v1434 = v1518
	goto L344
L363:
	;
	v1505 = v1500 - int32(97)
	if v1505 < int32(0) {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1505)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1511)>>(uint(v1505&int32(7))%32))&int32(1) != 0 {
		goto L343
	} else {
		goto L365
	}
L365:
	;
	goto L362
L367:
	;
	goto L341
L368:
	;
	if v1651 != 0 {
		goto L155
	} else {
		goto L392
	}
L369:
	;
	v1651 = v1644
	goto L368
L370:
	;
	if v1546 <= v1293 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	v1644 = int32(0)
	goto L369
L372:
	;
	v1651 = int32(-1)
	goto L368
L373:
	;
	goto L374
L374:
	;
	v1562 = int32(1)
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1293+v1547))))
	if base.Ui32(v1564) < base.Ui32(int32(192)) {
		v1621 = v1564
		v1622 = v1562
		goto L375
	} else {
		goto L376
	}
L375:
	;
	if int32(259) < v1621 {
		v1644 = v1622
		goto L369
	} else {
		goto L388
	}
L376:
	;
	v1568 = v1293 + int32(1)
	if v1568 == v1546 {
		v1621 = v1564
		v1622 = v1562
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1568+v1547))))
	v1573 = v1571 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1564) {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v1587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577+v1547))))
	v1589 = v1587 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1564) {
		goto L384
	} else {
		goto L385
	}
L379:
	;
	v1577 = v1293 + int32(2)
	if v1577 != v1546 {
		goto L378
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v1621 = v1564<<(uint(int32(6))%32)&int32(1984) | v1573
	v1622 = int32(2)
	goto L375
L382:
	;
	goto L381
L383:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1547+v1593))))
	v1621 = v1606&int32(63) | (v1564<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1573<<(uint(int32(12))%32) | v1589<<(uint(int32(6))%32))
	v1622 = int32(4)
	goto L375
L384:
	;
	v1593 = v1293 + int32(3)
	if v1593 != v1546 {
		goto L383
	} else {
		goto L387
	}
L385:
	;
	goto L386
L386:
	;
	v1621 = v1564<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1573<<(uint(int32(6))%32) | v1589
	v1622 = int32(3)
	goto L375
L387:
	;
	goto L386
L388:
	;
	v1626 = v1621 - int32(97)
	if v1626 < int32(0) {
		v1644 = v1622
		goto L369
	} else {
		goto L389
	}
L389:
	;
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1626)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1632)>>(uint(v1626&int32(7))%32))&int32(1) == int32(0) {
		v1644 = v1622
		goto L369
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1622 + v1293
	goto L391
L391:
	;
	goto L371
L392:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L395
L393:
	;
	if int32(0) <= v1706 {
		v1714 = v1706
		goto L156
	} else {
		goto L413
	}
L395:
	;
	goto L396
L396:
	;
	goto L397
L397:
	;
	v1661 = v1653
	v1663 = int32(1)
	goto L400
L399:
	;
	v1706 = v1691
	goto L393
L400:
	;
	if v1654 <= v1661 {
		goto L402
	} else {
		goto L403
	}
L401:
	;
	goto L399
L402:
	;
	v1706 = int32(-1)
	goto L393
L403:
	;
	goto L404
L404:
	;
	v1668 = v1661 + int32(1)
	v1670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652+v1661))))
	if base.Ui32(v1670) < base.Ui32(int32(192)) {
		v1691 = v1668
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v1692 = int32(1)
	if v1692 < v1663 {
		v1661 = v1691
		v1663 = v1663 - v1692
		goto L400
	} else {
		goto L412
	}
L406:
	;
	if v1654 <= v1668 {
		v1691 = v1668
		goto L405
	} else {
		goto L407
	}
L407:
	;
	v1677 = v1668
	goto L408
L408:
	;
	v1680 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1652+v1677))))
	if int32(-65) < v1680 {
		v1691 = v1677
		goto L405
	} else {
		goto L410
	}
L409:
	;
	v1691 = v1654
	goto L405
L410:
	;
	v1684 = v1677 + int32(1)
	if v1684 != v1654 {
		v1677 = v1684
		goto L408
	} else {
		goto L411
	}
L411:
	;
	goto L409
L412:
	;
	goto L401
L413:
	;
	goto L155
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2212
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2212
	v2216 = v2212 - int32(1)
	if v2216 <= v7 {
		goto L517
	} else {
		goto L518
	}
L415:
	;
	if v1835 < int32(0) {
		goto L414
	} else {
		goto L440
	}
L416:
	;
	v1835 = v1807
	goto L415
L417:
	;
	if v1731 <= v1740 {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1835 = int32(-1)
	goto L415
L420:
	;
	goto L421
L421:
	;
	v1747 = int32(1)
	v1749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1740+v1732))))
	if base.Ui32(v1749) < base.Ui32(int32(192)) {
		v1806 = v1749
		v1807 = v1747
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if int32(259) < v1806 {
		goto L435
	} else {
		goto L436
	}
L423:
	;
	v1753 = v1740 + int32(1)
	if v1753 == v1731 {
		v1806 = v1749
		v1807 = v1747
		goto L422
	} else {
		goto L424
	}
L424:
	;
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1753+v1732))))
	v1758 = v1756 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1749) {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1762+v1732))))
	v1774 = v1772 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1749) {
		goto L431
	} else {
		goto L432
	}
L426:
	;
	v1762 = v1740 + int32(2)
	if v1762 != v1731 {
		goto L425
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v1806 = v1749<<(uint(int32(6))%32)&int32(1984) | v1758
	v1807 = int32(2)
	goto L422
L429:
	;
	goto L428
L430:
	;
	v1791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1732+v1778))))
	v1806 = v1791&int32(63) | (v1749<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1758<<(uint(int32(12))%32) | v1774<<(uint(int32(6))%32))
	v1807 = int32(4)
	goto L422
L431:
	;
	v1778 = v1740 + int32(3)
	if v1778 != v1731 {
		goto L430
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v1806 = v1749<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1758<<(uint(int32(6))%32) | v1774
	v1807 = int32(3)
	goto L422
L434:
	;
	goto L433
L435:
	;
	v1824 = v1807 + v1740
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1824
	v1740 = v1824
	goto L417
L436:
	;
	v1811 = v1806 - int32(97)
	if v1811 < int32(0) {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v1817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1811)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1817)>>(uint(v1811&int32(7))%32))&int32(1) != 0 {
		goto L416
	} else {
		goto L438
	}
L438:
	;
	goto L435
L440:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1839 = v1838 + v1835
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1839
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1862 = v1839
	goto L443
L441:
	;
	if v1958 < int32(0) {
		goto L414
	} else {
		goto L465
	}
L442:
	;
	v1958 = v1929
	goto L441
L443:
	;
	if v1853 <= v1862 {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v1958 = int32(-1)
	goto L441
L446:
	;
	goto L447
L447:
	;
	v1869 = int32(1)
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1862+v1854))))
	if base.Ui32(v1871) < base.Ui32(int32(192)) {
		v1928 = v1871
		v1929 = v1869
		goto L448
	} else {
		goto L449
	}
L448:
	;
	if int32(259) < v1928 {
		goto L442
	} else {
		goto L461
	}
L449:
	;
	v1875 = v1862 + int32(1)
	if v1875 == v1853 {
		v1928 = v1871
		v1929 = v1869
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1875+v1854))))
	v1880 = v1878 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1871) {
		goto L452
	} else {
		goto L453
	}
L451:
	;
	v1894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884+v1854))))
	v1896 = v1894 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1871) {
		goto L457
	} else {
		goto L458
	}
L452:
	;
	v1884 = v1862 + int32(2)
	if v1884 != v1853 {
		goto L451
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v1928 = v1871<<(uint(int32(6))%32)&int32(1984) | v1880
	v1929 = int32(2)
	goto L448
L455:
	;
	goto L454
L456:
	;
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1854+v1900))))
	v1928 = v1913&int32(63) | (v1871<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v1880<<(uint(int32(12))%32) | v1896<<(uint(int32(6))%32))
	v1929 = int32(4)
	goto L448
L457:
	;
	v1900 = v1862 + int32(3)
	if v1900 != v1853 {
		goto L456
	} else {
		goto L460
	}
L458:
	;
	goto L459
L459:
	;
	v1928 = v1871<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v1880<<(uint(int32(6))%32) | v1896
	v1929 = int32(3)
	goto L448
L460:
	;
	goto L459
L461:
	;
	v1933 = v1928 - int32(97)
	if v1933 < int32(0) {
		goto L442
	} else {
		goto L462
	}
L462:
	;
	v1939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1933)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v1939)>>(uint(v1933&int32(7))%32))&int32(1) == int32(0) {
		goto L442
	} else {
		goto L463
	}
L463:
	;
	v1947 = v1929 + v1862
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1947
	v1862 = v1947
	goto L443
L465:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1962 = v1961 + v1958
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1962
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1962
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1986 = v1962
	goto L468
L466:
	;
	if v2081 < int32(0) {
		goto L414
	} else {
		goto L491
	}
L467:
	;
	v2081 = v2053
	goto L466
L468:
	;
	if v1977 <= v1986 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2081 = int32(-1)
	goto L466
L471:
	;
	goto L472
L472:
	;
	v1993 = int32(1)
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1986+v1978))))
	if base.Ui32(v1995) < base.Ui32(int32(192)) {
		v2052 = v1995
		v2053 = v1993
		goto L473
	} else {
		goto L474
	}
L473:
	;
	if int32(259) < v2052 {
		goto L486
	} else {
		goto L487
	}
L474:
	;
	v1999 = v1986 + int32(1)
	if v1999 == v1977 {
		v2052 = v1995
		v2053 = v1993
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v2002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1999+v1978))))
	v2004 = v2002 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1995) {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2008+v1978))))
	v2020 = v2018 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1995) {
		goto L482
	} else {
		goto L483
	}
L477:
	;
	v2008 = v1986 + int32(2)
	if v2008 != v1977 {
		goto L476
	} else {
		goto L480
	}
L478:
	;
	goto L479
L479:
	;
	v2052 = v1995<<(uint(int32(6))%32)&int32(1984) | v2004
	v2053 = int32(2)
	goto L473
L480:
	;
	goto L479
L481:
	;
	v2037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1978+v2024))))
	v2052 = v2037&int32(63) | (v1995<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v2004<<(uint(int32(12))%32) | v2020<<(uint(int32(6))%32))
	v2053 = int32(4)
	goto L473
L482:
	;
	v2024 = v1986 + int32(3)
	if v2024 != v1977 {
		goto L481
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	v2052 = v1995<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v2004<<(uint(int32(6))%32) | v2020
	v2053 = int32(3)
	goto L473
L485:
	;
	goto L484
L486:
	;
	v2070 = v2053 + v1986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2070
	v1986 = v2070
	goto L468
L487:
	;
	v2057 = v2052 - int32(97)
	if v2057 < int32(0) {
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2057)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2063)>>(uint(v2057&int32(7))%32))&int32(1) != 0 {
		goto L467
	} else {
		goto L489
	}
L489:
	;
	goto L486
L491:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2085 = v2084 + v2081
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2085
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2108 = v2085
	goto L494
L492:
	;
	if v2204 < int32(0) {
		goto L414
	} else {
		goto L516
	}
L493:
	;
	v2204 = v2175
	goto L492
L494:
	;
	if v2099 <= v2108 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2204 = int32(-1)
	goto L492
L497:
	;
	goto L498
L498:
	;
	v2115 = int32(1)
	v2117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2108+v2100))))
	if base.Ui32(v2117) < base.Ui32(int32(192)) {
		v2174 = v2117
		v2175 = v2115
		goto L499
	} else {
		goto L500
	}
L499:
	;
	if int32(259) < v2174 {
		goto L493
	} else {
		goto L512
	}
L500:
	;
	v2121 = v2108 + int32(1)
	if v2121 == v2099 {
		v2174 = v2117
		v2175 = v2115
		goto L499
	} else {
		goto L501
	}
L501:
	;
	v2124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2121+v2100))))
	v2126 = v2124 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v2117) {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	v2140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2130+v2100))))
	v2142 = v2140 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v2117) {
		goto L508
	} else {
		goto L509
	}
L503:
	;
	v2130 = v2108 + int32(2)
	if v2130 != v2099 {
		goto L502
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	v2174 = v2117<<(uint(int32(6))%32)&int32(1984) | v2126
	v2175 = int32(2)
	goto L499
L506:
	;
	goto L505
L507:
	;
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2100+v2146))))
	v2174 = v2159&int32(63) | (v2117<<(uint(int32(18))%32)&int32(_a_F_romanian_UTF_8_stem_4) | v2126<<(uint(int32(12))%32) | v2142<<(uint(int32(6))%32))
	v2175 = int32(4)
	goto L499
L508:
	;
	v2146 = v2108 + int32(3)
	if v2146 != v2099 {
		goto L507
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	v2174 = v2117<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v2126<<(uint(int32(6))%32) | v2142
	v2175 = int32(3)
	goto L499
L511:
	;
	goto L510
L512:
	;
	v2179 = v2174 - int32(97)
	if v2179 < int32(0) {
		goto L493
	} else {
		goto L513
	}
L513:
	;
	v2185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2179)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2185)>>(uint(v2179&int32(7))%32))&int32(1) == int32(0) {
		goto L493
	} else {
		goto L514
	}
L514:
	;
	v2193 = v2175 + v2108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2193
	v2108 = v2193
	goto L494
L516:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v2207 + v2204
	goto L414
L517:
	;
	v2307 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v2307)
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2309
	v2315 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_8), int32(46), v2307)
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L14
	} else {
		goto L549
	}
L518:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2218+v2216))))
	if base.B2i32(v2220&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2220)%32)&int32(_a_F_romanian_UTF_8_stem_9) == int32(0)) != 0 {
		goto L517
	} else {
		goto L519
	}
L519:
	;
	v2235 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_10), int32(16), int32(0))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L14
	} else {
		goto L520
	}
L520:
	;
	if v2235 == int32(0) {
		goto L517
	} else {
		goto L521
	}
L521:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2239
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2239 < v2241 {
		goto L517
	} else {
		goto L522
	}
L522:
	;
	switch v2235 - int32(1) {
	case 0:
		goto L529
	case 1:
		goto L528
	case 2:
		goto L527
	case 3:
		goto L526
	case 4:
		goto L525
	case 5:
		goto L524
	case 6:
		goto L523
	default:
		goto L517
	}
L523:
	;
	v2301 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_11))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L14
	} else {
		goto L546
	}
L524:
	;
	v2295 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_12))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L14
	} else {
		goto L544
	}
L525:
	;
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2267 = int32(2)
	v2269 = int32(0)
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2271-v2272 < v2267 {
		v2282 = v2269
		goto L538
	} else {
		goto L539
	}
L526:
	;
	v2262 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_13))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L14
	} else {
		goto L535
	}
L527:
	;
	v2256 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_14))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L14
	} else {
		goto L533
	}
L528:
	;
	v2250 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_15))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L14
	} else {
		goto L531
	}
L529:
	;
	v2245 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2245 {
		goto L517
	} else {
		goto L530
	}
L530:
	;
	v2786 = v2245
	goto L1
L531:
	;
	if int32(0) <= v2250 {
		goto L517
	} else {
		goto L532
	}
L532:
	;
	v2786 = v2250
	goto L1
L533:
	;
	if int32(0) <= v2256 {
		goto L517
	} else {
		goto L534
	}
L534:
	;
	v2786 = v2256
	goto L1
L535:
	;
	if int32(0) <= v2262 {
		goto L517
	} else {
		goto L536
	}
L536:
	;
	v2786 = v2262
	goto L1
L537:
	;
	if v2282 != 0 {
		goto L517
	} else {
		goto L541
	}
L538:
	;
	goto L537
L539:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2278 = F_memcmp(m, v2275+v2271-v2267, int32(_a_F_romanian_UTF_8_stem_16), v2267)
	mBase = m.M
	if v2278 != 0 {
		v2282 = v2269
		goto L538
	} else {
		goto L540
	}
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2271 - v2267
	v2282 = int32(1)
	goto L538
L541:
	;
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2283 + (v2239 - v2266)
	v2289 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_17))
	mBase = m.M
	v2290 = m.ExcPending
	if v2290 != 0 {
		goto L14
	} else {
		goto L542
	}
L542:
	;
	if int32(0) <= v2289 {
		goto L517
	} else {
		goto L543
	}
L543:
	;
	v2786 = v2289
	goto L1
L544:
	;
	if int32(0) <= v2295 {
		goto L517
	} else {
		goto L545
	}
L545:
	;
	v2786 = v2295
	goto L1
L546:
	;
	if v2301 < int32(0) {
		v2786 = v2301
		goto L1
	} else {
		goto L547
	}
L547:
	;
	goto L517
L548:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2384
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2384
	v2390 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_18), int32(62), int32(0))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L14
	} else {
		goto L576
	}
L549:
	;
	if v2315 == int32(0) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	v2321 = v2315
	goto L551
L551:
	;
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2325
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2325 < v2327 {
		goto L548
	} else {
		goto L553
	}
L552:
	;
	goto L548
L553:
	;
	switch v2321 - int32(1) {
	case 0:
		goto L560
	case 1:
		goto L559
	case 2:
		goto L558
	case 3:
		goto L557
	case 4:
		goto L556
	case 5:
		goto L555
	default:
		goto L554
	}
L554:
	;
	v2368 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v2368)
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2370
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2370
	v2376 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_8), int32(46), int32(0))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L14
	} else {
		goto L573
	}
L555:
	;
	v2363 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_19))
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L14
	} else {
		goto L571
	}
L556:
	;
	v2357 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_20))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L14
	} else {
		goto L569
	}
L557:
	;
	v2351 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_21))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L14
	} else {
		goto L567
	}
L558:
	;
	v2345 = F_slice_from_s(m, l0, int32(2), int32(_a_F_romanian_UTF_8_stem_22))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L14
	} else {
		goto L565
	}
L559:
	;
	v2339 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_23))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L14
	} else {
		goto L563
	}
L560:
	;
	v2333 = F_slice_from_s(m, l0, int32(4), int32(_a_F_romanian_UTF_8_stem_24))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L14
	} else {
		goto L561
	}
L561:
	;
	if int32(0) <= v2333 {
		goto L554
	} else {
		goto L562
	}
L562:
	;
	v2786 = v2333
	goto L1
L563:
	;
	if int32(0) <= v2339 {
		goto L554
	} else {
		goto L564
	}
L564:
	;
	v2786 = v2339
	goto L1
L565:
	;
	if int32(0) <= v2345 {
		goto L554
	} else {
		goto L566
	}
L566:
	;
	v2786 = v2345
	goto L1
L567:
	;
	if int32(0) <= v2351 {
		goto L554
	} else {
		goto L568
	}
L568:
	;
	v2786 = v2351
	goto L1
L569:
	;
	if int32(0) <= v2357 {
		goto L554
	} else {
		goto L570
	}
L570:
	;
	v2786 = v2357
	goto L1
L571:
	;
	if v2363 < int32(0) {
		v2786 = v2363
		goto L1
	} else {
		goto L572
	}
L572:
	;
	goto L554
L573:
	;
	if v2376 != 0 {
		v2321 = v2376
		goto L551
	} else {
		goto L574
	}
L574:
	;
	goto L552
L575:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2439
	v2441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v2441 != 0 {
		goto L45
	} else {
		goto L593
	}
L576:
	;
	if v2390 == int32(0) {
		goto L575
	} else {
		goto L577
	}
L577:
	;
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2394
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2394 < v2396 {
		goto L575
	} else {
		goto L578
	}
L578:
	;
	switch v2390 - int32(1) {
	case 0:
		goto L582
	case 1:
		goto L581
	case 2:
		goto L580
	default:
		goto L579
	}
L579:
	;
	v2436 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v2436)
	goto L45
L580:
	;
	v2431 = F_slice_from_s(m, l0, int32(3), int32(_a_F_romanian_UTF_8_stem_25))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L14
	} else {
		goto L591
	}
L581:
	;
	v2403 = int32(2)
	v2405 = int32(0)
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2407-v2408 < v2403 {
		v2418 = v2405
		goto L585
	} else {
		goto L586
	}
L582:
	;
	v2400 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2400 {
		goto L579
	} else {
		goto L583
	}
L583:
	;
	v2786 = v2400
	goto L1
L584:
	;
	if v2418 == int32(0) {
		goto L575
	} else {
		goto L588
	}
L585:
	;
	goto L584
L586:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2414 = F_memcmp(m, v2411+v2407-v2403, int32(_a_F_romanian_UTF_8_stem_26), v2403)
	mBase = m.M
	if v2414 != 0 {
		v2418 = v2405
		goto L585
	} else {
		goto L587
	}
L587:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2407 - v2403
	v2418 = int32(1)
	goto L585
L588:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2421
	v2425 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_27))
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L14
	} else {
		goto L589
	}
L589:
	;
	if int32(0) <= v2425 {
		goto L579
	} else {
		goto L590
	}
L590:
	;
	v2786 = v2425
	goto L1
L591:
	;
	if v2431 < int32(0) {
		v2786 = v2431
		goto L1
	} else {
		goto L592
	}
L592:
	;
	goto L579
L593:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2443 < v2444 {
		v2623 = int32(0)
		goto L594
	} else {
		goto L595
	}
L594:
	;
	if v2623 == int32(0) {
		goto L45
	} else {
		goto L626
	}
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2443
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2444
	v2449 = int32(0)
	v2453 = F_find_among_b(m, l0, int32(_a_F_romanian_UTF_8_stem_28), int32(94), v2449)
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L14
	} else {
		goto L597
	}
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v2447
	v2623 = v2618
	goto L594
L597:
	;
	if v2453 == int32(0) {
		v2618 = v2449
		goto L596
	} else {
		goto L598
	}
L598:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2457
	v2459 = int32(1)
	switch v2453 - v2459 {
	case 0:
		goto L600
	case 1:
		goto L599
	default:
		v2618 = v2459
		goto L596
	}
L599:
	;
	v2615 = F_slice_del(m, l0)
	mBase = m.M
	if v2615 < int32(0) {
		v2623 = v2615
		goto L594
	} else {
		goto L625
	}
L600:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2463 = int32(0)
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L603
L601:
	;
	if v2593 != 0 {
		goto L619
	} else {
		goto L620
	}
L602:
	;
	v2593 = v2586
	goto L601
L603:
	;
	if v2476 <= v2477 {
		v2586 = int32(-1)
		goto L602
	} else {
		goto L605
	}
L604:
	;
	v2586 = int32(0)
	goto L602
L605:
	;
	v2494 = int32(1)
	v2495 = v2476 - v2494
	v2497 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2478+v2495))))
	v2499 = v2497 & int32(255)
	if base.B2i32(v2495 == v2477)|base.B2i32(int32(0) <= v2497) != 0 {
		v2557 = v2499
		v2561 = v2494
		goto L606
	} else {
		goto L607
	}
L606:
	;
	if int32(259) < v2557 {
		goto L614
	} else {
		goto L615
	}
L607:
	;
	v2506 = v2499 & int32(63)
	v2508 = v2476 - int32(2)
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478+v2508))))
	v2512 = v2510 << (uint(int32(6)) % 32)
	if base.B2i32(v2508 != v2477)&base.B2i32(base.Ui32(v2510) < base.Ui32(int32(192))) == int32(0) {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	v2557 = v2512&int32(1984) | v2506
	v2561 = int32(2)
	goto L606
L609:
	;
	goto L610
L610:
	;
	v2525 = v2512&int32(4032) | v2506
	v2527 = v2476 - int32(3)
	v2529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2478+v2527))))
	if base.B2i32(v2527 != v2477)&base.B2i32(base.Ui32(v2529) < base.Ui32(int32(224))) == int32(0) {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v2557 = v2529<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_5) | v2525
	v2561 = int32(3)
	goto L606
L612:
	;
	goto L613
L613:
	;
	v2547 = int32(4)
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2476+v2478-v2547))))
	v2557 = v2529<<(uint(int32(12))%32)&int32(_a_F_romanian_UTF_8_stem_29) | v2549&int32(7)<<(uint(int32(18))%32) | v2525
	v2561 = v2547
	goto L606
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2476 - v2561
	goto L618
L615:
	;
	v2563 = v2557 - int32(97)
	if v2563 < int32(0) {
		goto L614
	} else {
		goto L616
	}
L616:
	;
	v2569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2563)>>(uint(int32(3))%32)))+uint32(_c_F_romanian_UTF_8_stem[0]))))
	if int32(base.Ui32(v2569)>>(uint(v2563&int32(7))%32))&int32(1) == int32(0) {
		goto L614
	} else {
		goto L617
	}
L617:
	;
	v2593 = v2561
	goto L601
L618:
	;
	goto L604
L619:
	;
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2596 = v2594 + (v2457 - v2462)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2596
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2596 <= v2598 {
		v2618 = v2463
		goto L596
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	v2612 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2612 {
		v2618 = int32(1)
		goto L596
	} else {
		goto L624
	}
L622:
	;
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2600+v2596-int32(1)))))
	if v2604 != int32(117) {
		v2618 = v2463
		goto L596
	} else {
		goto L623
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2596 - int32(1)
	goto L621
L624:
	;
	v2623 = v2612
	goto L594
L625:
	;
	v2618 = v2459
	goto L596
L626:
	;
	v2628 = int32(0)
	v2629 = base.B2i32(v2623 < v2628)
	if v2629 == v2628 {
		goto L45
	} else {
		goto L627
	}
L627:
	;
	if v2623 < v2628 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v2633 = v2623
	goto L630
L629:
	;
	v2633 = int32(1)
	goto L630
L630:
	;
	v2786 = v2633
	goto L1
L631:
	;
	if int32(0) <= v2636 {
		goto L43
	} else {
		goto L632
	}
L632:
	;
	v2786 = v2636
	goto L1
L633:
	;
	if v2666 < int32(0) {
		v2786 = v2666
		goto L1
	} else {
		goto L640
	}
L634:
	;
	if v2652 == int32(0) {
		v2666 = v2646
		goto L633
	} else {
		goto L635
	}
L635:
	;
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2656
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v2656 < v2658 {
		v2666 = v2646
		goto L633
	} else {
		goto L636
	}
L636:
	;
	v2661 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2661 {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v2664 = int32(1)
	goto L639
L638:
	;
	v2664 = v2661
	goto L639
L639:
	;
	v2666 = v2664
	goto L633
L640:
	;
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2669
	goto L642
L641:
	;
	if v2777 < int32(0) {
		v2786 = v2777
		goto L1
	} else {
		goto L679
	}
L642:
	;
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2677
	v2679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2679 <= v2677 {
		goto L645
	} else {
		goto L646
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2677
	v2777 = int32(1)
	goto L641
L644:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L658
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2677
	v2717 = v2677
	v2718 = v2679
	goto L644
L646:
	;
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2681+v2677))))
	v2685 = v2683 - int32(73)
	v2686 = int32(0)
	if base.B2i32(v2685 == v2686)|base.B2i32(v2685 == int32(12)) == v2686 {
		goto L645
	} else {
		goto L647
	}
L647:
	;
	v2696 = F_find_among(m, l0, int32(_a_F_romanian_UTF_8_stem_30), int32(3), int32(0))
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L14
	} else {
		goto L648
	}
L648:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2698
	switch v2696 - int32(1) {
	case 0:
		goto L650
	case 1:
		goto L649
	case 2:
		goto L651
	default:
		goto L642
	}
L649:
	;
	v2711 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_31))
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L14
	} else {
		goto L654
	}
L650:
	;
	v2705 = F_slice_from_s(m, l0, int32(1), int32(_a_F_romanian_UTF_8_stem_32))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L14
	} else {
		goto L652
	}
L651:
	;
	v2702 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2717 = v2698
	v2718 = v2702
	goto L644
L652:
	;
	if int32(0) <= v2705 {
		goto L642
	} else {
		goto L653
	}
L653:
	;
	v2777 = v2705
	goto L641
L654:
	;
	if int32(0) <= v2711 {
		goto L642
	} else {
		goto L655
	}
L655:
	;
	v2777 = v2711
	goto L641
L656:
	;
	if int32(0) <= v2771 {
		goto L676
	} else {
		goto L677
	}
L658:
	;
	goto L659
L659:
	;
	goto L660
L660:
	;
	v2726 = v2717
	v2728 = int32(1)
	goto L663
L662:
	;
	v2771 = v2756
	goto L656
L663:
	;
	if v2718 <= v2726 {
		goto L665
	} else {
		goto L666
	}
L664:
	;
	goto L662
L665:
	;
	v2771 = int32(-1)
	goto L656
L666:
	;
	goto L667
L667:
	;
	v2733 = v2726 + int32(1)
	v2735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2719+v2726))))
	if base.Ui32(v2735) < base.Ui32(int32(192)) {
		v2756 = v2733
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v2757 = int32(1)
	if v2757 < v2728 {
		v2726 = v2756
		v2728 = v2728 - v2757
		goto L663
	} else {
		goto L675
	}
L669:
	;
	if v2718 <= v2733 {
		v2756 = v2733
		goto L668
	} else {
		goto L670
	}
L670:
	;
	v2742 = v2733
	goto L671
L671:
	;
	v2745 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2719+v2742))))
	if int32(-65) < v2745 {
		v2756 = v2742
		goto L668
	} else {
		goto L673
	}
L672:
	;
	v2756 = v2718
	goto L668
L673:
	;
	v2749 = v2742 + int32(1)
	if v2749 != v2718 {
		v2742 = v2749
		goto L671
	} else {
		goto L674
	}
L674:
	;
	goto L672
L675:
	;
	goto L664
L676:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2771
	goto L642
L677:
	;
	goto L678
L678:
	;
	goto L643
L679:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2669
	return int32(1)
}
func F_rstacktoodeep(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_rstacktoodeep[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_rstacktoodeep[1]))
	v7 = m.G0
	v8 = v6 - v7
	v10 = v8 >> (uint(int32(31)) % 32)
	return base.B2i32(v4 < v8^v10-v10) & base.B2i32(v6 != int32(0))
}
func F_rtrim(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14239(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
